package main

import (
	"net"
	"strings"
)

type User struct {
	Name string
	Addr string
	C chan string //接收来自server的广播消息，
	Conn net.Conn

	server *Server
}

func GetNewUser(conn net.Conn, server *Server) *User {
	userAddr := conn.RemoteAddr().String() //地址
	user := &User {
		Name: userAddr,
		Addr: userAddr,
		C: make(chan string),
		Conn: conn,
		server: server,
	}

	//启动 监听当前用户 channel的协程
	go user.ListennerMsg()

	return user
}

func (user *User) Online() {

	//将用户添加到map中
	user.server.MapLock.Lock()
	user.server.OnlineMap[user.Name] = user
	user.server.MapLock.Unlock()

	//广播
	user.server.Broadcast(user, "已上线")
}


func (user *User) Offline() {

	//用户下线 ， 从onlineMap中删除
	user.server.MapLock.Lock()
	delete(user.server.OnlineMap, user.Name)
	user.server.MapLock.Unlock()

	//广播
	user.server.Broadcast(user, "已下线")
}

//给当前user对应的用户发送消息
func (user *User) sendMsg(msg string) {
	user.Conn.Write([]byte(msg))
}


func (user *User) DoMessage(msg string) {

	if msg == "who" {
		//查询当前在线用户
		user.server.MapLock.Lock()
		for _, U:= range user.server.OnlineMap {
			sendMsg := "[" + U.Name + "]在线\n"
			user.sendMsg(sendMsg)
		}
		user.server.MapLock.Unlock()
		return
	}

	//修改用户名  协议  rename|张三
	if len(msg) > 7 && msg[:7] == "rename|" {

		oldName := user.Name
		newName := strings.Split(msg, "|")[1]

		if _, ok := user.server.OnlineMap[newName]; ok {
			user.sendMsg("名称已存在:" + newName)
		} else {
			user.server.MapLock.Lock()
			delete(user.server.OnlineMap, oldName)
			user.server.OnlineMap[newName] = user
			user.server.MapLock.Unlock()

			user.Name = newName
			user.sendMsg("用户名已更新, 新用户名:" + user.Name + "\n")
		}

		return
	}

	//私聊 to|张三|msg
	if len(msg) > 4 && msg[:3] == "to|" {
		//获取对方用户名
		remoteName := strings.Split(msg, "|")[1]
		//解析msg内容
		msg := strings.Split(msg, "|")[2]

		if msg == "" {
			user.sendMsg("消息内容为空")
			return
		}

		//获取remoteUser对象
		remoteUser , ok := user.server.OnlineMap[remoteName]
		if !ok {
			user.sendMsg("用户不存在")
			return
		}

		//发送消息
		remoteUser.sendMsg(user.Name + "对您说:" + msg)

		return
	}

	user.server.Broadcast(user, msg)
}


//监听当前user 管道， 有消息发送给客户端
func (user *User) ListennerMsg()  {
	for {
		msg := <- user.C
		user.Conn.Write([]byte(msg + "\n"))
	}
}