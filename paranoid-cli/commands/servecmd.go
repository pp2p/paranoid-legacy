package commands

import (
	"fmt"
	"github.com/codegangsta/cli"
	pb "github.com/cpssd/paranoid/proto/fileserver"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
	"time"
)

//Adds files from Paranoid File Server
func Serve(c *cli.Context) {
	args := c.Args()
	if len(args) < 2 {
		cli.ShowCommandHelp(c, "serve")
		os.Exit(1)
	}
	requestTimeout := c.Int("timeout")
	requestLimit := c.Int("access")
	file := args[1]

	usr, err := user.Current()
	if err != nil {
		fmt.Println("Unable to get information on current user:", err)
		Log.Fatal("Could not get user information:", err)
	}
	ip, port, uuid, pool := getFsMeta(usr, args[0])

	address := ip + ":" + port
	serveFilePath, err := filepath.Abs(file)
	serveData, err := ioutil.ReadFile(serveFilePath)

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTimeout(2*time.Second))
	opts = append(opts, grpc.WithInsecure())
	connection, err := grpc.Dial(address, opts...)
	if err != nil {
		fmt.Println("Failed to Connect to Discovery Share Server")
		Log.Fatal("Unable to Connect to Discovery Share Server", err)
	}
	defer connection.Close()

	serverClient := pb.NewFileserverClient(connection)
	response, err := serverClient.ServeFile(context.Background(),
		&pb.ServeRequest{
			Uuid:     uuid,
			FilePath: serveFilePath,
			FileData: serveData,
			Pool:     pool,
			Timeout:  int32(requestTimeout),
			Limit:    int32(requestLimit),
		})
	if err != nil {
		fmt.Println("Unable to send File to Discovery Share Server")
		Log.Fatal("Couldn't message Discovery Share Server", err)
	}
	fmt.Println("File now avaliable at:", "http://"+ip+response.ServerPort+"/"+response.ServeResponse)
}
