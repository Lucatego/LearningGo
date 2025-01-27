package main

import (
	"QuizGame/src/funcs"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

const (
	question uint8 = 0
	answer   uint8 = 1
)

var (
	wg           sync.WaitGroup
	qNa          chan []string = make(chan []string, 5)
	end          chan struct{} = make(chan struct{}, 1)
	displayMutex sync.Mutex
)

func server() {
	// 1.
	var file *os.File
	var err error
	file, err = os.Open("data/problems.csv")
	if err != nil {
		end <- struct{}{}
		panic(err)
	}
}

func client() {

}

func main() {
	/*
		1. Read the questions
		2. Ask for answer
		3. Check input with answer on csv
	*/
	var file *os.File
	var err errorz
	file, err = os.Open("data/problems.csv")
	if err != nil {
		end <- struct{}{}
		panic(err)
	}
	// Using goroutines to improve performance, read and show at the same time
	wg.Add(2)
	go func() {
		// Here we use slices since we don´t know the size of the CSV file
		var items []string
		var reader *bufio.Scanner = bufio.NewScanner(file)
		// To send the end signal
		defer wg.Done()
		// Loop for all the CSV file
		for {
			// Get a token
			items = funcs.ReadCSVLineToDisplay(reader)
			// End case
			if items == nil {
				displayMutex.Lock()
				fmt.Println("End of reading")
				displayMutex.Unlock()
				qNa <- nil
				return
			}
			// Send it to the channel
			select {
			case qNa <- items:
				displayMutex.Lock()
				fmt.Println("Question send to client")
				displayMutex.Unlock()
			default:
				displayMutex.Lock()
				fmt.Println("Waiting for client to free queue")
				displayMutex.Unlock()
				time.Sleep(5 * time.Second)
			}
		}
	}()

	// Using a goroutine to recive the items slice of strings to display
	var correct, total uint32 = 0, 0
	go func() {
		var items []string
		var correctAnswer int
		var userAnswer int
		// To send the end signal
		defer wg.Done()
		// Now we wait for the items to be filled by the previous goroutine
		for {
			select {
			case items = <-qNa:
				if items == nil {
					displayMutex.Lock()
					fmt.Println("End of the quz")
					displayMutex.Unlock()
					return
				}
				// Question recieved
				total++
				displayMutex.Lock()
				fmt.Printf("Question #%s:\n", items[question])
				correctAnswer, _ = strconv.Atoi(items[answer]) // TODO: Handle this error
				// Read from user
				fmt.Printf("Your answer is: ")
				fmt.Scanf("%d", &userAnswer)
				if userAnswer == correctAnswer {
					correct++
					fmt.Println("Correct!")
				} else {
					fmt.Printf("The correct answer is: %d\n", correctAnswer)
				}
				displayMutex.Unlock()
			case <-end:
				displayMutex.Lock()
				fmt.Println("Error from server")
				displayMutex.Unlock()
			default:
				displayMutex.Lock()
				fmt.Println("Waiting for a question from \"server\"")
				displayMutex.Unlock()
				time.Sleep(2 * time.Second)
			}
		}
	}()
	// To end
	for i := 0; i < 2; i++ {
		wg.Wait()
	}
	close(qNa)
	// End
	fmt.Printf("Results: %d correct of %d questions\n", correct, total)
}
