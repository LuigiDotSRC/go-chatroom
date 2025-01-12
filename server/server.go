package server

import (
	"log"
	"net"
)

func StartServer() {
	listener, err := net.Listen("tcp4", "127.0.0.1:5000")
	if err != nil {
		log.Fatalf("Could not initialize tcp socket: %v", err)
	}

	listener.Accept()
}
