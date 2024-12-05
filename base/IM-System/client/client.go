package main

import (
	"flag"
	"fmt"
	"net"
)

type Client struct {
	ServerIp   string
	ServerPort int
	Name       string
	Conn       net.Conn
	clientFlag int // 当前客户端的模式
}

func NewClient(serverIp string, serverPort int) *Client {
	// 创建客户端对象
	client := &Client{
		ServerIp:   serverIp,
		ServerPort: serverPort,
		clientFlag: 999,
	}
	// 连接到server
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", serverIp, serverPort))
	if err != nil {
		fmt.Println("net.Dial err, ", err)
		return nil
	}

	client.Conn = conn

	// 返回对象
	return client
}

var serverIp string
var serverPort int

func init() {
	flag.StringVar(&serverIp, "ip", "127.0.0.1", "设置服务器IP地址(默认127.0.0.1)")
	flag.IntVar(&serverPort, "port", 8888, "设置服务器端口(默认8888)")
}

func (client *Client) menu() bool {
	var clientFlag int
	fmt.Println("1.公聊模式")
	fmt.Println("2.私聊模式")
	fmt.Println("3.更改用户名")
	fmt.Println("0.退出")

	fmt.Scanln(&clientFlag)

	if clientFlag >= 0 && clientFlag <= 3 {
		client.clientFlag = clientFlag
		return true
	} else {
		fmt.Println("请输入合法范围内的数字")
		return false
	}
}

func (client *Client) run() {
	for client.clientFlag != 0 {
		for client.menu() != true {
		}
		switch client.clientFlag {
		case 1:
			// 公聊模式
			fmt.Println("选择了公聊模式")
			break
		case 2:
			// 私聊模式
			fmt.Println("选择了私聊模式")
			break
		case 3:
			// 更改用户名
			fmt.Println("选择了更改用户名")
			break
		}
	}
}

func main() {
	flag.Parse()
	client := NewClient(serverIp, serverPort)
	if client == nil {
		fmt.Println(">>>>>连接失败...")
		return
	}

	fmt.Println(">>>>>连接成功...")

	client.run()
}
