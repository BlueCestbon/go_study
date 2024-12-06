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
	//for {
	//	msg := <-user.C
	//	// 当超时下线的时候，这里user.C已经关闭了，取到的msg是零值。
	//	// 并且下线的时候user.Conn已经close了，下面的Conn.Write就会报错
	//	// 所以得让ListenMessage知道这个conn关闭了，或者是channel已经关闭了，就不要执行下面的了
	//	_, err := user.Conn.Write([]byte(msg + "\n"))
	//	if err != nil {
	//		// 会一直输出err
	//		fmt.Println("conn write err, ", err)
	//		//return
	//	}
	//}

	for msg := range user.C {
		_, err := user.Conn.Write([]byte(msg + "\n"))
		if err != nil {
			//fmt.Println("conn write err, ", err)
			panic(err)
		}
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
	patternRename := `^rename\|.*.$`
	reRename := regexp.MustCompile(patternRename)

	// 私聊消息的正则
	patternPrivateChat := `^to\|.*.\|.*.$`
	rePrivateChat := regexp.MustCompile(patternPrivateChat)

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
	} else if rePrivateChat.MatchString(msg) {
		userMsg := strings.Split(msg, "|")
		toUserName, theMsg := userMsg[1], userMsg[2]
		toUser, ok := user.Server.OnlineMap[toUserName]
		if !ok {
			user.SendMsg("当前用户不在线，请重新选择用户")
			return
		}
		if len(theMsg) == 0 {
			user.SendMsg("信息为空，请重新发送")
			return
		}
		// 发给指定用户
		toUser.SendMsg("[" + user.Addr + "]" + user.Name + ":" + theMsg)
	} else {
		go user.Server.BroadCast(user, msg)
	}
}

// SendMsg 发消息给当前客户端
func (user *User) SendMsg(msg string) {
	user.Conn.Write([]byte(msg + "\n"))
}
