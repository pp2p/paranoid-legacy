package commands

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/cpssd/paranoid/paranoid-cli/tls"
	"os"
	"os/user"
	"path"
	"path/filepath"
)

func Secure(c *cli.Context) {
	args := c.Args()
	if len(args) < 1 {
		cli.ShowCommandHelp(c, "secure")
		os.Exit(0)
	}

	pfsname := args[0]
	usr, err := user.Current()
	if err != nil {
		fmt.Println("FATAL: Couldn't get current user")
		Log.Fatal("Couldn't get current user:", err)
	}

	homeDir := usr.HomeDir
	pfsDir, err := filepath.Abs(path.Join(homeDir, ".pfs", "filesystems", pfsname))
	if err != nil {
		fmt.Println("FATAL: Could not get absolute path to paranoid filesystem.")
		Log.Fatal("Could not get absolute path to paranoid filesystem.")
	}
	if !pathExists(pfsDir) {
		fmt.Println("FATAL: Paranoid filesystem does not exist:", pfsname)
		Log.Fatal("Paranoid filesystem does not exist:", pfsname)
	}

	certPath := path.Join(pfsDir, "meta", "cert.pem")
	keyPath := path.Join(pfsDir, "meta", "key.pem")
	if c.Bool("force") {
		os.Remove(certPath)
		os.Remove(keyPath)
	} else {
		if pathExists(certPath) || pathExists(keyPath) {
			fmt.Println("FATAL: Paranoid filesystem already secured.",
				"Run with --force to overwrite existing security files.")
			Log.Fatal("Paranoid filesystem already secured.",
				"Run with --force to overwrite existing security files.")
		}
	}

	if (c.String("cert") != "") && (c.String("key") != "") {
		Log.Info("Using existing certificate.")
		err = os.Link(c.String("cert"), certPath)
		if err != nil {
			fmt.Println("FATAL: Failed to copy cert file:", err)
			Log.Fatal("Failed to copy cert file")
		}
		err = os.Link(c.String("key"), keyPath)
		if err != nil {
			fmt.Println("FATAL: Failed to copy key file:", err)
			Log.Fatal("Failed to copy key file")
		}
	} else {
		fmt.Println("Generating TLS certificate. Please follow the given instructions.")
		err = tls.GenCertificate(pfsDir)
		if err != nil {
			fmt.Println("FATAL: Failed to generate certificate:", err)
			Log.Fatal("Failed to generate certificate:", err)
		}
	}
}
