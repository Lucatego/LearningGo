package main

/*
	Important: In Go, slices, maps, interfaces are implicit pointers. On the other hand, small data types like uint8,
	float32 or struct are copied.

	Be careful with structs that have pointers like string:

	type string struct {
		data *byte  // Pointer
		len  int
	}

	In this case, by passing string as an argument, it copies everything, so if the data is modified, it will only
	affect to that copy of the pointer but not the original one.
	Remember that the attributes of the string are immutable because of being private.
*/

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

func printHeader() {
	fmt.Println("-----------------------------------")
	fmt.Printf("%20s\n", "The calculator")
	fmt.Printf("usage: <FirstNumber> <operator> <SecondNumber>\n")
	fmt.Print("Operators: +  -  *  /  ^\n")
}

func main() {
	var err error = nil
	var con net.Conn
	// Open a connection
	for {
		fmt.Println("Trying to connect...")
		con, err = net.Dial("tcp", "127.0.0.1:1234")
		if err != nil {
			fmt.Println("Error connecting: ", err)
			fmt.Println("Trying again in 5 seconds...")
			time.Sleep(5 * time.Second)
		} else {
			fmt.Println("Connection successful...")
			break
		}
	}
	defer func() {
		err = con.Close()
		if err != nil {
			panic(err)
		}
	}()

	// Readers for user input and server output
	reader := bufio.NewReader(os.Stdin)
	response := bufio.NewReader(con)

	printHeader()

	for {
		fmt.Print("Enter your operation: ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		_, err := con.Write([]byte(input + "\n"))
		if err != nil {
			panic(err)
		}

		ans, _ := response.ReadString('\n')
		fmt.Print("Answer: " + ans + "\n")
	}

}
