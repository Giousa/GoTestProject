package main

import (
	"fmt"
	"time"
)

func producer(out chan<- int)  {
	for i := 0; i < 10; i++ {
		fmt.Println("生产：",i)
		out <- i
	}

	close(out)
}

func consumer(in <-chan int)  {
	for num := range in{
		fmt.Println("消费者拿到数据：",num)
		time.Sleep(time.Second)
	}
}


func main() {

	ch := make(chan int)

	go producer(ch)//子go 生产者
	consumer(ch)//主 消费者
	
}
