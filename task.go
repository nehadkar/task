//This is a command line task manager
package main

import (
	"os"
	"github.com/mitchellh/cli"
	"task/database"

	"bytes"
	"fmt"
	"sort"
	"strings"
)

func main() {
	//Initialize the database
	db := database.DBSession{}
	db.InitDB("localhost")


	//initialize and run the CLI
	c := cli.NewCLI("task", "1.0.0")
	c.Args = os.Args[1:]
	c.HelpFunc = taskHelpFunc

	ui := &cli.BasicUi{
		Reader:      os.Stdin,
		Writer:      os.Stdout,
		ErrorWriter: os.Stderr,
	}

	//Add the list, add and do commands to the cli 
	c.Commands = map[string]cli.CommandFactory{
		"list": func() (cli.Command, error) {
			return &ListTask{ui,&db},nil
		},
		"add": func() (cli.Command, error) {
			return &AddTask{ui,&db},nil
		},
		"do": func() (cli.Command, error) {
			return &DoTask{ui,&db},nil
		},
	}

	exitStatus, err := c.Run()
	if err != nil {
		fmt.Printf("%v",err)
	}

	os.Exit(exitStatus)
}

//CLI custom helper function 
func taskHelpFunc(commands map[string]cli.CommandFactory) string {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("task is a CLI task manager\n\n"))
	buf.WriteString(fmt.Sprintf("Usage: task [command]\n\n"))
	buf.WriteString("Available Commands:\n")

	// Get the list of keys so we can sort them, and also get the maximum
	// key length so they can be aligned properly.
	keys := make([]string, 0, len(commands))
	maxKeyLen := 0
	for key := range commands {
		if len(key) > maxKeyLen {
			maxKeyLen = len(key)
		}

		keys = append(keys, key)
	}
	sort.Strings(keys)

	for _, key := range keys {
		commandFunc, ok := commands[key]
		if !ok {
			// This should never happen since we JUST built the list of
			// keys.
			panic("command not found: " + key)
		}

		command, err := commandFunc()
		if err != nil {
			continue
		}

		key = fmt.Sprintf("%s%s", key, strings.Repeat(" ", maxKeyLen-len(key)))
		buf.WriteString(fmt.Sprintf("    %s    %s\n", key, command.Synopsis()))
	}

	buf.WriteString(fmt.Sprintf("Flags:\n"))
	buf.WriteString(fmt.Sprintf("    -h, --help       help for task\n"))
	buf.WriteString(fmt.Sprintf("    -v, --version    task version\n"))
	buf.WriteString(fmt.Sprintf("\nUse \"task [command] --help\" for more information about a command\n"))
	return buf.String()
}

