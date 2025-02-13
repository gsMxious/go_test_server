package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func udpMain() {
	addr := net.UDPAddr{
		Port: 9000,
		IP:   net.IPv4zero,
	}

	conn, err := net.ListenUDP("udp", &addr)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	defer conn.Close()
	logudp := log.New(os.Stdout, "[UDP] ", log.LstdFlags)
	logudp.Println("Server is listening on port 9000")

	buffer := make([]byte, 1024)
	for {
		n, clientAddr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println("Error: ", err)
			continue
		}

		logudp.Printf("Received %s from %s\n", string(buffer[:n]), clientAddr)

		_, err = conn.WriteToUDP([]byte("Message received"), clientAddr)
		if err != nil {
			fmt.Println("Error: ", err)
			continue
		}
	}
}
