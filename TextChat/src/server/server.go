package server

import (
	"errors"
	"fmt"
	"net"
	"sync"
	"sync/atomic"
	"syscall"
	"time"
)

type Server struct {
	numberOfHandlers  uint64
	clientsPerHandler uint64
	url               string
	protocol          string

	isRunning    bool
	skt          net.Listener
	wg           sync.WaitGroup
	clients      chan net.Conn
	queueMaxSize uint64
	handlers     []Handler

	terminalInput chan string
}

/*
 *	Public functions to work with the Server structure
 */

func (server *Server) SetServer(numberOfHandlers, clientsPerHandler uint64) {
	server.isRunning = false
	// Set the parameters
	server.numberOfHandlers = numberOfHandlers
	server.clientsPerHandler = clientsPerHandler
	server.queueMaxSize = numberOfHandlers * clientsPerHandler
	// Allocate memory
	server.clients = make(chan net.Conn, server.queueMaxSize)
	server.handlers = make([]Handler, numberOfHandlers)
	fmt.Printf("Server > Server initialized with %d handlers.\n", numberOfHandlers)
}

func (server *Server) CreateSocket(ipAddress, port, protocol string) {
	var err error

	server.url = ipAddress + ":" + port
	server.protocol = protocol

	server.skt, err = net.Listen(protocol, server.url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Server > Socket created on %s.\n", server.skt.Addr().String())
}

func (server *Server) RunServer() {
	// Initialize handlers
	for i := range server.handlers {
		server.handlers[i].Initialize(uint64(i), server.clientsPerHandler)
		go server.handlers[i].HandleClients()
		server.wg.Add(1)
	}
	server.isRunning = true
	// In other goroutine that receives the clients
	go server.distributeClients()
	// The server main work
	for server.isRunning {
		select {
		case input := <-server.terminalInput:
			fmt.Printf("Server > %s\n", input)
		default:
			// Accept a client and validate
			conn, err := server.skt.Accept()
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Printf("Server > New connection from %s\n", conn.RemoteAddr().String())

			// Send clients to the distributor goroutine
			server.clients <- conn
		}
	}

	server.closeServer()
}

/*
 *	Private functions used inside the server
 */

func (server *Server) distributeClients() {
	for {
		select {
		case client := <-server.clients:
			// Select a handler
			id, err := server.selectHandler()
			if err != nil {
				fmt.Println(err)
				server.clients <- client
				continue
			}
			// Send the client
			server.handlers[id].ReceiveClient(client)
		default:
			// Wait until a client connects
			fmt.Printf("Distributor > Waiting for a client to connect...\n")
			time.Sleep(2 * time.Second)
		}
	}
}

func (server *Server) selectHandler() (uint64, error) {
	// Search
	lowestUsage, bestId := uint64(syscall.INFINITE), server.numberOfHandlers+1
	for i := range server.handlers {
		usage := atomic.LoadUint64(&server.handlers[i].ClientsInQueue)
		if usage < lowestUsage {
			lowestUsage = usage
			bestId = uint64(i)
		}
	}
	// Return found
	if bestId >= server.numberOfHandlers || lowestUsage > server.clientsPerHandler {
		return bestId, errors.New("error: out of range handler (not found id)")
	}
	return bestId, nil
}

func (server *Server) closeServer() {
	// Wait for handlers to finish
	server.wg.Wait()
	// Terminate them
	for i := range server.handlers {
		server.handlers[i].CloseHandler()
	}
	close(server.clients)
	// Delete socket and free memory
	_ = server.skt.Close()
	server.handlers = nil
}
