package server

import (
	"log"
	"net"
)

var (
	connections = make([]net.Conn, 0)
)

func StartServer() {
	listener, err := net.Listen("tcp4", "127.0.0.1:5000")
	if err != nil {
		log.Fatalf("Could not initialize tcp socket: %v", err)
	}
	defer listener.Close()

	log.Printf("Server running on: %s", listener.Addr().String())
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Could not accept connection: %v", err)
			continue
		}
		connections = append(connections, conn)
		go handler(conn)
	}
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
			broadcast(in)
		}
	}
}

func broadcast(msg []byte) {
	for _, c := range connections {
		c.Write(msg)
	}
}
