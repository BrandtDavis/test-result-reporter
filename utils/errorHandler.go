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
}
