/*
	The client should have a graphical interface.
	That is something to be done in the future (maybe).
*/

package client

import (
	"fmt"
	"net"
	"time"
)

func Client(ip_address, port string) {
	con, err := net.Dial("tcp", ip_address+":"+port)
	if err != nil {
		panic(err)
	}
	defer con.Close()
	for {
		fmt.Println("Client > Im connected with ", con.RemoteAddr())
		time.Sleep(10 * time.Second)
	}
}
