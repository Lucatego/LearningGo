package main

import "TextChat/src/server"

const (
	ipAddress, port, protocol = "127.0.0.1", "1080", "tcp"

	numberOfHandlers  uint64 = 4
	clientsPerHandler uint64 = 4
)

func main() {
	var s server.Server
	// Initialize
	s.SetServer(numberOfHandlers, clientsPerHandler)
	s.CreateSocket(ipAddress, port, protocol)
	// Run
	s.RunServer()
}
