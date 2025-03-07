/*
	The client should have a graphical interface.
	That is something to be done in the future (maybe).
*/

package client

import (
	"fmt"
	"net"
	"sync"
	"time"
)

func Client(ipAddress, port string, wg *sync.WaitGroup, id int) {
	if wg != nil {
		defer (*wg).Done()
	}

	con, err := net.Dial("tcp", ipAddress+":"+port)
	if err != nil {
		panic(err)
	}
	defer con.Close()

	fmt.Printf("Client #%d> Im connected with %s\n", id, con.RemoteAddr())
	time.Sleep(15 * time.Second)

}
