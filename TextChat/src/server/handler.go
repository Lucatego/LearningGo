package server

import (
	"fmt"
	"net"
	"sync/atomic"
	"time"
)

type Handler struct {
	id uint64

	Clients        chan net.Conn
	ClientsInQueue uint64
}

/*
 *	Public functions to work with the Server structure
 */

func (handler *Handler) Initialize(id, numberOfClients uint64) {
	handler.id = id
	handler.Clients = make(chan net.Conn, numberOfClients)
	handler.ClientsInQueue = 0
}

func (handler *Handler) HandleClients() {
	defer fmt.Println("I'm leaving! ID:", handler.id)
	fmt.Printf("Handler #%d > I'm running.\n", handler.id)
	// Receive the client assigned
	for {
		select {
		case client := <-handler.Clients:
			// Start
			fmt.Printf("Handler #%d > Received connection from %s\n", handler.id, client.RemoteAddr().String())
			time.Sleep(3 * time.Second)
			fmt.Printf("Handler #%d > Ending connection\n", handler.id)
			// End
			_ = client.Close()
			atomic.AddUint64(&handler.ClientsInQueue, ^uint64(0))
		}
	}
}

func (handler *Handler) ReceiveClient(client net.Conn) {
	handler.Clients <- client
	atomic.AddUint64(&handler.ClientsInQueue, 1)
}

func (handler *Handler) CloseHandler() {
	close(handler.Clients)
}
