package main

import (
	"fmt"
	"net"
	"strings"
)

func HandlerConnect(conn net.Conn)  {
	defer conn.Close()

	addr := conn.RemoteAddr()
	fmt.Printf("客户端：%s 连接成功\n",addr)

	buf := make([]byte,1024*4)

	//多次读取
	for {
		n,err := conn.Read(buf)

		if err != nil{
			fmt.Println("读取失败:",err)
			return
		}

		if n == 0{
			fmt.Printf("客户端：%s 断开连接\n",addr)
			return
		}


		str := string(buf[:n])

		if "exit\n" == str{
			fmt.Printf("客户端：%s 发出退出指令\n",addr)
			return
		}

		fmt.Printf("服务端读取 : %s 发送的数据 : %s",addr,str)

		str2 := strings.ToUpper(str)
		conn.Write([]byte(str2))

	}

}

func main() {

	listen,err := net.Listen("tcp","127.0.0.1:8282")

	if err != nil{
		fmt.Println("tcp创建失败:",err)
		return
	}

	defer listen.Close()
	fmt.Println("TCP服务端等待客户端连接：")


	//多个连接
	for{
		conn,err := listen.Accept()
		if err != nil{
			fmt.Println("监听失败:",err)
			return
		}

		go HandlerConnect(conn)
	}


}
