package main

import (
	"365-Project/src/costestimate"
	"365-Project/src/login"
	"365-Project/src/renovproject"
	"365-Project/src/reviews"
	"fmt"
)

func main() {
	var loggedInUser string
	for {
		fmt.Println("Welcome to Jauregui Renovations!")
		fmt.Println("1. Login System")
		fmt.Println("2. Flooring Cost Estimator")
		fmt.Println("3. Renovation Projects")
		fmt.Println("4. Reviews")
		fmt.Println("5. Exit")
		fmt.Print("Choose an option: ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			// Call the login system
			loggedInUser = login.RunLoginSystem()
		case 2:
			// Call cost estimate
			costestimate.RunCostEstimator()
		case 3:
			if loggedInUser == "" {
				fmt.Println("You must be logged in to access renovation projects.")
			} else {
				renovproject.RunRenovProject(loggedInUser)
			}
		case 4:
			review := reviews.CreateNewReview()
			fmt.Println("\nReview Summary: ")
			fmt.Printf("Reviewer: %s\n", review.ReviewerName)
			fmt.Printf("Comment: %s\n:", review.ReviewComment)
			fmt.Printf("Rating: %d/5\n", review.ReviewRating)
		case 5:
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
