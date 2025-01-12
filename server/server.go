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

	conn, _ := listener.Accept()
	log.Printf("Server running on: %s", listener.Addr().String())
	handler(conn)
}

func handler(conn net.Conn) {
	defer conn.Close()
	var out []byte = []byte("PONG")
	for {
		in := make([]byte, 1024)
		n, err := conn.Read(in)
		if err != nil {
			log.Printf("Could not read input: %v", err)
			return
		}
		if n > 0 {
			conn.Write(out)
		}
	}
}
