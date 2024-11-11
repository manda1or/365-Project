// Task Managing feature
package tasks

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type TaskStruct struct {
	TaskID          int
	TaskDescription string
	TaskComments    []string
	TaskIsComplete  bool
}

type TaskManager struct {
	Tasks []TaskStruct
}

func(tm *TaskManager) CreateNewTask() {
	var taskSlice []TaskStruct
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Task Management List")
	for {
		// Task Description
		fmt.Println("Enter a new task (Type 'exit' to quit.):")
		scanner.Scan()
		taskDescription := scanner.Text()
		if strings.ToLower(taskDescription) == "exit" {
			break
		}

		// Task Comments. Comma splitter
		fmt.Println("Enter comment:")
		scanner.Scan()
		taskCommentInput := scanner.Text()
		var taskComments []string
		if taskCommentInput != "" {
			taskComments = strings.Split(taskCommentInput, ",")
		}

		// Task completion status
		var taskIsComplete bool
		for {
			fmt.Println("Task Completion: Yes/No:")
			scanner.Scan()
			taskIsCompleteInput := strings.TrimSpace(strings.ToLower(scanner.Text()))
			if taskIsCompleteInput == "yes" {
				taskIsComplete = true
				break
			} else if taskIsCompleteInput == "no" {
				taskIsComplete = false
				break
			} else {
				fmt.Println("Invalid input. \"Yes\" or \"No\" only.")
			}
		}

		// Creates new task to slice
		newTask := TaskStruct{
			TaskID:          len(taskSlice) + 1,
			TaskDescription: taskDescription,
			TaskComments:    taskComments,
			TaskIsComplete:  taskIsComplete,
		}
		taskSlice = append(taskSlice, newTask)

		fmt.Println("New task added!")
	}
}
	func(tm *TaskManager) ShowTasks() {
	fmt.Println("All Tasks:")
	for _, task := range tm.Tasks {
		fmt.Printf("ID: %d, Description: %s, Complete: %v, Comments: %v\n",
			task.TaskID, task.TaskDescription, task.TaskComments, task.TaskIsComplete)
	}
}


