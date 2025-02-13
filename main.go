package main

import (
	"time"
)

func main() {
	go httpMain()
	go tcpMain()
	go udpMain()

	for {
		time.Sleep(1 * time.Second)
	}
}
