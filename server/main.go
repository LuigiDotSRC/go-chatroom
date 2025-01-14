package main

import (
	"log"
	"net"
	"sync"
)

var (
	connections = make([]net.Conn, 0)
	mu          sync.Mutex
)

func main() {
	StartServer()
}

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

		mu.Lock()
		connections = append(connections, conn)
		mu.Unlock()

		go handler(conn)
	}
}

func handler(conn net.Conn) {
	defer conn.Close()

	for {
		in := make([]byte, 1024)
		n, err := conn.Read(in)
		if err != nil {
			log.Printf("Could not read input: %v", err)
			return
		}
		if n > 0 {
			log.Print(string(in))
			broadcast(in[:n])
		}
	}
}

func broadcast(msg []byte) {
	mu.Lock()
	defer mu.Unlock()
	for _, c := range connections {
		_, err := c.Write(msg)
		if err != nil {
			log.Printf("could not broadcast message: %v", err)
		}
	}
}
