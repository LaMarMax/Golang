package main

import (
	"fmt"
	"net"
	"strings"
)

type User struct {
	Name string
	Addr string
	C    chan string //用户channel
	conn net.Conn    //与客户端的链接

	svr *Server //当前用户属于哪个Server
}

func NewUser(connect net.Conn, server *Server) *User {
	//以地址名作为用户名
	userAddr := connect.RemoteAddr().String()

	user := &User{
		Name: userAddr,
		Addr: userAddr,
		C:    make(chan string),
		conn: connect,
		svr:  server,
	}

	//启动监听当前User channel消息的goroutine
	go user.ListenMessage()

	return user
}

// 监听当前User channel的方法，一旦有消息，就直接发送给客户端
func (usr *User) ListenMessage() {
	for {
		msg := <-usr.C

		usr.conn.Write([]byte(msg + "\n")) //使用二进制数组形式写入消息
	}
}

// 用户上线
func (usr *User) Online() {
	//用户上线，加入到onlinemap中
	usr.svr.mapLock.Lock()
	usr.svr.OnlineMap[usr.Name] = usr
	usr.svr.mapLock.Unlock()
	//广播当前用户上线消息
	usr.svr.BroadCast(usr, "已上线")
}

// 用户下线
func (usr *User) Offline() {
	//用户上线，加入到onlinemap中
	usr.svr.mapLock.Lock()
	delete(usr.svr.OnlineMap, usr.Name)
	usr.svr.mapLock.Unlock()
	//广播当前用户上线消息
	usr.svr.BroadCast(usr, "已下线")
}

// 给当前User对应的客户端发消息(单发)
func (usr *User) SendMsg(msg string) {
	usr.conn.Write([]byte(msg))
}

// 用户处理消息
func (usr *User) DoMessage(msg string) {
	if msg == "who" {
		//用户想查询所有在线用户
		usr.svr.mapLock.Lock()
		for _, user := range usr.svr.OnlineMap {
			onlineUser := "[" + user.Addr + "]" + user.Name + ":" + "在线...\n"
			usr.SendMsg(onlineUser)
		}
		usr.svr.mapLock.Unlock()
	} else if len(msg) > 7 && msg[:7] == "rename|" {
		//修改指令是 rename|...
		newName := strings.Split(msg, "|")[1]
		//首先查询onlinemap中是否存在该用户名
		_, ok := usr.svr.OnlineMap[newName]
		if ok {
			fmt.Println("当前用户名已经被使用！")
		} else {
			//删除原用户名对应的user，并加入新的用户名
			usr.svr.mapLock.Lock()
			delete(usr.svr.OnlineMap, usr.Name)
			usr.svr.OnlineMap[newName] = usr
			usr.svr.mapLock.Unlock()

			//更新User的用户名
			usr.Name = newName
			usr.SendMsg("您已更改用户名为:" + newName + "\n")
		}
	} else if len(msg) > 4 && msg[:3] == "to|" {
		//消息格式：to|张三|张三你好...
		//获取对方用户名
		remoteName := strings.Split(msg, "|")[1]
		if remoteName == "" {
			usr.SendMsg("私聊消息格式不正确，请使用\"to|张三|张三你好...\"格式\n")
			return
		}
		//根据用户名找到User对象
		remoteUser, ok := usr.svr.OnlineMap[remoteName]
		if !ok {
			usr.SendMsg(remoteName + "用户不存在\n")
			return
		}
		//获取消息内容，通过User对象发送给对方
		privateMsg := strings.Split(msg, "|")[2]
		remoteUser.SendMsg(usr.Name + " : " + privateMsg)
	} else {
		usr.svr.BroadCast(usr, msg)
	}
}
