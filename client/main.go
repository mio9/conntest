package main

import (
	"bufio"
	"fmt"
	"net"
	"time"
)

func connectToServer() net.Conn {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return nil
	}
	fmt.Println("Connected to server")
	return conn
}

func main() {
	var conn net.Conn

	for {
		if conn == nil {
			conn = connectToServer()
			if conn == nil {
				time.Sleep(5 * time.Second) // wait before retrying connection
				continue
			}
		}

		message := "Hello Server\n"
		_, err := conn.Write([]byte(message))
		if err != nil {
			fmt.Println("Error writing to server:", err)
			conn.Close()
			conn = nil
			time.Sleep(2 * time.Second) // wait before retrying connection
			continue
		}

		reader := bufio.NewReader(conn)
		response, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading from server:", err)
			conn.Close()
			conn = nil
			time.Sleep(2 * time.Second) // wait before retrying connection
			continue
		}

		fmt.Printf("Server response: %s", response)

		time.Sleep(1 * time.Second) // send a message every second
	}
}
