package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

// sseHandler 是用来处理 SSE 请求的处理器函数。
func sseHandler(w http.ResponseWriter, r *http.Request) {
	// 设置响应头以支持 Server-Sent Events
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	// 创建一个通道，用于控制循环
	done := make(chan bool)

	// 启动一个 goroutine 来监听客户端关闭连接的情况
	go func() {
		// 当客户端关闭连接时，读取请求体将返回错误
		n := 0
		for {
			time.Sleep(3 * time.Second)
			n++
			if n == 60 {
				done <- true
				return
			}
		}
	}()

	// 循环发送 "Hello, World!" 消息
	for {
		select {
		case <-done:
			log.Println("Sent: Hello, World! done")
			return
		default:
			// 发送消息
			fmt.Fprintf(w, "data: Hello, World!\n\n")
			// 刷新缓冲区，确保数据立即发送给客户端
			w.(http.Flusher).Flush()
			log.Println("Sent: Hello, World!")
			// 等待一段时间再发送下一个消息
			time.Sleep(3 * time.Second)
		}
	}
}

func httpMain() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/sse", sseHandler)
	log.Println("Starting server at port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err)
	}
}
