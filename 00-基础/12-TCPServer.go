package main

import (
	"fmt"
	"net"
)

func main() {

	listen,err := net.Listen("tcp","127.0.0.1:8181")

	if err != nil{
		fmt.Println("tcp创建失败:",err)
		return
	}

	defer listen.Close()

	fmt.Println("TCP服务端等待客户端连接：")
	conn,err := listen.Accept()
	if err != nil{
		fmt.Println("监听失败:",err)
		return
	}

	defer conn.Close()

	fmt.Println("客户端连接成功!")
	buf := make([]byte,1024*4)
	n,err := conn.Read(buf)
	if err != nil{
		fmt.Println("读取失败:",err)
		return
	}
	fmt.Println("TCP server read : ",string(buf[:n]))

	conn.Write([]byte("服务端已收到数据"))
	
}
