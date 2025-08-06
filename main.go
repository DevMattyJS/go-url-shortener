package main

import (
	"fmt"
	"os"
)

var urlMappings = make(map[string]URLMapping)

type URLMapping struct {
	Alias   string
	FullURL string
}

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
	// Adds a new entry to urlMappings map if all parameters are received
	case "add":

		if len(os.Args) < 4 {
			fmt.Println("Error: Missing paramaters! Usage: add <alias> <url>")
			return
		}

		alias := os.Args[2]
		fullURL := os.Args[3]
		newEntry := URLMapping{Alias: alias, FullURL: fullURL}
		urlMappings[alias] = newEntry
		fmt.Printf("Added new URL mapping: %s -> %s\n", alias, fullURL)

	// Retrieves an URL stored for an alias if all parameters are received
	case "get":
		if len(os.Args) < 3 {
			fmt.Println("Error: Missing parameters! Usage: get <alias>")
			return
		}

		alias := os.Args[2]
		mapping, ok := urlMappings[alias]
		if ok {
			fmt.Printf("The full url for alias %s: %s", alias, mapping.FullURL)
		} else {
			fmt.Printf("Alias %s not found!", alias)
		}
	default:
		fmt.Printf("Unknown command: %s\n", command)
	}
}
