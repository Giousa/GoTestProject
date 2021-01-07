package main

import (
	"fmt"
	"net"
	"os"
)

func main() {

	conn,err := net.Dial("tcp","127.0.0.1:8282")
	if err != nil{
		fmt.Println("客户端连接服务端失败:",err)
		return
	}

	defer conn.Close()

	//获取键盘输入，将输入数据发送给服务器
	go func() {
		str := make([]byte,1024*4)
		for{
			n,err := os.Stdin.Read(str)
			if err != nil{
				fmt.Println("录入数据异常:",err)
				continue
			}

			conn.Write(str[:n])
		}
	}()

	//读取服务器发来数据
	buf := make([]byte,1024*4)
	for{
		n,err := conn.Read(buf)

		if n == 0{
			fmt.Printf("客户端退出")
			return
		}

		if err != nil{
			fmt.Println("读取失败:",err)
			return
		}


		str := string(buf[:n])
		fmt.Printf("客户端读取 :%s",str)

	}
}
