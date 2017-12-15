package server

import "github.com/urfave/cli"

// Main exposes the main CLI entry point to the server
func Main() {
	app := cli.NewApp()
	app.Name = "websnake"
	app.Usage = "API backend for a small web snake game"
	app.EnableBashCompletion = true
	app.Commands = []cli.Command{serveCommand}
	app.RunAndExitOnError()
}
