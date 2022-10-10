package main

import (
	"io"
	"log"
	"os"

	"golang.org/x/net/websocket"
)


func main() {
	ws, err := websocket.Dial("ws://127.0.0.1:3000", "", "http://127.0.0.1:3000")
	if err != nil {
		panic(err)
	}

	done := make(chan struct{})

	go func() {
		io.Copy(os.Stdout, ws)
		log.Println("done")
		done <- struct{}{}  // 向 main goroutine 发信号
	}()

	// 阻塞
	wsCopy(ws, os.Stdin)

	ws.Close()
	<- done
}


func wsCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}


