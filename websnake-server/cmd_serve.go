package server

import (
	"os"
	"os/signal"

	"github.com/fsufitch/websnake/websnake-server/config"
	"github.com/fsufitch/websnake/websnake-server/log"
	"github.com/urfave/cli"
)

var serveCommand = cli.Command{
	Name:   "serve",
	Usage:  "start the prez-tweet server using configuration from the environment",
	Action: serve,
	Flags:  []cli.Flag{},
}

func serve(_ *cli.Context) error {
	err := config.ConfigureLogging()
	if err != nil {
		return err
	}
	interruptChan := make(chan os.Signal)
	signal.Notify(interruptChan, os.Interrupt)

	serverQuitChan := make(chan error)
	go func() {
		serverQuitChan <- startServer()
	}()

	select {
	case <-interruptChan:
		log.Error.Print("Received interrupt, quitting.")
		return nil
	case err := <-serverQuitChan:
		log.Error.Print("Fatal error:", err)
		return err
	}
}
