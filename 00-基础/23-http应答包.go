package main

import (
	"fmt"
	"net/http"
)

/**
	writer:写回给客户端的数据
	request:从客户端，读到的数据
 */
func handle(writer http.ResponseWriter,request *http.Request)  {
	fmt.Println("回调函数触发")
	writer.Write([]byte("连接成功"))
}

func main() {

	//func HandleFunc(pattern string, handler func(ResponseWriter, *Request)) {
	//参数1：用户访问位置   参数2：回调函数
	http.HandleFunc("/test",handle)

	//func ListenAndServe(addr string, handler Handler) error
	//第二个参数，若是传入nil，调用我们之前定义的函数
	http.ListenAndServe("127.0.0.1:8000",nil)


}
