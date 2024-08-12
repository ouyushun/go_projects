package main

import (
	"fmt"
	"io"
	"net"
	"sync"
	"time"
)

type Server struct {
	Ip string
	Port int
	//在线用户列表
	OnlineMap map[string]*User
	MapLock sync.Mutex

	ChanMessage chan string
}

func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip             : ip,
		Port           : port,
		OnlineMap      : make(map[string]*User),
		ChanMessage:     make(chan string),
	}
	return server
}

//从message管道读取数据发送给所有用户
func (server *Server) ListenMessage()  {
	for  {

		msg := <- server.ChanMessage

		server.MapLock.Lock()
		for _, clientUser := range server.OnlineMap {
			clientUser.C <- msg
		}
		server.MapLock.Unlock()
	}
}

//监听广播Chan, 有消息立即发送给所有在线User
func (server *Server) Broadcast(user *User, msg string) {
	sendMsg := "[" + user.Addr + "] " + user.Name + ":" + msg
	server.ChanMessage <- sendMsg
}

func (server *Server) Handler(conn net.Conn)  {

	//创建user
	user := GetNewUser(conn, server)
	user.Online()
	//接收客户端发送的消息
	isLive := make(chan bool)
	go func() {
		buf := make([]byte, 4096)
		for {
			n, err := conn.Read(buf)
			if n == 0 {
				user.Offline()
				return
			}
			if err != nil && err != io.EOF {
				fmt.Println("read err: ", err)
				return
			}

			//提取用户的消息
			msg := string(buf[:n-1])
			user.DoMessage(msg)

			//用户的任意消息， 代表当前用户是活跃的
			isLive <- true
		}


	}()


	//todo 当前handler  阻塞  ?
	for {
		select {
		case <-isLive:

		case <- time.After(time.Second * 120):
			user.sendMsg("超时强制下线")
			close(user.C)
			conn.Close()
			//退出当前用户的Handler
			return
		}

	}

}

func (server *Server) Start() {
	listenner, err := net.Listen("tcp", fmt.Sprintf("%s:%d", server.Ip, server.Port))

	if err != nil {
		fmt.Println("连接失败, err: ", err)
		return
	}

	defer listenner.Close()

	//启动监听message的协程, 有msg推送给ChanMessage时候， 就会广播给在线用户
	go server.ListenMessage()

	for  {
		//阻塞
		conn, err := listenner.Accept()

		if err != nil {
			fmt.Println("listnner accept err")
			continue
		}

		go server.Handler(conn)
	}
}
