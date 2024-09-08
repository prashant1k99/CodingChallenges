package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func IsValidJSON(content string) bool {
	// If the content either starts with [ or {, then it's a valid Json
	if len(content) == 0 {
		return false
	}
	if string(content[0]) == "{" || string(content[0]) == "[" {
		return true
	}
	return false
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	content, err := io.ReadAll(reader)
	if err != nil {
		fmt.Println("Unable to get the Stdin, Err:", err)
		os.Exit(1)
	}

	// Check for valid opening of content
	isValid := IsValidJSON(string(content))
	if !isValid {
		fmt.Println("Invalid JSON is passed")
		os.Exit(1)
	}
	fmt.Println(string(content))
	os.Exit(0)
}
