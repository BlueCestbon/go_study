package main

import (
	"net"
	"regexp"
	"strings"
)

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

// DoMsg 处理消息
func (user *User) DoMsg(msg string) {
	// 以rename|开头，Unicode字符结尾
	pattern := `^rename\|.*.$`
	reRename := regexp.MustCompile(pattern)
	if msg == "who" {
		for _, onlineUser := range user.Server.OnlineMap {
			msg = "[" + onlineUser.Addr + "]" + onlineUser.Name + ":" + "在线...\n"
			user.SendMsg(msg)
		}
	} else if reRename.MatchString(msg) {
		newName := strings.Split(msg, "|")[1]
		user.Server.mapLock.Lock()
		// 判断新用户名是否存在
		_, ok := user.Server.OnlineMap[newName]
		if ok {
			user.SendMsg("当前用户名已存在，请更换")
		} else {
			// 删除旧的
			delete(user.Server.OnlineMap, user.Name)
			user.Name = newName
			// 添加新的
			user.Server.OnlineMap[user.Name] = user
			user.Server.mapLock.Unlock()
			user.SendMsg("rename success to " + newName)
		}
	} else {
		go user.Server.BroadCast(user, msg)
	}
}

// SendMsg 发消息给当前客户端
func (user *User) SendMsg(msg string) {
	user.Conn.Write([]byte(msg))
}
