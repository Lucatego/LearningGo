package test

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func OpenReadFileTest() {
	// The os.Getwd() stands for working directory (I guess)
	var basePath string
	var err error
	basePath, err = os.Getwd()
	if err != nil {
		panic(err)
	}
	fmt.Println("base path:", basePath)
	/*
		start
	*/
	// The os.Open() function returns 2 values, the *os.File and the error type.
	var file *os.File
	var ferr error
	file, ferr = os.Open("data/problems.csv")
	// Error handling
	if ferr != nil {
		panic(ferr)
	}
	// Create a buffer
	var scanner *bufio.Scanner
	scanner = bufio.NewScanner(file)
	// Loop to print all the tokens (lines) in the file
	var counter uint64 = 1
	// Remember that Go has a garbage collector, so it's not necessary to delete the array of strings
	for scanner.Scan() {
		// Use the scanner to read a line from the csv file
		line := scanner.Text()
		// Now use a splitter to create an array of strings ([]string) that represents every item on a token
		items := strings.Split(line, ",")
		fmt.Println("*----------------*")
		fmt.Printf("Question #%d: %s\n\tAnswer: %s\n", counter, items[0], items[1])
		counter++
	}
}
