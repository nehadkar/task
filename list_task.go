//ListTask is responsible for handling the task list subcommands 
//Implements the Run, Help and Synopsis functions
package main

import (
	"fmt"
	"github.com/mitchellh/cli"
	"task/database"
)

type ListTask struct {
	ui cli.Ui 
	db *database.DBSession 
}

func (c *ListTask) Run(_ []string) int {
	c.ui.Output("You have the following tasks:")
	//Call database's Gettasks method
	Tasks, err := c.db.GetTasks()
	if err != nil {
		str := fmt.Sprintf("Error getting tasks from db, %v", err)
		c.ui.Output(str)
	}
	for i,v := range(Tasks){
		str := fmt.Sprintf("%d. %s", i+1, v.Tsk)
		c.ui.Output(str)
	}
	return 0
}

func (c *ListTask) Help() string {
	return "List all your TODO tasks"
}

func (c *ListTask) Synopsis() string {
	return "Lists the ToDOs"
}
