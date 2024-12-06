package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
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

func (client *Client) PublicChat() {
	fmt.Println("选择了公聊模式")
	fmt.Println("请输入消息内容，输入exit退出")
	var msg string
	fmt.Scanln(&msg)

	for msg != "exit" {
		if len(msg) == 0 {
			fmt.Println("不能发送空消息")
			fmt.Scanln(&msg)
			continue
		}
		_, err := client.Conn.Write([]byte(msg + "\n"))
		if err != nil {
			fmt.Println("conn write err, ", err)
			break
		}
		// 置空，方便下次输入
		msg = ""
		fmt.Scanln(&msg)
	}
}

func (client *Client) ListOnlineUser() {
	_, err := client.Conn.Write([]byte("who\n"))
	if err != nil {
		fmt.Println("conn write err, ", err)
	}
}

func (client *Client) PrivateChat() {
	fmt.Println("选择了私聊模式")
	client.ListOnlineUser()
	fmt.Println("请输入要聊天的对象[用户名]，exit退出")
	var toUserName string
	fmt.Scanln(&toUserName)

	for toUserName != "exit" {
		fmt.Println("请输入消息内容，exit退出")
		var msg string
		fmt.Scanln(&msg)
		for msg != "exit" {
			if len(msg) == 0 {
				fmt.Println("不能发送空消息")
				fmt.Scanln(&msg)
				continue
			}
			sendMsg := "to|" + toUserName + "|" + msg
			_, err := client.Conn.Write([]byte(sendMsg + "\n"))
			if err != nil {
				fmt.Println("conn write err, ", err)
				break
			}
			// 置空，方便下次输入
			msg = ""
			fmt.Scanln(&msg)
		}
		client.ListOnlineUser()
		fmt.Println("请输入要聊天的对象[用户名]，exit退出")
		fmt.Scanln(&toUserName)
	}
}

func (client *Client) Rename() {
	fmt.Println("请输入用户名")
	fmt.Scanln(&client.Name)

	// 模拟手动输入这个协议
	sendMsg := "rename|" + client.Name
	_, err := client.Conn.Write([]byte(sendMsg + "\n"))
	if err != nil {
		fmt.Println("conn write err, ", err)
	}
}

func (client *Client) run() {
	for client.clientFlag != 0 {
		for client.menu() != true {
		}
		switch client.clientFlag {
		case 1:
			// 公聊模式
			client.PublicChat()
			break
		case 2:
			// 私聊模式
			client.PrivateChat()
			break
		case 3:
			// 更改用户名
			client.Rename()
			break
		}
	}
}

// DealResponse 显示返回结果，显示到标准输出
func (client *Client) DealResponse() {
	// 把client的消息拷贝到stdout输出上，永久阻塞监听
	io.Copy(os.Stdout, client.Conn)
}

func main() {
	flag.Parse()
	client := NewClient(serverIp, serverPort)
	if client == nil {
		fmt.Println(">>>>>连接失败...")
		return
	}

	fmt.Println(">>>>>连接成功...")

	// 开启一个goroutine去处理server的回执消息
	go client.DealResponse()

	// 客户端业务
	client.run()
}
