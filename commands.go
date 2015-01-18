package main

import (
	"encoding/json"
	"log"
	"os"
	"strconv"

	"github.com/codegangsta/cli"
)

var Commands = []cli.Command{
	commandAdd,
	commandList,
	commandDelete,
}

var commandAdd = cli.Command{
	Name:  "add",
	Usage: "add command",
	Description: `
`,
	Action: doAdd,
}

var commandList = cli.Command{
	Name:  "list",
	Usage: "show cheat sheet",
	Description: `
`,
	Action: doList,
}

var commandDelete = cli.Command{
	Name:  "delete",
	Usage: "delete command",
	Description: `
`,
	Action: doDelete,
}

func debug(v ...interface{}) {
	if os.Getenv("DEBUG") != "" {
		log.Println(v...)
	}
}

func assert(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	_, err := os.Stat(CHEAT_FILE)
	if err != nil {
		_, err := os.Create(CHEAT_FILE)
		if err != nil {
			panic(err)
		}
	}
}

type Cheat struct {
	Command     string
	Description string
}

func doAdd(c *cli.Context) {
	args := c.Args()
	command := args.Get(0)
	description := args.Get(1)
	cheats := getAll()

	cheat := Cheat{command, description}
	cheats = append(cheats, cheat)
	write(cheats)
}

func doList(c *cli.Context) {
	cheats := getAll()
	for i, cheat := range cheats {
		println(i+1, ")", cheat.Command)
		println("   ", cheat.Description)
	}
}

func doDelete(c *cli.Context) {
	delete_index, err := strconv.Atoi(c.Args().Get(0))
	if err != nil {
		panic(err)
	}
	cheats := getAll()

	var result []Cheat
	for i, cheat := range cheats {
		if i+1 != delete_index {
			result = append(result, cheat)
		}
	}

	write(result)
}

func getAll() []Cheat {
	f, err := os.OpenFile(CHEAT_FILE, os.O_RDONLY, 0600)
	assert(err)
	defer f.Close()

	var cheat []Cheat
	j := json.NewDecoder(f)
	j.Decode(&cheat)

	return cheat
}

func write(cheat []Cheat) {
	f, err := os.OpenFile(CHEAT_FILE, os.O_TRUNC|os.O_WRONLY, 0600)
	assert(err)
	defer f.Close()

	j, err := json.Marshal(cheat)
	assert(err)

	_, err = f.Write(j)
	assert(err)
}
