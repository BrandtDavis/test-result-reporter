package utils

import (
	"fmt"
	"os"
)

// checkArgs should validate command line args and display helpful error messages where possible
func CheckArgs(args []string) {
	// Check if there are enough args
	if len(args) < 2 {
		fmt.Println("You must provide an input and an output file name as arguments")
		os.Exit(1)
	}

	// Check that there are not too many args
	if len(args) > 2 {
		fmt.Println("You must provide only 2 arguments: an input and an output file")
		os.Exit(1)
	}

	// Check if first arg is an existing file
	fileInfo, err := os.Stat(args[0])
	if err != nil {
		fmt.Println("Error on first arg:", err)
		os.Exit(1)
	}
	fmt.Println(fileInfo)
	// if args[1]
}
