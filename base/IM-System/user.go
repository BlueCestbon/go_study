package main

import "net"

type User struct {
	Name   string
	Addr   string
	C      chan string
	Conn   net.Conn
	Server *Server
}

// NewUser 创建用户的API
func NewUser(conn net.Conn, server *Server) *User {
	addr := conn.RemoteAddr().String()

	user := &User{
		Name:   addr,
		Addr:   addr,
		C:      make(chan string),
		Conn:   conn,
		Server: server,
	}
	// 初始化之后，就监听自己的这个channel，有消息就输出到客户端
	go user.ListenMessage()

	return user
}

func (user *User) ListenMessage() {
	for {
		msg := <-user.C
		user.Conn.Write([]byte(msg + "\n"))
	}
}

// Online 上线
func (user *User) Online() {
	user.Server.mapLock.Lock()
	user.Server.OnlineMap[user.Name] = user
	user.Server.mapLock.Unlock()

	// 广播当前用户上线消息
	go user.Server.BroadCast(user, "已上线")
}

// Offline 下线
func (user *User) Offline() {
	user.Server.mapLock.Lock()
	delete(user.Server.OnlineMap, user.Name)
	user.Server.mapLock.Unlock()
	// 广播当前用户下线消息
	go user.Server.BroadCast(user, "下线")
}

// SendMsg 处理消息
func (user *User) SendMsg(msg string) {
	go user.Server.BroadCast(user, msg)
}
