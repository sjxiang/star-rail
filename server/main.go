package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"golang.org/x/net/websocket"
	
	"star-rail/server/game"
)

func main() {
	fmt.Println("开始")

	go func() {
		game.GetManageMatch().Run()	
	}()

	http.Handle("/", websocket.Handler(WebsocketHandler))

	log.Fatal(http.ListenAndServe(":3000", nil))
}


func WebsocketHandler(ws *websocket.Conn) {
	var player *game.Player

	for {

		var msg []byte
		
		ws.SetReadDeadline(time.Now().Add(3 * time.Second))  // 3 s 没有消息交互，连接是否超时，确保畅通。
		err := websocket.Message.Receive(ws, &msg)
		
		if err != nil {
			netErr, ok := err.(net.Error)
			if ok && netErr.Timeout() {  // 3 s 超时，没关系
				continue
			}

			fmt.Println(err)
			break
		}

		fmt.Println(string(msg))

		// Player 为空，则登录；不为空，则在线，把消息广播出去
		if player == nil {
			player = game.GetManagePlayer().PlayerLogin(ws)
		}

		if player != nil {
			testMsg := fmt.Sprintf("id %d \n", player.UserId)
			game.GetManagePlayer().BoardCast([]byte(testMsg))
		}

	}
}