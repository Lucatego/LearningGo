package main

import server "TextChat/src/server"

func main() {
	const ipAddress, port = "127.0.0.1", "1080"
	server.Server(ipAddress, port)
}
