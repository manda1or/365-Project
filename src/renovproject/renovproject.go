package renovproject

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

// RunRenovProject handles renovation project creation and report generation
func RunRenovProject(username string) {
    fmt.Println("\n--- Renovation Projects ---")
    fmt.Println("1. Create New Project")
    fmt.Println("2. Create Project Report")  // New option for creating a project report
    fmt.Println("3. Back to Main Menu")
    fmt.Print("Choose an option: ")

    var choice int
    fmt.Scan(&choice)

    switch choice {
    case 1:
        createProject(username)
    case 2:
        createProjectReport(username)
    case 3:
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

// createProjectReport allows the user to create a report for an existing project
func createProjectReport(username string) {
reader := bufio.NewReader(os.Stdin)

fmt.Print("Enter the project name for the report: ")
projectName, _ := reader.ReadString('\n')
projectName = strings.TrimSpace(projectName)

fmt.Print("Enter the report date (YYYY-MM-DD): ")
reportDate, _ := reader.ReadString('\n')
reportDate = strings.TrimSpace(reportDate)

fmt.Println("Enter your comments for the report:")
reportComment, _ := reader.ReadString('\n')
reportComment = strings.TrimSpace(reportComment)

// Ensure the directory exists for the report
reportDir := "src/renovproject/reports"
err := os.MkdirAll(reportDir, os.ModePerm)
if err != nil {
	fmt.Println("Error creating report directory:", err)
	return
}

// Create the report file
reportFilename := fmt.Sprintf("%s/%s_report_%s.txt", reportDir, projectName, reportDate)
file, err := os.Create(reportFilename)
if err != nil {
	fmt.Println("Error creating project report file:", err)
	return
}
defer file.Close()

    reportContent := fmt.Sprintf(
        "Project: %s\nReport Date: %s\nReported By: %s\nComments: %s\n",
        projectName, reportDate, username, reportComment,
    )

    _, err = file.WriteString(reportContent)
    if err != nil {
        fmt.Println("Error writing project report:", err)
        return
    }

    fmt.Printf("Project report for '%s' on %s has been successfully created and saved to '%s'.\n", projectName, reportDate, reportFilename)
}