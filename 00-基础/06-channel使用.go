package main

import "fmt"

func send(out chan <- int) {
	out <- 111
	close(out)
}

func recv(in <- chan int)  {
	num := <- in
	fmt.Println("读取数据Num = ",num)
}

func main() {

	ch := make(chan int)

	go func() {
		send(ch)
	}()

	recv(ch)




}
