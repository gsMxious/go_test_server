package main

import (
	"fmt"
	"net"
	"os"
)

func udpMain() {
	addr := net.UDPAddr{
		Port: 8080,
		IP:   net.ParseIP("127.0.0.1"),
	}

	conn, err := net.ListenUDP("udp", &addr)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	defer conn.Close()

	buffer := make([]byte, 1024)
	for {
		n, clientAddr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println("Error: ", err)
			continue
		}

		fmt.Printf("Received %s from %s\n", string(buffer[:n]), clientAddr)

		_, err = conn.WriteToUDP([]byte("Message received"), clientAddr)
		if err != nil {
			fmt.Println("Error: ", err)
			continue
		}
	}
}
