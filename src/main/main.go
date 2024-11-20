package main

import (
	"fmt"
	"src/login" // Import the login package
)

func main() {
	for {
		fmt.Println("Welcome to the Main Application")
		fmt.Println("1. Login System")
		fmt.Println("2. Other Tasks")
		fmt.Println("3. Exit")
		fmt.Print("Choose an option: ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			// Call the login system
			login.RunLoginSystem()
		case 2:
			// Placeholder for other tasks
			fmt.Println("Performing other tasks...")
		case 3:
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

