package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {

	//测试url  http://pic.netbian.com/4kdongman/index_3.html

	resp,err := http.Get("http://pic.netbian.com/4kdongman/index_3.html")

	if err != nil{
		fmt.Println("Get err ",err)
		return
	}
	defer resp.Body.Close()


	buf := make([]byte,1024*4)

	var result string
	for{
		n,err := resp.Body.Read(buf)

		if n == 0{
			fmt.Println("----Read finish----")
			break
		}

		if err != nil && err != io.EOF{
			fmt.Println("read err ",err)
			return
		}

		result += string(buf[:n])
	}


	fmt.Println(result)

}
