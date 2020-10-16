package main

import (
	"fmt"
	"net"
	"sort"
)

func worker(ports, results chan int) {
	for p := range ports {
		address := fmt.Sprintf("scanme.nmap.org:%d", p)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			results <- 0
			continue
		}
		conn.Close()
		results <- p
	}
}

func main() {
	ports := make(chan int, 10)
	results := make(chan int)
	var openPorts []int

	for i := 0; i < cap(ports); i++ {
		go worker(ports, results)
	}

	go func() {
		for i := 21; i <= 54; i++ {
			ports <- i
		}
	}()

	for i := 21; i <= 54; i++ {
		res := <-results
		if res != 0 {
			openPorts = append(openPorts, res)
		}
	}

	close(ports)
	close(results)

	sort.Ints(openPorts)
	for _, p := range openPorts {
		fmt.Println(p)
	}
}
