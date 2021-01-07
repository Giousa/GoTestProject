package main

import (
	"fmt"
	"net"
	"time"
)

type Client struct {
	C chan string
	Name string
	Addr string
}

//全局map，存储用户
var onlineMap map[string]Client

//全局channel，传递用户消息
var messageCh = make(chan string)

func Manager()  {
	//初始化map
	onlineMap = make(map[string]Client)

	//循环读取数据
	for{
		//监听全局channel中是否有数据，若是有，就存储至map，否则就阻塞
		msg := <- messageCh

		//循环发送消息给所有在线用户
		for _,v := range onlineMap{
			//fmt.Println("2")
			v.C <- msg
		}
	}
}


func WriteMsgToClient(client Client,conn net.Conn)  {
	for msg:= range client.C{
		//fmt.Println("3")
		conn.Write([]byte(msg))
	}
}

func MakeMsg(client Client,msg string) (buf string) {
	buf = fmt.Sprintf("用户[%v] (%v) : %v\n",client.Addr,client.Name,msg)
	return
}

func HandlerConnectFunc(conn net.Conn)  {
	defer conn.Close()

	//判断用户是否活跃
	isActive := make(chan bool)

	//获取用户信息
	netAddr := conn.RemoteAddr().String()

	//创建新连接用户的结构体， 默认用户名称： ip+port
	client := Client{
		make(chan string),
		netAddr,
		netAddr,
	}

	onlineMap[netAddr] = client

	//创建专门用来给当前用户发送消息的go程
	go WriteMsgToClient(client,conn)

	//发送用户上线通知
	//msg := fmt.Sprintf("[%v] 上线了",netAddr)
	//fmt.Println(msg)
	//messageCh <- msg
	messageCh <- MakeMsg(client,"加入聊天室")

	//坚持用户是否退出
	isQuit := make(chan bool)


	//go程，专门处理用户发送的消息
	go func() {
		buf := make([]byte,1024*4)
		for{
			n,err := conn.Read(buf)
			if n == 0{
				isQuit <- true
				fmt.Printf("用户:%v 退出\n",netAddr)
				return
			}

			if err != nil{
				fmt.Println("Read err = ",err)
				return
			}

			//将服务器读取的用户消息，也就是用户发来的消息，发送到全局channe，并全局广播
			//[115 115 10]
			//因为发来的消息，最后一个是\n空格
			str := string(buf[:n-1])
			if "users" == str && len(str) == 5{
				//查询用户列表
				conn.Write([]byte("当前在线用户列表:\n"))

				for _,v := range onlineMap{
					addr := v.Addr
					if addr != netAddr{
						msg := fmt.Sprintf("在线用户:[%v] (%v) \n",addr,v.Name)
						conn.Write([]byte(msg))
					}

				}
			}else if len(str) >= 8 && str[:6] == "rename"{
				//rename|不笑猫
				//改名操作
				newName := str[7:]
				//for _,v := range onlineMap{
				//	if v.Addr == netAddr{
				//		v.Name = newName
				//	}
				//}
				client.Name = newName
				onlineMap[netAddr] = client

			}else {
				messageCh <- MakeMsg(client,str)
			}

			isActive <- true

		}
	}()

	//保证不退出
	for{
		select{
			case <-isQuit:
				delete(onlineMap,client.Addr)
				messageCh <- MakeMsg(client,"退出聊天室")
				fmt.Println("执行了退出指令")
				return

			case <-isActive:
					//活跃
					//目的：重置下面的计时器

			case <-time.After(time.Second * 10):
				delete(onlineMap,client.Addr)
				messageCh <- MakeMsg(client,"超时退出聊天室")
				fmt.Println("超时，退出聊天室")
				return
		}
	}

}

func main() {

	listen,err := net.Listen("tcp","127.0.0.1:8000")

	if err != nil{
		fmt.Println("listen err = ",err)
		return
	}

	defer listen.Close()

	fmt.Println("服务器开启：")

	//创建管理者go程，管理map和channel
	go Manager()

	for{
		conn,err := listen.Accept()

		if err != nil{
			fmt.Println("accept err = ",err)
			return
		}

		go HandlerConnectFunc(conn)

	}
	
}
