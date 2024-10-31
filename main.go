package main

import (
	"fmt"
	"os"
)

func main() {
	checkArgs(os.Args[1:])

	filename := os.Args[1]
	fmt.Println(filename)

	// If existing test execution, set testExecution key
	// retrieve value from command line args

	// Else, add required information to data
	getTestExecutionData()

	// for test case in json file
	// get key, and overall status (pass/fail)

	// --------------------------
	// f, err := os.Open(filename)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	os.Exit(1)
	// }

	// fileBytes := make([]byte, 9999)
	// fileText, err := f.Read(fileBytes)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	os.Exit(1)
	// }

	// fmt.Printf("%v bytes: %v\n", fileText, string(fileBytes[:fileText]))
}

// checkArgs should validate command line args and display helpful error messages where possible
func checkArgs(args []string) {
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

func getTestExecutionData() {

}
