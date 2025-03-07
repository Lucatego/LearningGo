package server

import (
	"fmt"
	"net"
	"sync"
)

const (
	protocol             string = "tcp"
	queueMaxSize         uint32 = 1024
	numberOfHandlers     uint32 = 32
	maxClientsPerHandler uint32 = 32
)

var (
	// To safely stop the server
	wg   sync.WaitGroup
	stop chan struct{} = make(chan struct{}, 1)

	// To manage connections
	clients chan net.Conn = make(chan net.Conn, queueMaxSize)

	// To print on server console without interfering with the user input
	output  chan string = make(chan string, queueMaxSize)
	printMx sync.Mutex
)

func server(ipAddress string, port string) {
	// Create the socket
	skt, err := net.Listen(protocol, ipAddress+":"+port)
	if err != nil {
		panic(err)
		return
	}
	fmt.Printf("Server listening on %s\n", skt.Addr().String())

	// TODO: Console input
	// go ConsoleInput()

	// Distribute clients in other goroutine
	go DistributeClients()

	// Start listening for clients trying to connect to the socket
	RunServer(skt)

	// Finish case
	wg.Wait()
	err = skt.Close()
	if err != nil {
		panic(err)
		return
	}
}

func RunServer(skt net.Listener) {
	// Running server
	for {
		// Listen
		select {
		case <-stop:
			return
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
}

/*
// ConsoleInput
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
*/

func DistributeClients() {
	// Here we will receive the connections from the channel to evaluate and distribute them.
	var isHandlerListening [numberOfHandlers]bool
	var handlerChan [numberOfHandlers]chan net.Conn
	var handlerCapacity [numberOfHandlers]uint32
	// Create handlers
	for i := uint32(0); i < numberOfHandlers; i++ {
		// Create a handler
		handlerChan[i] = make(chan net.Conn, maxClientsPerHandler)
		go HandleClients(i, handlerChan[i])
		// Set its status
		isHandlerListening[i] = true
		handlerCapacity[i] = maxClientsPerHandler
	}
	activeHandlers := numberOfHandlers

	// Status control
	for {
		// Saturated servers case
		if activeHandlers == 0 {
			continue
		}
		fmt.Printf("# of free handlers: %d\n", activeHandlers)
		// Give the connection to the freest server
		select {
		case conn := <-clients:
			designedHandler := FreestHandler(&isHandlerListening, &handlerCapacity)
			// Not valid case or no free handlers
			if (designedHandler >= numberOfHandlers) || (designedHandler < 0) {
				clients <- conn
				break
			}
		default:
		}
	}
}

func FreestHandler(active *[numberOfHandlers]bool, capacity *[numberOfHandlers]uint32) uint32 {
	var bestId, bestCapacity uint32 = maxClientsPerHandler, maxClientsPerHandler
	for i := uint32(0); i < numberOfHandlers; i++ {
		if (capacity[i] < bestCapacity) && (active[i]) {
			bestCapacity = capacity[i]
			bestId = i
		}
	}
	return bestId
}

func HandleClients(id uint32, clients chan net.Conn) {
	// The queue is automatically managed by the distributor
	for {
		select {
		case conn := <-clients:
			// Manage the connection

			// Finish
			conn.Close()
		}
	}
}
