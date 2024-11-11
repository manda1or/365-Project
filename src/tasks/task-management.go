// Task Managing feature
package tasks

import(
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type taskStruct struct {
	TaskID int
	TaskDescription string
	TaskComments []string
	TaskIsComplete bool
}

func createNewTask() {
	var taskSlice []taskStruct 
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
	}

	// Task Comments. Comma splitter
	fmt.Println("Enter comment:")
	scanner.Scan()
	taskCommentInput := scanner.Text()
	taskComments := strings.Split(taskCommentInput, ",")

	// Task completion status
	var taskIsComplete bool
	for{
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
	 newTask := taskStruct {
		TaskID: len(taskSlice) +1,
		TaskDescription: taskDescription,
		TaskComments: taskCommentInput,
		TaskIsComplete: taskIsComplete,
	}
	tasks := append(taskSlice, newTask)
	fmt.Println("New task added!")
}
