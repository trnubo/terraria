package main

import (
	"flag"
	"strings"

	// "github.com/joshbarrass/TerrariaServerWrapper/internal"
	"github.com/trnubo/terraria/TerrariaServerWrapper/internal"
	"github.com/sirupsen/logrus"
)

func main() {
	// Use flag to get command line args, we may want to accept other flags in future
	flag.Parse()
	logrus.Infof("command: %s", strings.Join(flag.Args()," "))

	// Create the command wrapper
	server, err := internal.NewServer(flag.Args())
	if err != nil {
		logrus.Fatalf("An error occurred starting the server: %s", err)
	}

	// Start the server and wait for it to exit
	err = server.Start()
	if err != nil {
		logrus.Errorf("An error occurred in the server: %s", err)
	} else {
		logrus.Infof("Exited with status code %d", server.GetExitCode())
	}
}
