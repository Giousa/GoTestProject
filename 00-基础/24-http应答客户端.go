package main

import (
	"fmt"
	"net"
)

func main() {
	conn,_ := net.Dial("tcp","127.0.0.1:8000")
	defer conn.Close()

	httpRequest := "GET /test HTTP/1.1\r\nHost: localhost:8000\r\n\r\n"
	conn.Write([]byte(httpRequest))

	buf := make([]byte,1024*4)
	n,_ := conn.Read(buf)
	fmt.Println(string(buf[:n]))
}
