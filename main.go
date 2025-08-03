package main

import (
	"fmt"
	"os"
)

func runApp() {
	const appName = "URL Shortener CLI"
	var alias string
	fullURL := ""

	fmt.Println(appName)
	fmt.Print("Please enter an alias for your URL: ")
	_, _ = fmt.Scanln(&alias)

	fmt.Print("Please enter the full URL: ")
	_, _ = fmt.Scanln(&fullURL)

	fmt.Printf("Saved URL for %s: %s\n", alias, fullURL)
}

func main() {

	if len(os.Args) < 2 {
		// No arguments provided, run the interactive app
		runApp()
		return
	}

	command := os.Args[1]

	switch command {
	case "add":
		fmt.Println("Add command received.")
	default:
		fmt.Printf("Unknown command: %s\n", command)
	}
}
