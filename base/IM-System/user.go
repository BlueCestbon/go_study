package main

import "net"

type User struct {
	Name string
	Addr string
	C    chan string
	Conn net.Conn
}

// NewUser 创建用户的API
func NewUser(conn net.Conn) *User {
	addr := conn.RemoteAddr().String()

	user := &User{
		Name: addr,
		Addr: addr,
		C:    make(chan string),
		Conn: conn,
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
