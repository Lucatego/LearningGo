package funcs

import (
	"bufio"
	"os"
	"strings"
	"sync"
)

func ReadCSV(file *os.File) [][]string {
	var reader *bufio.Scanner = bufio.NewScanner(file)
	var items []string
	var csv [][]string = nil
	// Read the CSV file by tokens (lines)
	for reader.Scan() {
		// Here we split the token by the comas
		items = strings.Split(reader.Text(), ",")
		// And here we append the slice of strings (csv) to the final slice of []string
		csv = append(csv, items)
	}
	return csv
}

func ReadCSVLineToDisplay(file *os.File, mutex *sync.Mutex) []string {
	var reader *bufio.Scanner = bufio.NewScanner(file)
	var items []string = nil
	// Scan for tokens
	for reader.Scan() {
		
	}
	return items
}
