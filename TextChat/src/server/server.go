package server

import (
	"errors"
	"fmt"
	"net"
	"sync"
	"time"
)

const (
	connectionTimeout = 5 * time.Second
)

type Server struct {
	numberOfHandlers  uint64
	clientsPerHandler uint64
	url               string
	protocol          string

	isRunning bool
	skt       net.Listener
	wg        sync.WaitGroup
	// TODO: Implement more channels for different kind of messages sent to handlers or distributor
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
		server.handlers[i].Initialize(uint64(i), server.clientsPerHandler, &server.wg)
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
			if input == "exit" {
				server.isRunning = false
				break
			}
			fmt.Printf("Server > %s\n", input)
		default:
			// Accept a client and validate
			conn, err := server.skt.Accept()
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Printf("Server > New connection from %s\n", conn.RemoteAddr().String())

			// TODO: Evaluate, retrying or closing connection if server is full.
			// Send clients to the distributor goroutine avoiding deadlocks
			select {
			case server.clients <- conn:
				break
			case <-time.After(connectionTimeout):
				fmt.Printf("Server > Rejected connection from %s.\n", conn.RemoteAddr().String())
				_ = conn.Close()
			}
		}
	}
	// End safely
	server.closeServer()
}

/*
 *	Private functions used inside the server
 */

func (server *Server) distributeClients() {
	// TODO: Evaluate the usage of for { select { case end: case client: } } to manage more messages.
	for client := range server.clients {
		// End server case
		if client == nil {
			fmt.Printf("Server > Stopping distributor.\n")
			// Send signal for handlers to end
			for i := range server.handlers {
				server.handlers[i].Clients <- nil
			}
			break
		}
		fmt.Printf("Server > Distributing client from %s\n", client.RemoteAddr().String())
		// Select a handler
		id, err := server.selectHandler()
		if err != nil {
			fmt.Println(err)
			// TODO: Check this send, this is the cause of a deadlock, but right now, because of the
			//       > operator in selectHandler(), the deadlock won´t happen.
			server.clients <- client
			continue
		}
		// Send the client
		server.handlers[id].Clients <- client
	}
}

func (server *Server) selectHandler() (uint64, error) {
	// Search
	lowestUsage, bestId := ^uint64(0), server.numberOfHandlers+1
	for i := range server.handlers {
		usage := uint64(len(server.handlers[i].Clients))
		if usage == 0 {
			return uint64(i), nil
		}
		if usage < lowestUsage {
			lowestUsage = usage
			bestId = uint64(i)
		}
	}
	// Fail case - Note: the > avoids deadlocks, while the >= can make them happen.
	// Here we must evaluate, aborting connection or retrying to send it to a channel.
	// Note: I think it's better to send the connection to the first handler found (id: 0)
	//		 because in this case, we avoid the deadlock and keep managing the connections.
	if bestId >= server.numberOfHandlers || lowestUsage > server.clientsPerHandler {
		return bestId, errors.New("error: out of range handler (not found id)")
	}
	return bestId, nil
}

func (server *Server) closeServer() {
	// Close the socket
	_ = server.skt.Close()
	// Send end signal to distributor
	server.clients <- nil
	// Wait for handlers to finish
	server.wg.Wait()
	// Free memory
	close(server.clients)
	server.handlers = nil
}
