package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/cpssd/paranoid/libpfs/commands"
	"github.com/cpssd/paranoid/libpfs/encryption"
	"github.com/cpssd/paranoid/logger"
	"github.com/cpssd/paranoid/pfsd/dnetclient"
	"github.com/cpssd/paranoid/pfsd/globals"
	"github.com/cpssd/paranoid/pfsd/intercom"
	"github.com/cpssd/paranoid/pfsd/keyman"
	"github.com/cpssd/paranoid/pfsd/pfi"
	"github.com/cpssd/paranoid/pfsd/pnetclient"
	"github.com/cpssd/paranoid/pfsd/pnetserver"
	"github.com/cpssd/paranoid/pfsd/upnp"
	pb "github.com/cpssd/paranoid/proto/paranoidnetwork"
	rpb "github.com/cpssd/paranoid/proto/raft"
	"github.com/cpssd/paranoid/raft"
	"github.com/cpssd/paranoid/raft/raftlog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
	"net"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
)

var (
	srv *grpc.Server
	log *logger.ParanoidLogger

	certFile   = flag.String("cert", "", "TLS certificate file - if empty connection will be unencrypted")
	keyFile    = flag.String("key", "", "TLS key file - if empty connection will be unencrypted")
	noNetwork  = flag.Bool("no_networking", false, "Do not perform any networking")
	skipVerify = flag.Bool("skip_verification", false,
		"skip verification of TLS certificate chain and hostname - not recommended unless using self-signed certs")
	verbose = flag.Bool("v", false, "Use verbose logging")
)

func startRPCServer(lis *net.Listener) {
	var opts []grpc.ServerOption
	if globals.TLSEnabled {
		log.Info("Starting ParanoidNetwork server with TLS.")
		creds, err := credentials.NewServerTLSFromFile(*certFile, *keyFile)
		if err != nil {
			log.Fatal("Failed to generate TLS credentials:", err)
		}
		opts = []grpc.ServerOption{grpc.Creds(creds)}
	} else {
		log.Info("Starting ParanoidNetwork server without TLS.")
	}
	srv = grpc.NewServer(opts...)

	pb.RegisterParanoidNetworkServer(srv, &pnetserver.ParanoidServer{})
	nodeDetails := raft.Node{
		IP:         globals.ThisNode.IP,
		Port:       globals.ThisNode.Port,
		CommonName: globals.ThisNode.CommonName,
		NodeID:     globals.ThisNode.UUID,
	}

	//First node to join a given cluster
	if len(globals.Nodes.GetAll()) == 0 {
		globals.RaftNetworkServer = raft.NewRaftNetworkServer(
			nodeDetails,
			globals.ParanoidDir,
			path.Join(globals.ParanoidDir, "meta", "raft"),
			&raft.StartConfiguration{
				Peers: []raft.Node{},
			},
			globals.TLSEnabled,
			globals.TLSSkipVerify)
	} else {
		globals.RaftNetworkServer = raft.NewRaftNetworkServer(
			nodeDetails,
			globals.ParanoidDir,
			path.Join(globals.ParanoidDir, "meta", "raft"),
			nil,
			globals.TLSEnabled,
			globals.TLSSkipVerify)
	}

	rpb.RegisterRaftNetworkServer(srv, globals.RaftNetworkServer)

	globals.Wait.Add(1)
	go func() {
		defer globals.Wait.Done()
		err := srv.Serve(*lis)
		log.Info("Paranoid network server stopped")
		if err != nil && globals.ShuttingDown == false {
			log.Fatal("Server stopped because of an error:", err)
		}
	}()

	//Do we need to request to join a cluster
	if globals.RaftNetworkServer.State.Configuration.HasConfiguration() == false {
		log.Info("Attempting to join raft cluster")
		err := dnetclient.JoinCluster()
		if err != nil {
			log.Fatal("Unable to join a raft cluster")
		}
	}
}

func setupLogging() {
	logDir := path.Join(globals.ParanoidDir, "meta", "logs")

	log = logger.New("main", "pfsd", logDir)
	dnetclient.Log = logger.New("dnetclient", "pfsd", logDir)
	pnetclient.Log = logger.New("pnetclient", "pfsd", logDir)
	pnetserver.Log = logger.New("pnetserver", "pfsd", logDir)
	upnp.Log = logger.New("upnp", "pfsd", logDir)
	keyman.Log = logger.New("keyman", "pfsd", logDir)
	raft.Log = logger.New("raft", "pfsd", logDir)
	raftlog.Log = logger.New("raftlog", "pfsd", logDir)
	commands.Log = logger.New("libpfs", "pfsd", logDir)
	intercom.Log = logger.New("intercom", "pfsd", logDir)
	globals.Log = logger.New("globals", "pfsd", logDir)

	log.SetOutput(logger.STDERR | logger.LOGFILE)
	dnetclient.Log.SetOutput(logger.STDERR | logger.LOGFILE)
	pnetclient.Log.SetOutput(logger.STDERR | logger.LOGFILE)
	pnetserver.Log.SetOutput(logger.STDERR | logger.LOGFILE)
	upnp.Log.SetOutput(logger.STDERR | logger.LOGFILE)
	keyman.Log.SetOutput(logger.STDERR | logger.LOGFILE)
	raft.Log.SetOutput(logger.STDERR | logger.LOGFILE)
	raftlog.Log.SetOutput(logger.STDERR | logger.LOGFILE)
	commands.Log.SetOutput(logger.STDERR | logger.LOGFILE)
	intercom.Log.SetOutput(logger.STDERR | logger.LOGFILE)
	globals.Log.SetOutput(logger.STDERR | logger.LOGFILE)

	if *verbose {
		commands.Log.SetLogLevel(logger.VERBOSE)
	}
}

func getFileSystemAttributes() {
	attributesJson, err := ioutil.ReadFile(path.Join(globals.ParanoidDir, "meta", "attributes"))
	if err != nil {
		log.Fatal("unable to read file system attributes:", err)
	}

	attributes := &globals.FileSystemAttributes{}
	err = json.Unmarshal(attributesJson, attributes)
	if err != nil {
		log.Fatal("unable to read file system attributes:", err)
	}

	globals.Encrypted = attributes.Encrypted
	encryption.Encrypted = attributes.Encrypted

	if attributes.Encrypted {
		if !attributes.KeyGenerated {
			//If a key has not yet been generated for this file system, one must be generated
			globals.EncryptionKey, err = keyman.GenerateKey(32)
			if err != nil {
				log.Fatal("unable to generate encryption key:", err)
			}
			attributes.KeyGenerated = true

			cipherB, err := encryption.GenerateAESCipherBlock(globals.EncryptionKey.GetBytes())
			if err != nil {
				log.Fatal("unable to generate cipher block:", err)
			}
			encryption.SetCipher(cipherB)

			if attributes.NetworkOff {
				//If networking is turned off, save the key to a file
				attributes.EncryptionKey = *globals.EncryptionKey
			}
		} else if attributes.NetworkOff {
			//If networking is off, load the key from the file
			globals.EncryptionKey = &attributes.EncryptionKey
			cipherB, err := encryption.GenerateAESCipherBlock(globals.EncryptionKey.GetBytes())
			if err != nil {
				log.Fatal("unable to generate cipher block:", err)
			}
			encryption.SetCipher(cipherB)
		}
	}

	attributesJson, err = json.Marshal(attributes)
	if err != nil {
		log.Fatal("unable to save new file system attributes to file:", err)
	}

	err = ioutil.WriteFile(path.Join(globals.ParanoidDir, "meta", "attributes"), attributesJson, 0600)
	if err != nil {
		log.Fatal("unable to save new file system attributes to file:", err)
	}
}

func main() {
	flag.Parse()

	if len(flag.Args()) < 5 {
		fmt.Print("Usage:\n\tpfsd <paranoid_directory> <mount_point> <Discovery Server> <Discovery Port> <Discovery Pool>\n")
		os.Exit(1)
	}

	paranoidDirAbs, err := filepath.Abs(flag.Arg(0))
	if err != nil {
		fmt.Println("FATAL: Could not get absolute paranoid dir:", err)
		os.Exit(1)
	}

	mountPointAbs, err := filepath.Abs(flag.Arg(1))
	if err != nil {
		fmt.Println("FATAL: Could not get absolute mount point:", err)
		os.Exit(1)
	}

	globals.ParanoidDir = paranoidDirAbs
	globals.MountPoint = mountPointAbs
	setupLogging()

	getFileSystemAttributes()

	globals.TLSSkipVerify = *skipVerify
	if *certFile != "" && *keyFile != "" {
		globals.TLSEnabled = true
		if !globals.TLSSkipVerify {
			cn, err := getCommonNameFromCert(*certFile)
			if err != nil {
				log.Fatal("Could not get CN from provided TLS cert:", err)
			}
			globals.ThisNode.CommonName = cn
		}
	} else {
		globals.TLSEnabled = false
	}

	if !*noNetwork {
		discoveryPort, err := strconv.Atoi(flag.Arg(3))
		if err != nil || discoveryPort < 1 || discoveryPort > 65535 {
			log.Fatal("Discovery port must be a number between 1 and 65535, inclusive.")
		}

		uuid, err := ioutil.ReadFile(path.Join(globals.ParanoidDir, "meta", "uuid"))
		if err != nil {
			log.Fatal("Could not get node UUID:", err)
		}
		globals.ThisNode.UUID = string(uuid)

		ip, err := upnp.GetIP()
		if err != nil {
			log.Fatal("Could not get IP:", err)
		}

		//Asking for port 0 requests a random free port from the OS.
		lis, err := net.Listen("tcp", ip+":0")
		if err != nil {
			log.Fatalf("Failed to start listening : %v.\n", err)
		}
		splits := strings.Split(lis.Addr().String(), ":")
		port := splits[len(splits)-1]
		portInt, err := strconv.Atoi(port)
		if err != nil {
			log.Fatal("Could not parse port", splits[len(splits)-1], " Error :", err)
		}
		globals.ThisNode.Port = port

		//Try and set up uPnP. Otherwise use internal IP.
		globals.UPnPEnabled = false
		err = upnp.DiscoverDevices()
		if err == nil {
			log.Info("UPnP devices available")
			externalPort, err := upnp.AddPortMapping(ip, portInt)
			if err == nil {
				log.Info("UPnP port mapping enabled")
				port = strconv.Itoa(externalPort)
				globals.ThisNode.Port = port
				globals.UPnPEnabled = true
			}
		}

		globals.ThisNode.IP, err = upnp.GetIP()
		if err != nil {
			log.Fatal("Can't get IP. Error : ", err)
		}
		log.Info("Peer address:", globals.ThisNode.IP+":"+globals.ThisNode.Port)

		if _, err := os.Stat(globals.ParanoidDir); os.IsNotExist(err) {
			log.Fatal("Path", globals.ParanoidDir, "does not exist.")
		}
		if _, err := os.Stat(path.Join(globals.ParanoidDir, "meta")); os.IsNotExist(err) {
			log.Fatal("Path", globals.ParanoidDir, "is not valid PFS root.")
		}

		dnetclient.SetDiscovery(flag.Arg(2), flag.Arg(3))
		dnetclient.JoinDiscovery(flag.Arg(4))
		startRPCServer(&lis)
	}
	createPid("pfsd")
	pfi.StartPfi(*verbose)

	if globals.SystemLocked {
		globals.Wait.Add(1)
		go UnlockWorker()
	}
	intercom.RunServer(path.Join(globals.ParanoidDir, "meta"))

	HandleSignals()
}

func createPid(processName string) {
	processID := os.Getpid()
	pid := []byte(strconv.Itoa(processID))
	err := ioutil.WriteFile(path.Join(globals.ParanoidDir, "meta", processName+".pid"), pid, 0600)
	if err != nil {
		log.Fatal("Failed to create PID file", err)
	}
}
