package main

import (
	"fmt"
	"os"
)

func main() {

	//获取命令行参数
	list := os.Args

	fmt.Println("------------------")
	//if len(list) != 2{
	//	fmt.Println("格式为xxx")
	//	return
	//}
	fmt.Println(list)

	
}
