package main

import "fmt"

func main() {

	fmt.Println("指针的应用")

	var a = 10
	fmt.Println(fmt.Sprint(a))

	var p*int = &a
	fmt.Println(p)

	//a = 100
	//fmt.Println(fmt.Sprint(a))
	//fmt.Println(p)

	*p = 100
	fmt.Println(fmt.Sprint(a))
	fmt.Println(p)

	var b = 20
	p = &b
	fmt.Println(p)

	
}
