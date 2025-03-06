package server

import (
	"fmt"
	"net"
	"sync"
)

const (
	protocol             string = "tcp"
	queueMaxSize         uint32 = 1024
	numberOfHandlers     uint8  = 8
	maxClientsPerHandler uint32 = 64
)

var (
	// To safely stop the server
	wg   sync.WaitGroup
	stop chan struct{} = make(chan struct{}, 1)
	// To manage connections
	clients     chan net.Conn = make(chan net.Conn, queueMaxSize)
	fullHandler chan uint8    = make(chan uint8, numberOfHandlers)
	// To print on server console without interfering with the user input
	output  chan string = make(chan string, queueMaxSize)
	printMx sync.Mutex
)

func server(ipAddress string, port string) {

	var runServer bool = true

	// Create the socket
	skt, err := net.Listen(protocol, ipAddress+":"+port)
	if err != nil {
		panic(err)
		return
	}
	fmt.Printf("Server listening on %s\n", skt.Addr().String())

	// TODO: Console input
	// go ConsoleInput()

	// Handle the client in other process
	go DistributeClients()

	// Start listening for clients
	for runServer {
		// Listen
		select {
		case <-stop:
			runServer = false
		default:
			// Accept a client and validate
			conn, err := skt.Accept()
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Printf("New connection from %s\n", conn.RemoteAddr().String())

			// Send the connection to a channel
			clients <- conn
		}
	}

	// Finish
	wg.Wait()
	skt.Close()
}

// ConsoleInput()
// works as an input for the server. It handles commands like "stop" to stop accepting connections.
func ConsoleInput() {
	// Open a terminal
	var input string
	for {
		fmt.Printf("> ")
		fmt.Scanln(&input)
	}
}

func SafePrint(output string) {
	printMx.Lock()
	fmt.Printf("%s", output)
	printMx.Unlock()
}

func DistributeClients() {
	// Here we will receive the connections from the channel to evaluate and distribute them.
	var handlerListening [numberOfHandlers]bool
	activeHandlers := numberOfHandlers
	// Forever loop
	for i := uint8(0); i < numberOfHandlers; i++ {
		go HandleClients(i)
		handlerListening[i] = true
	}

	// Status control
	for {
		if activeHandlers == 0 {
			break
		}
		select {
		case id := <-fullHandler:
			handlerListening[id] = false
			activeHandlers--
		}
	}
}

func HandleClients(id uint8) {
	numberOfClients := uint32(0)

	// Welcome
	fmt.Println("Server id: %d\n", id)

	// For the moment, if the server is full, it breaks (Temporal solution).
	// In a real solution, the DistributeClients() function must create more handlers
	for {
		select {
		case conn := <-clients:
			// If the handler is full, then give to another
			if numberOfClients == maxClientsPerHandler {
				clients <- conn
				fullHandler <- id
				break
			}

			ServeClient(conn)
		}
	}
}

func ServeClient(conn net.Conn) {
	// Look in the database
}
