package main

import (
	"fmt"
	"net"
	"os"
)

func errfunc(err error,info string)  {
	if err != nil{
		fmt.Println(info,err)

		os.Exit(1)
	}
}

func main() {

	listen,err := net.Listen("tcp","127.0.0.1:8000")
	errfunc(err,"监听异常")

	defer listen.Close()

	conn,err := listen.Accept()
	errfunc(err,"Accept err")

	defer conn.Close()

	buf := make([]byte,1024*4)
	n,err := conn.Read(buf)
	if n == 0{
		return
	}

	errfunc(err,"读取异常")

	fmt.Println(string(buf[:n]))


}
