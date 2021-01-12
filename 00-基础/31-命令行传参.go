package main

import (
	"flag"
	"fmt"
)

//定义命令行参数
var name = flag.String("name", "", "input your name")
var age = flag.Int("Age", 0, "input your age")
func main() {
	//解析命令行参数
	flag.Parse()
	//输出命令行参数
	fmt.Println(*name, *age)
}