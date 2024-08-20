package main

import (
	"bufio"
	"fmt"
	"github.com/gorilla/websocket"
	"os"
)



func main() {

	dialer := websocket.Dialer{}
	conn, _, err := dialer.Dial("ws://127.0.0.1:8888/ws", nil)
	if err != nil {
		return
	}

	//向服务端写操作
	go func(conn *websocket.Conn) {
		for {
			reader := bufio.NewReader(os.Stdin)
			line, _, _ := reader.ReadLine()
			conn.WriteMessage(websocket.TextMessage,  line)
		}
	}(conn)

	//从服务端读
	for {
		_, p, err := conn.ReadMessage()
		if err != nil {
			continue
		}
		fmt.Println(string(p))
	}
}