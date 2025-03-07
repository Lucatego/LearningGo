package main

import client "TextChat/src/client"

func main() {
	const ipAddress, port = "127.0.0.1", "1080"
	client.Client(ipAddress, port)
}
