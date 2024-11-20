package renovproject

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

// RunRenovProject handles renovation project creation
func RunRenovProject(username string) {
    fmt.Println("\n--- Renovation Projects ---")
    fmt.Println("1. Create New Project")
    fmt.Println("2. Back to Main Menu")
    fmt.Print("Choose an option: ")

    var choice int
    fmt.Scan(&choice)

    switch choice {
    case 1:
        createProject(username)
    case 2:
        return
    default:
        fmt.Println("Invalid choice. Please try again.")
    }
}

func createProject(username string) {
    var projectName, budget, deadline, comment string

    fmt.Print("Enter project name: ")
    fmt.Scan(&projectName)
    projectName = strings.TrimSpace(projectName)

    fmt.Print("Enter budget (in USD): ")
    fmt.Scan(&budget)

    fmt.Print("Enter deadline (e.g., YYYY-MM-DD): ")
    fmt.Scan(&deadline)

    fmt.Println("Enter additional comments:")
    scanner := bufio.NewScanner(os.Stdin)
    if scanner.Scan() {
        comment = scanner.Text()
    }

    // Ensure the directory exists
    projectDir := "src/renovproject"
    err := os.MkdirAll(projectDir, os.ModePerm)
    if err != nil {
        fmt.Println("Error creating project directory:", err)
        return
    }

    // Create the project file
    filename := fmt.Sprintf("%s/%s.txt", projectDir, projectName)
    file, err := os.Create(filename)
    if err != nil {
        fmt.Println("Error creating project file:", err)
        return
    }
    defer file.Close()

    projectDetails := fmt.Sprintf(
        "Project Name: %s\nCreated By: %s\nBudget: $%s\nDeadline: %s\nComments: %s\n",
        projectName, username, budget, deadline, comment,
    )

    _, err = file.WriteString(projectDetails)
    if err != nil {
        fmt.Println("Error saving project details:", err)
        return
    }

    fmt.Printf("Project '%s' has been successfully created and saved to '%s'.\n", projectName, filename)
}