package main

import (
	"bytes"
	"net"
	"testing"
	"time"
)

func TestServer(t *testing.T) {
	go StartServer()
	time.Sleep(1 * time.Second)

	client1, err := net.Dial("tcp4", "127.0.0.1:5000")
	if err != nil {
		t.Fatalf("Could not connect client1: %v", err)
	}
	defer client1.Close()

	client2, err := net.Dial("tcp4", "127.0.0.1:5000")
	if err != nil {
		t.Fatalf("Could not connect client2: %v", err)
	}
	defer client2.Close()

	// send a message from client1
	message := []byte("Hello from client1")
	_, err = client1.Write(message)
	if err != nil {
		t.Fatalf("client1 could not send message: %v", err)
	}

	// wait for message broadcast
	time.Sleep(500 * time.Microsecond)

	// read message from client2
	buffer := make([]byte, 1024)
	client2.SetReadDeadline(time.Now().Add(1 * time.Second))
	n, err := client2.Read(buffer)
	if err != nil {
		t.Fatalf("client2 could not read message %v", err)
	}

	received := bytes.TrimSpace(buffer[:n])
	expected := bytes.TrimSpace(message)

	recieved := buffer[:n]
	if string(recieved) != string(expected) {
		t.Errorf("Expeceted '%s' got '%s'", expected, received)
	}
}
