package main

import (
	"os"

	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "cht"
	app.Version = Version
	app.Usage = ""
	app.Author = "takady"
	app.Email = "takadyuichi@gmail.com"
	app.Commands = Commands

	app.Run(os.Args)
}
