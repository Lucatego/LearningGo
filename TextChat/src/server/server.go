package server

import (
	"fmt"
	"net"
	"sync"
	"time"
)

const (
	protocol             string = "tcp"
	numberOfHandlers     uint32 = 4
	maxClientsPerHandler uint32 = 4
	queueMaxSize                = numberOfHandlers * maxClientsPerHandler
)

var (
	// To safely stop the server
	wg   sync.WaitGroup
	stop = make(chan struct{}, 1)

	// To manage connections
	clients = make(chan net.Conn, queueMaxSize)
	/*
		// To print on server console without interfering with the user input
		output  = make(chan string, queueMaxSize)
		printMx sync.Mutex
	*/
)

func Server(ipAddress, port string) {
	// Create the socket
	skt, err := net.Listen(protocol, ipAddress+":"+port)
	if err != nil {
		panic(err)
		return
	}
	fmt.Printf("Server > Server listening on %s\n", skt.Addr().String())

	// TODO: Console input
	// go ConsoleInput()

	// Distribute clients in other goroutine
	go distributeClients()

	// Start listening for clients trying to connect to the socket
	runServer(skt)

	// Finish case
	wg.Wait()
	err = skt.Close()
	if err != nil {
		panic(err)
		return
	}
}

func runServer(skt net.Listener) {
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
			fmt.Printf("Server > New connection from %s\n", conn.RemoteAddr().String())

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

func distributeClients() {
	// Here we will receive the connections from the channel to evaluate and distribute them.
	var handlerMutex = make([]sync.Mutex, numberOfHandlers)
	var handlerChan = make([]chan net.Conn, numberOfHandlers)
	// Create handlers
	for i := uint32(0); i < numberOfHandlers; i++ {
		// Create a handler
		handlerChan[i] = make(chan net.Conn, maxClientsPerHandler)
		go handleClients(i, handlerChan[i], &handlerMutex[i])
	}

	// Status control
	for {
		// Give the connection to the freest server
		select {
		case conn := <-clients:
			designedHandler := freestHandler(handlerChan, handlerMutex)
			// Not valid case or no free handlers
			if (designedHandler >= numberOfHandlers) || (designedHandler < 0) {
				clients <- conn
				break
			}
			// Valid case
			handlerChan[designedHandler] <- conn
		default:
			fmt.Printf("Distributor > Waiting for a client to connect...\n")
			// Wait until a client connects
			time.Sleep(2 * time.Second)
		}
	}
}

func freestHandler(channels []chan net.Conn, mutexes []sync.Mutex) uint32 {
	// Capacity means free space because we have the information about the number of items in the channel
	var bestId, lowestUsage = numberOfHandlers, maxClientsPerHandler
	// This should be done FAST...
	for i := range channels {
		// We block the handler to stop receiving connections to evaluate its capacity
		mutexes[i].Lock()
		temp := uint32(len(channels[i]))
		if temp < lowestUsage {
			lowestUsage = temp
			bestId = uint32(i)
		}
	}
	// After all handlers finished, we unlock them all
	for i := range mutexes {
		mutexes[i].Unlock()
	}
	return bestId
}

func handleClients(id uint32, clients chan net.Conn, mutex *sync.Mutex) {
	// The queue is automatically managed by the distributor
	for {
		mutex.Lock()
		select {
		// The len(clients) is the critic zone here.
		case conn := <-clients:
			// Receive the connection
			mutex.Unlock()
			// Manage the connection
			fmt.Printf("Handler #%d > Handling connection from %s\n", id, conn.RemoteAddr().String())
			time.Sleep(1 * time.Second)
			fmt.Printf("Handler #%d > Ending connection with %s\n", id, conn.RemoteAddr().String())
			// Finish
			err := conn.Close()
			if err != nil {
				fmt.Println(err)
			}
		default:
			mutex.Unlock()
		}
	}
}
