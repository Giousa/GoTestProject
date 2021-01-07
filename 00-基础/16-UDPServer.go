package main

import (
	"fmt"
	"net"
	"time"
)

func main() {

	udpAddr,err := net.ResolveUDPAddr("udp","127.0.0.1:8003")
	if err != nil{
		fmt.Println("ResolveUDPAddr err:",err)
		return
	}
	udpConn,err := net.ListenUDP("udp",udpAddr)
	if err != nil{
		fmt.Println("ListenUDP err:",err)
		return
	}

	defer udpConn.Close()

	fmt.Println("UDP服务器启动成功：")

	//读数据
	buf := make([]byte,1024*4)
	n,clientAddr,err := udpConn.ReadFromUDP(buf)
	if err != nil{
		fmt.Println("ReadFromUDP err:",err)
		return
	}
	readStr := string(buf[:n])
	fmt.Printf("客户端：%v:%v 发送来的数据:%v",clientAddr.IP,clientAddr.Port,readStr)

	//写数据
	dayTime := time.Now().String()
	_,err = udpConn.WriteToUDP([]byte("时间："+dayTime+"\n"),clientAddr)
	if err != nil{
		fmt.Println("WriteToUDP err:",err)
		return
	}
}
