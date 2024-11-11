package main

import (
	"365-Project/src/reviews"
	"365-Project/src/tasks"
	"fmt"
)

func main() {

	//Task Manager instance
	tm := tasks.TaskManager{}
	// Task Creation
	tm.CreateNewTask()
	// Displays All Tasks
	tm.ShowTasks()

	// Review Code
	review := reviews.CreateNewReview()
	fmt.Println("Review List:")
	fmt.Printf("Name: %s\n", review.ReviewerName)
	fmt.Printf("Comment: %s\n", review.ReviewComment)
	fmt.Printf("Rating: %d\n", review.ReviewRating)
}
