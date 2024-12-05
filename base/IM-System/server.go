package main

import (
	"fmt"
	"io"
	"net"
	"sync"
)

type Server struct {
	Ip   string
	Port int
	// 在线的用户
	OnlineMap map[string]*User
	mapLock   sync.RWMutex
	// server监控channel，有信息就触发广播
	ServerChannel chan string
}

func NewServer(ip string, port int) *Server {
	return &Server{
		Ip:            ip,
		Port:          port,
		OnlineMap:     make(map[string]*User),
		ServerChannel: make(chan string),
	}
}

// BroadCast 广播消息
func (server *Server) BroadCast(user *User, msg string) {
	sendMsg := "[" + user.Addr + "]" + user.Name + ":" + msg
	server.ServerChannel <- sendMsg
}

// Handler 处理连接上线之后的操作
func (server *Server) Handler(conn net.Conn) {
	// fmt.Println("成功建立连接")
	// 用户上线，添加到OnlineMap
	user := NewUser(conn, server)
	user.Online()

	// 接收客户端消息
	go func() {
		buf := make([]byte, 4096)
		for {
			n, err := conn.Read(buf)
			// 如果读到的数据是0，就说明是正常关闭的
			if n == 0 {
				user.Offline()
				return
			}
			// 以EOF结尾的
			if err != nil && err != io.EOF {
				fmt.Println("Conn Read err:", err)
				return
			}

			// 去除用户信息的’\n‘
			msg := string(buf[:n-1])

			// 用户处理这个消息
			user.SendMsg(msg)
		}
	}()

	// 阻塞当前的handler
	select {}
}

// ListenServerChannel 监听server的广播信道，一旦有消息，就发送给全部在线的user
func (server *Server) ListenServerChannel() {
	for {
		sendMsg := <-server.ServerChannel
		server.mapLock.Lock()
		for _, user := range server.OnlineMap {
			user.C <- sendMsg
		}
		server.mapLock.Unlock()
	}
}

// Start 启动服务器的接口
func (server *Server) Start() {
	// socket listen
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", server.Ip, server.Port))
	if err != nil {
		fmt.Println("net.Listen err: ", err)
		return
	}

	// 监听sever的广播信道
	go server.ListenServerChannel()

	// close
	defer listener.Close()

	for {
		// accept
		fmt.Printf("start listen %s:%d\n", server.Ip, server.Port)
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener conn err: ", err)
			continue
		}
		// handle
		go server.Handler(conn)
	}
}
