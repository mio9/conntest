package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	log.Println("New connection from", conn.RemoteAddr())

	reader := bufio.NewReader(conn)
	lastMsg := "" // Store the last received message
	for {
		message, err := reader.ReadString('\n')

		if err != nil {
			log.Println("Error reading:", err)
			log.Printf("Last message is %s", lastMsg)
			break
		}
		lastMsg = message
		// fmt.Printf("Received: %s", message)

		_, err = conn.Write([]byte("Echo: " + message))
		if err != nil {
			log.Println("Error writing:", err)
			log.Printf("Last message is %s", lastMsg)
			break
		}
	}
}

func main() {
	path, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory:", err)
		os.Exit(1)
	}
	fmt.Printf("Current directory: %s\n", path)
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error starting server:", err)
		os.Exit(1)
	}
	defer listener.Close()
	fmt.Println("Server started on :8080")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		go handleConnection(conn)
	}
}
