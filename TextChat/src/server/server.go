package server

import (
	"fmt"
	"net"
	"sync"
)

const (
	protocol       string = "tcp"
	queue_max_size uint32 = 1024
)

var (
	// To safely stop the server
	wg   sync.WaitGroup
	stop chan struct{} = make(chan struct{})
	// To manage connections
	clients chan net.Conn = make(chan net.Conn, queue_max_size)
	// To print on server console without interfering with the user input
	output   chan string = make(chan string, queue_max_size)
	print_mx sync.Mutex
)

func server(ip_address string, port string) {

	var run_server bool = true

	// Create the socket
	skt, err := net.Listen(protocol, ip_address+":"+port)
	if err != nil {
		panic(err)
		return
	}
	fmt.Printf("Server listening on %s\n", skt.Addr().String())

	// Console input
	go server_console_input()

	// Start listening for clients
	for run_server {
		// Listen
		select {
		case <-stop:
			run_server = false
			stop <- struct{}{}
		default:
			// Accept a client and validate
			conn, err := skt.Accept()
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Printf("New connection from %s\n", conn.RemoteAddr().String())

			// Handle the client
			go handle_client(conn)
		}
	}

	// Finish
	wg.Wait()
	skt.Close()
}

/*
 * This function works as an input for the server.
 * It handles commands like "stop" to stop accepting connections.
 */
func server_console_input() {
	// Open a terminal
	var input string
	for {
		fmt.Printf("> ")
		fmt.Scanln(&input)
	}
}

func safe_print() {
	print_mx.Lock()

	print_mx.Unlock()
}

func handle_client(conn net.Conn) {
	for {
		var conn net.Conn = nil

		select {
		case conn = <-clients:
			// To sync
			wg.Add(1)
			defer wg.Done()

		default:

		}
	}
}
