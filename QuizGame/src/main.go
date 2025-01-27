package main

import (
	"QuizGame/src/funcs"
	"fmt"
	"os"
	"sync"
)

var (
	question uint8 = 0
	answer   uint8 = 1
)

var (
	wg        sync.WaitGroup
	mtx       sync.Mutex
	doneItems chan struct{} = make(chan struct{}, 1)
)

func main() {
	/*
		1. Read the questions
		2. Ask for answer
		3. Check input with answer on csv
	*/
	// 1.
	var file *os.File
	var err error
	file, err = os.Open("data/problems.csv")
	if err != nil {
		panic(err)
	}
	// Here we use slices since we don´t know the size of the CSV file
	var items []string
	// Using goroutine to improve performance, read and show at the same time
	wg.Add(2)
	go func() {
		// To send the end signal
		defer wg.Done()
		// Read
		for {
			items = funcs.ReadCSVLineToDisplay(file, &mtx)
			if items == nil {
				break
			}
		}
	}()
	// Using a goroutine to recive the items slice of strings to display
	go func() {
		// To send the end signal
		defer wg.Done()
		// Now we wait for the items to be filled by the previous goroutine
		for i, item := range items {
			fmt.Printf("Question #%d: %s\n", i+1, item[0])

		}
	}()
	// To end
	for i := 0; i < 2; i++ {
		wg.Wait()
	}
}
