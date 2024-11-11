package main

import (
	//"fmt"
	"365-Project/src/tasks"
)

func main() {

	//Task Manager instance
	tm := tasks.TaskManager{}

	// Task Creation
	tm.CreateNewTask()

	// Displays All Tasks
	tm.ShowTasks()
}
