package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
)

type Client struct {
	ServerIP   string
	ServerPort int
	ClientName string
	conn       net.Conn
	flag       int // 当前客户端的模型
}

func NewClient(serverIP string, serverPotr int) *Client {
	//创建客户端对象
	client := &Client{
		ServerIP:   serverIP,
		ServerPort: serverPort,
		flag:       999,
	}

	//链接server
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", serverIP, serverPotr))
	if err != nil {
		fmt.Println("net.Dial error:", err)
		return nil
	}
	client.conn = conn

	//返回对象
	return client
}

func (clt *Client) menu() bool {
	var flag int
	fmt.Println("*****************************")
	fmt.Println("**********1.公聊模式**********")
	fmt.Println("**********2.私聊模式**********")
	fmt.Println("**********3.更新用户名********")
	fmt.Println("**********0.退出**************")
	fmt.Println("******************************")

	fmt.Scanln(&flag) // 输入flag

	if flag >= 0 && flag <= 3 {
		clt.flag = flag
		return true
	} else {
		fmt.Println(">>>>>>请输入合法的flag!<<<<<<")
		return false
	}
}

func (clt *Client) Run() {
	for clt.flag != 0 {
		//模式不对就一直循环菜单
		for clt.menu() == false {
		}
		//根据不同的模式处理不同的业务
		switch clt.flag {
		case 1:
			//公聊模式
			fmt.Println(">>>>>>公聊模式,输入exit退出")
			clt.PublicChat()
		case 2:
			//私聊模式
			fmt.Println(">>>>>>私聊模式,输入exit退出")
			clt.PrivateChat()
		case 3:
			//更新用户名
			clt.UpdateName()
		}
	}
}

func (clt *Client) UpdateName() bool {
	fmt.Println(">>>>请输入用户名：")
	fmt.Scanln(&clt.ClientName)

	sendMsg := "rename|" + clt.ClientName + "\n"
	_, err := clt.conn.Write([]byte(sendMsg))
	if err != nil {
		fmt.Println("conn.Write error:", err)
		return false
	}
	return true
}

func (clt *Client) DealResponse() {
	//不断阻塞等待套接字conn的数据，并copy到stdout中
	//一旦clt.conn有数据，就直接拷贝到stdout中，是永久阻塞的
	io.Copy(os.Stdout, clt.conn)
	//上面的一行代码等价于下面的代码
	// for {
	// 	buf := make([]byte, 4096)
	// 	clt.conn.Read(buf)
	// 	fmt.Println(buf)
	// }
}

func (clt *Client) PublicChat() {
	var chatMsg string
	fmt.Println(">>>>>请输入聊天内容:")
	fmt.Scanln(&chatMsg)

	for chatMsg != "exit" {
		//消息不为空则发送给服务器
		if len(chatMsg) != 0 {
			sendMsg := chatMsg + "\n"
			_, err := clt.conn.Write([]byte(sendMsg))
			if err != nil {
				fmt.Println("conn.Write error:", err)
				break
			}
		}

		chatMsg = ""
		fmt.Println(">>>>>请输入聊天内容:")
		fmt.Scanln(&chatMsg)
	}
}

// 查询在线用户
func (clt *Client) SelectUsers() {
	sendMsg := "who\n"
	_, err := clt.conn.Write([]byte(sendMsg))
	if err != nil {
		fmt.Println("conn.Write error:", err)
		return
	}
}

func (clt *Client) PrivateChat() {
	clt.SelectUsers()

	var remoteName string
	var chatMsg string
	fmt.Println(">>>>>>请输入聊天对象(用户名),输入exit退出:")
	fmt.Scanln(&remoteName)

	for remoteName != "exit" {
		fmt.Println(">>>>>>请输入消息内容,输入exit退出:")
		fmt.Scanln(&chatMsg)
		for chatMsg != "exit" {
			//消息不为空则发送给服务器
			if len(chatMsg) != 0 {
				sendMsg := "to|" + remoteName + "|" + chatMsg + "\n\n"
				_, err := clt.conn.Write([]byte(sendMsg))
				if err != nil {
					fmt.Println("conn.Write error:", err)
					break
				}
			}

			chatMsg = ""
			fmt.Println(">>>>>请输入消息内容,输入exit退出:")
			fmt.Scanln(&chatMsg)
		}

		clt.SelectUsers()
		fmt.Println(">>>>>>请输入聊天对象(用户名),输入exit退出:")
		fmt.Scanln(&remoteName)
	}
}

var serverIP string
var serverPort int

func init() {
	flag.StringVar(&serverIP, "ip", "127.0.0.1", "设置服务器IP地址(默认为127.0.0.1)")
	flag.IntVar(&serverPort, "port", 8888, "设置服务器端口号(默认为8888)")
}

func main() {
	//命令行解析
	flag.Parse()

	client := NewClient("127.0.0.1", 8888)
	if client == nil {
		fmt.Println(">>>>>>>>链接服务器失败>>>>>>>>")
		return
	}

	fmt.Println(">>>>>>>>链接服务器成功>>>>>>>>>")

	//单独开启一个goroutine，去处理server的回执消息
	go client.DealResponse()

	//启动客户端业务
	client.Run()
}
