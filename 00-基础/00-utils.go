package main

import (
	"fmt"
	"strings"
)

func main() {

	//str := "今天天气真的很好呀"
	str := "today is a good day"
	byt := []byte(str)
	fmt.Println(byt)

	str2 := strings.ToUpper(str)
	fmt.Println(str2)
}
