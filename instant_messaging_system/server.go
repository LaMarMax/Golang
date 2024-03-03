package main

import (
	"fmt"
	"io"
	"net"
	"sync"
	"time"
)

type Server struct {
	IP   string
	Port int

	//在线用户列表
	OnlineMap map[string]*User
	//读写锁
	mapLock sync.RWMutex //go中同步的机制全在sync包中

	//消息广播的channel
	Message chan string
}

// 创建一个server的接口
func NewServer(ip string, port int) *Server {
	server := &Server{
		IP:        ip,
		Port:      port,
		OnlineMap: make(map[string]*User),
		Message:   make(chan string),
	}

	return server
}

// 不断监听Message channel的goroutine，一旦有消息就发送给全部的User
func (svr *Server) ListenMessage() {
	for {
		//从Message channel中读取数据
		msg := <-svr.Message

		//遍历onlinemap，将消息发送给所有的在线User
		svr.mapLock.Lock()
		for _, cli := range svr.OnlineMap {
			cli.C <- msg
		}
		svr.mapLock.Unlock()
	}
}

// 广播消息的方法
func (svr *Server) BroadCast(user *User, msg string) {
	sendMsg := "[" + user.Addr + "]" + user.Name + ":" + msg

	svr.Message <- sendMsg
}

func (svr *Server) Handler(conn net.Conn) {
	//创建新用户
	user := NewUser(conn, svr)

	//用户上线
	user.Online()

	//监听用户是否活跃的channel
	isLive := make(chan bool)

	//接收用户发送的消息
	go func() {
		buf := make([]byte, 4096)
		for {
			//从链接中读取消息
			n, err := user.conn.Read(buf)
			//如果读取的消息长度为0，就代表当前用户已下线
			if n == 0 {
				user.Offline()
				return
			}

			//	err不是文件末尾
			if err != nil && err != io.EOF {
				fmt.Println("Conn read error:", err)
				return
			}

			// 读取消息成功，提取用户消息，去除\n，然后广播
			msg := string(buf[:n-1])
			user.DoMessage(msg)

			//每当用户发一次消息，就代表活跃
			isLive <- true
		}
	}()

	//当前Handler阻塞，否则go程就结束了
	for {
		select {
		case <-isLive:
			//当从isLive中读到数据时，下面的case也会执行一次，就重置了定时器
			//不作任何处理，只是为了更新定时器
		case <-time.After(60 * time.Second):
			//该case触发，表明超时
			//将当前的User强制关闭
			user.SendMsg("你被踢了")

			//销毁用户资源
			close(user.C)
			//关闭链接
			conn.Close()
			//退出当前Handler
			return //或者runtime.Goexit()
		}
	}
}

// 启动服务器的接口
func (svr *Server) Start() {
	// socket listen
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", svr.IP, svr.Port))
	if err != nil {
		fmt.Println("listen error :", err)
		return
	}
	//close listen socket
	defer listener.Close()

	//循环监听Message的go程
	go svr.ListenMessage()

	//accept
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("accpet error :", err)
			continue
		}

		//accept成功后，代表有个用户上线了
		// do handler
		go svr.Handler(conn)
	}
}
