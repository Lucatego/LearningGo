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
	// To send the end signal
	defer wg.Done()
	// 1.
	var file *os.File
	var err error
	file, err = os.Open("data/problems.csv")
	if err != nil {
		end <- struct{}{}
		panic(err)
	}
	// Here we use slices since we don´t know the size of the CSV file
	var items []string
	var reader *bufio.Scanner = bufio.NewScanner(file)
	// Loop for all the CSV file
	for {
		// Get a token (line)
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
}

func client(correct *uint32, total *uint32) {
	// To send the end signal
	defer wg.Done()
	// To read from server and user
	var items []string
	var correctAnswer, userAnswer int
	// 2.
	for {
		select {
		case items = <-qNa:
			if items == nil {
				displayMutex.Lock()
				fmt.Println("End of the quiz")
				displayMutex.Unlock()
				return
			}
			// Question recieved
			(*total)++
			displayMutex.Lock()
			fmt.Printf("Question #%s:\n", items[question])
			correctAnswer, _ = strconv.Atoi(items[answer]) // TODO: Handle this error
			// 3.
			fmt.Printf("Your answer is: ")
			fmt.Scanln(&userAnswer)
			if userAnswer == correctAnswer {
				(*correct)++
				fmt.Println("Correct!")
			} else {
				fmt.Printf("The correct answer is: %d\n", correctAnswer)
			}
			displayMutex.Unlock()
		case <-end:
			fmt.Println("Fatal error from server")
			return
		default:
			displayMutex.Lock()
			fmt.Println("Waiting for a question from \"server\"")
			displayMutex.Unlock()
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	/*
		1. Read the questions
		2. Ask for answer
		3. Check input with answer on csv
	*/
	// Using goroutines to improve performance, read and show at the same time
	wg.Add(2)
	go server()
	// Using a goroutine to recive the items slice of strings to display
	var correct, total uint32 = 0, 0
	go client(&correct, &total)
	// To end
	for i := 0; i < 2; i++ {
		wg.Wait()
	}
	close(qNa)
	// End
	fmt.Printf("Results: %d correct of %d questions\n", correct, total)
}
