package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	// Open a connection
	con, err := net.Dial("tcp", "127.0.0.1:1234")
	if err != nil {
		panic(err)
	}
	defer con.Close()

	reader := bufio.NewReader(os.Stdin)
	response := bufio.NewReader(con)

	fmt.Print("The calculator:\n usage: <FirstNumber> <operator> <SecondNumber>\n Operators: +  -  *  /  ^\n")

	for {
		fmt.Print("Enter your operation: ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		_, err := con.Write([]byte(input + "\n"))
		if err != nil {
			panic(err)
		}

		ans, _ := response.ReadString('\n')
		fmt.Print(ans + "\n")
	}

}
