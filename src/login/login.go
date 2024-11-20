package login

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// File to store login credentials
const credentialsFile = "src/login/credentials.txt"

// RunLoginSystem handles the login system loop
func RunLoginSystem() string {
    for {
        fmt.Println("\n--- Login System ---")
        fmt.Println("1. Register")
        fmt.Println("2. Login")
        fmt.Println("3. Back to Main Menu")
        fmt.Print("Choose an option: ")

        var choice int
        fmt.Scan(&choice)

        switch choice {
        case 1:
            register()
        case 2:
            username := login()
            if username != "" {
                return username
            }
        case 3:
            return ""
        default:
            fmt.Println("Invalid choice. Please try again.")
        }
    }
}

// Register a new user
func register() {
	var username, password string
	fmt.Print("Enter username: ")
	fmt.Scan(&username)
	fmt.Print("Enter password: ")
	fmt.Scan(&password)

	if usernameExists(username) {
		fmt.Println("Username already exists. Please try a different one.")
		return
	}

	err := saveCredentials(username, password)
	if err != nil {
		fmt.Println("Error saving credentials:", err)
		return
	}

	fmt.Println("Registration successful!")
}

// Log in an existing user
func login() string {
    var username, password string
    fmt.Print("Enter username: ")
    fmt.Scan(&username)
    fmt.Print("Enter password: ")
    fmt.Scan(&password)

    if validateCredentials(username, password) {
        fmt.Println("Login successful! Welcome,", username)
        return username
    }
    fmt.Println("Invalid username or password.")
    return ""
}

// Save credentials to file
func saveCredentials(username, password string) error {
	dir :=  filepath.Dir(credentialsFile)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.MkdirAll(dir, 0755)
	}
	file, err := os.OpenFile(credentialsFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(fmt.Sprintf("%s:%s\n", username, password))
	return err
}

// Check if username exists
func usernameExists(username string) bool {
	file, err := os.Open(credentialsFile)
	if err != nil {
		return false
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		credentials := strings.Split(scanner.Text(), ":")
		if len(credentials) > 0 && credentials[0] == username {
			return true
		}
	}
	return false
}

// Validate username and password
func validateCredentials(username, password string) bool {
	file, err := os.Open(credentialsFile)
	if err != nil {
		return false
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		credentials := strings.Split(scanner.Text(), ":")
		if len(credentials) == 2 && credentials[0] == username && credentials[1] == password {
			return true
		}
	}
	return false
}
