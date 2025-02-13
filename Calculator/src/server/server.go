package main

import (
	calc "Calculator/src"
	"bufio"
	"fmt"
	"net"
	"strconv"
	"strings"
	"sync"
)

var ioLk sync.Mutex

func printSync(s string) {
	defer ioLk.Unlock()
	ioLk.Lock()
	fmt.Println(s)
}

func server(con net.Conn) {
	var err error
	// Must close the connection
	defer func() {
		err = con.Close()
		if err != nil {
			panic(err)
		}
	}()
	// The calc for this session
	c := calc.Calculator{}
	// Reading from the connection
	reader := bufio.NewReader(con)

	var msg string

	for {
		// Get the message
		msg, err = reader.ReadString('\n')
		if err != nil {
			_, _ = con.Write([]byte(err.Error() + "\n"))
			return
		}
		// Transform the message
		msg = strings.TrimSpace(msg)
		argv := strings.Fields(msg)
		// Validate
		err = c.GetArguments(argv)
		if err != nil {
			_, _ = con.Write([]byte(err.Error() + "\n"))
			continue
		}
		// Operate
		ans, err := c.Operate()
		if err != nil {
			_, _ = con.Write([]byte(err.Error() + "\n"))
			continue
		}
		_, err = con.Write([]byte(strconv.FormatFloat(ans, 'f', -1, 64) + "\n"))
		if err != nil {
			printSync(err.Error())
			return
		}
	}
}

func main() {
	// Create the Listener
	skt, err := net.Listen("tcp", "127.0.0.1:1234")
	if err != nil {
		panic(err)
		return
	}

	defer func(skt net.Listener) {
		err := skt.Close()
		if err != nil {
			panic(err)
		}
	}(skt)

	fmt.Println("Listening on " + skt.Addr().String())

	for {
		con, err := skt.Accept()
		if err != nil {
			printSync(err.Error())
			return
		}

		printSync("New connection from " + con.RemoteAddr().String())

		go server(con)
	}
}
