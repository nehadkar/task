//DoTask is responsible for handling the task do subcommands 
//Implements the Run, Help and Synopsis functions
package main

import (
	"fmt"
	"github.com/mitchellh/cli"
	"task/database"
	"strconv"
)

type DoTask struct {
	ui cli.Ui 
	db *database.DBSession 
}

func (c *DoTask) Run(args []string) int {
	taskId,_ := strconv.ParseInt(args[0],10,32)
	taskId  = taskId - 1
	
	//Call database's gettask method
	//Retrieve the id of the task to be marked done/ deleted
	//Mark it completed
	Tasks, err := c.db.GetTasks()
	if err != nil {
		str := fmt.Sprintf("Error getting tasks from db, %v", err)
		c.ui.Output(str)
	}
	task := Tasks[taskId]
	err = c.db.DoTask(task)
	if err != nil {
		str := fmt.Sprintf("Error deleting task from db, %v", err)
		c.ui.Output(str)
		return 0
	}
	str := fmt.Sprintf("You have completed the \"%s\" task", task.Tsk)
	c.ui.Output(str)
	return 0
}

func (c *DoTask) Help() string {
	return "Complete a task from your TODO list"
}

func (c *DoTask) Synopsis() string {
	return "Complete ToDOs"
}
