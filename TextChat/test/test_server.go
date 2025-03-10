package main

import (
	"TextChat/src/client"
	"sync"
)

const numberOfClients = 64

func main() {
	const ipAddress, port = "127.0.0.1", "1080"
	// To emulate a lot of clients
	var wg sync.WaitGroup
	wg.Add(numberOfClients)
	for i := 0; i < numberOfClients; i++ {
		go client.Client(ipAddress, port, &wg, i)
	}
	wg.Wait()
}
