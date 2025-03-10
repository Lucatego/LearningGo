package server

import (
	"fmt"
	"net"
	"sync"
	"time"
)

type Handler struct {
	id        uint64
	handlerWG *sync.WaitGroup // Should only allow wg.Done()

	// TODO: Use more channels to use more types of messages
	Clients chan net.Conn
}

/*
 *	Public functions to work with the Server structure
 */

func (handler *Handler) Initialize(id, numberOfClients uint64, wg *sync.WaitGroup) {
	handler.id = id
	handler.handlerWG = wg

	handler.Clients = make(chan net.Conn, numberOfClients)
}

func (handler *Handler) HandleClients() {
	fmt.Printf("Handler #%d > I'm running.\n", handler.id)
	// Receive the client assigned
	for client := range handler.Clients {
		if client == nil {
			fmt.Printf("Handler #%d is closed.\n", handler.id)
			break
		}
		// Start
		fmt.Printf("Handler #%d > Processing connection from %s\n", handler.id, client.RemoteAddr().String())
		time.Sleep(3 * time.Second)
		// End
		_ = client.Close()
	}
	// End
	handler.closeHandler()
}

/*
 *	Private functions used inside the server
 */

func (handler *Handler) closeHandler() {
	close(handler.Clients)
	handler.handlerWG.Done()
}
