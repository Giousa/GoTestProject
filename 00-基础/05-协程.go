package main

import (
	"fmt"
	"runtime"
)

func test()  {

	fmt.Println("333333333333333")

	runtime.Goexit()
	fmt.Println("444444444444444")
}

func main() {

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("111111111111111")
			test()
			fmt.Println("222222222222222")
		}
	}()

	for{
		;
	}


}
