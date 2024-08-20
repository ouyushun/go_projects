package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
)

var UP = websocket.Upgrader{
	ReadBufferSize : 1024,
	WriteBufferSize : 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func Handler(w http.ResponseWriter,  r *http.Request) {
	if !websocket.IsWebSocketUpgrade(r) {
		w.Write([]byte("不是websocket"))
		return
	}

	conn, err := UP.Upgrade(w, r, nil)

	defer conn.Close()
	if err != nil {
		fmt.Println("连接失败: ", err)
		return
	}
	for {
		_, p, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("read error: ", err)
			break
		}
		conn.WriteMessage(websocket.TextMessage,  []byte("你说的是" + string(p) + "吗？"))
	}
	fmt.Println("服务端关闭")
}

func main() {
	http.HandleFunc("/ws", Handler)
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		return
	}
}
