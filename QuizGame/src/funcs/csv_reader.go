package funcs

import (
	"bufio"
	"os"
	"strings"
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

func ReadCSVLineToDisplay(reader *bufio.Scanner) []string {
	var items []string = nil
	// Scan for tokens
	if reader.Scan() {
		items = strings.Split(reader.Text(), ",")
	}
	return items
}
