//AddTask is responsible for handling the task add subcommands 
//Implements the Run, Help and Synopsis functions
package main

import (
	"fmt"
	"github.com/mitchellh/cli"
	"task/database"
	"task/model"
	"strings"
)

type AddTask struct {
	ui cli.Ui 
	db *database.DBSession 
}

func (c *AddTask) Run(args []string) int {
	task := model.Task{Tsk:strings.Join(args, " ")}
	err := c.db.AddTask(task)
	if err != nil {
		str := fmt.Sprintf("Error adding task to db, %v", err)
		c.ui.Output(str)
		return 0
	}
	str := fmt.Sprintf("Added \"%s\" to your task list", task.Tsk)
	c.ui.Output(str)
	return 0
}

func (c *AddTask) Help() string {
	return "Add a task to your TODO list"
}

func (c *AddTask) Synopsis() string {
	return "Add ToDOs"
}
