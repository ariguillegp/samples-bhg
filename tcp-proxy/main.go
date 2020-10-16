package main

import (
	"io"
	"log"
	"net"
)

func handler(pxy net.Conn) {
	// Forward traffic to final endpoint
	end := "example.com:8080"
	dst, err := net.Dial("tcp", end)
	if err != nil {
		log.Fatalln("Destination host unreachable")
	}
	// Close connection to final destination after all the processing
	defer dst.Close()

	go func() {
		// Copying data from client to final destination
		if _, err := io.Copy(dst, pxy); err != nil {
			log.Fatalln(err)
		}
	}()

	// Copying reply from final destination back to client
	if _, err := io.Copy(pxy, dst); err != nil {
		log.Fatalln(err)
	}
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
