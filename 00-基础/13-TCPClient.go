package main

import (
	"fmt"
	"net"
)

func main() {

	conn,err := net.Dial("tcp","127.0.0.1:8181")
	if err != nil{
		fmt.Println("客户端连接服务端失败:",err)
		return
	}

	defer conn.Close()

	//主动发送数据给服务器
	conn.Write([]byte("这个是一个简单的小测试"))

	buf := make([]byte,1024*4)
	n,err := conn.Read(buf)
	if err != nil{
		fmt.Println("client读取失败:",err)
		return
	}
	fmt.Println("TCP client read : ",string(buf[:n]))

}
