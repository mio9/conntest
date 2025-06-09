package main

import (
	"bufio"
	"fmt"
	"net"
	"time"
)

var conn net.Conn
var msg_id uint16 = 0

func connectToServer() net.Conn {
	conn, err := net.Dial("tcp", "172.28.10.1:8080")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return nil
	}
	fmt.Println("Connected to server")
	return conn
}

func killConnAndWait() {
	conn.Close()
	conn = nil
	time.Sleep(2 * time.Second) // wait before retrying connection
}

func main() {
	for {
		msg_id += 1
		if conn == nil {
			conn = connectToServer()
			if conn == nil {
				time.Sleep(2 * time.Second) // wait before retrying connection
				continue
			}
		}

		message := fmt.Sprintf("id:%d/ping\n", msg_id) //") "Hello Server\n"
		_, err := conn.Write([]byte(message))
		if err != nil {
			fmt.Println("Error writing to server:", err)
			killConnAndWait()
			continue
		}

		reader := bufio.NewReader(conn)
		response, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading from server:", err)
			killConnAndWait()
			continue
		}

		fmt.Printf("Server response: %s", response)

		time.Sleep(1 * time.Second) // send a message every second
	}
}
