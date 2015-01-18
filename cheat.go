package main

import (
	"os"

	"github.com/codegangsta/cli"
)

var CHEAT_FILE = os.Getenv("HOME") + "/.cheat"

func main() {
	app := cli.NewApp()
	app.Name = "cheat"
	app.Version = Version
	app.Usage = ""
	app.Author = "takady"
	app.Email = "takadyuichi@gmail.com"
	app.Commands = Commands

	app.Run(os.Args)
}
