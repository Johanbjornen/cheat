package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

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

func doAdd(c *cli.Context) {
	args := c.Args()
	add_command := args.First()
	add_description := args[:2]
	commands := getAll()

	fmt.Println(add_command)
	fmt.Println(add_description)

	for _, add_task := range add_command {
		commands = append(commands, add_task)
	}

	write(commands)
}

func doList(c *cli.Context) {
	commands := getAll()
	for i, command := range commands {
		println(i+1, ":", string(command))
	}
}

func doDelete(c *cli.Context) {
	delete_command := c.Args()[0]
	commands := getAll()
	var result []string

	for _, command := range commands {
		if command != delete_command {
			result = append(result, command)
		}
	}

	write(result)
}

func getAll() []string {
	f, err := os.OpenFile(CHEAT_FILE, os.O_RDONLY, 0600)
	assert(err)
	defer f.Close()

	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	assert(scanner.Err())

	return lines
}

func write(commands []string) {
	f, err := os.OpenFile(CHEAT_FILE, os.O_TRUNC|os.O_WRONLY, 0600)
	assert(err)
	defer f.Close()

	_, err = f.WriteString(strings.Join(commands, "\n"))
	assert(err)
}
