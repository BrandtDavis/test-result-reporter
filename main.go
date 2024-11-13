package main

import (
	"fmt"
	"os"

	"jsonParser/utils"
)

func main() {
	utils.CheckArgs(os.Args[1:])

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

func getTestExecutionData() {

}
