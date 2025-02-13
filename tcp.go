package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func tcpMain() {
	listener, err := net.Listen("tcp", ":8081")
	if err != nil {
		fmt.Println("Error starting TCP server:", err)
		os.Exit(1)
	}
	defer listener.Close()

	logtcp := log.New(os.Stdout, "[TCP] ", log.LstdFlags)
	logtcp.Println("Server tcp is listening on port 8081")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	logtcp := log.New(os.Stdout, "[TCP] ", log.LstdFlags)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading from connection:", err)
			return
		}
		logtcp.Print("Message received:", message)
		conn.Write([]byte("Message received\n"))
	}
}
