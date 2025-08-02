package main

import "fmt"

func main() {

	const appName = "URL Shortener CLI"
	var alias string
	fullURL := ""

	fmt.Println(appName)
	fmt.Print("Please enter an alias for your URL:")
	_, _ = fmt.Scanln(&alias)

	fmt.Print("Please enter the full URL:")
	_, _ = fmt.Scanln(&fullURL)

	fmt.Printf("Saved URL for %s: %s\n", alias, fullURL)
}
