package main

import (
	"fmt"
	"net/http"
	"os"
)

func openSendFile(fName string, w http.ResponseWriter)  {
	pathFileName := "/Users/zhangmengmeng/Desktop/images"+fName
	f,err := os.Open(pathFileName)
	if err != nil {
		w.Write([]byte("文件不存在"))
		return
	}

	defer f.Close()

	buf := make([]byte,1024*4)
	for{
		n,err := f.Read(buf)
		if n == 0{
			return
		}

		if err != nil{
			return
		}

		w.Write(buf[:n])
	}
}


func myHandler(w http.ResponseWriter,r *http.Request){
	fmt.Println("客户端请求：",r.URL)
	openSendFile(r.URL.String(),w)
}


func main() {

	http.HandleFunc("/",myHandler)

	http.ListenAndServe("127.0.0.1:8000",nil)

}
