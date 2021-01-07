package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var rwMutex sync.RWMutex

func readGo(in <-chan int,idx int)  {
	for{
		rwMutex.RLock()
		num := <-in
		fmt.Printf("-------第 %d 读go程，读取：%d\n",idx,num)
		rwMutex.RUnlock()
	}

}

func writeGo(out chan<- int,idx int)  {

	for{
		num := rand.Intn(1000)
		rwMutex.Lock()
		out <- num
		fmt.Printf("第 %d 写go程，写入：%d\n",idx,num)
		time.Sleep(time.Microsecond*300)
		rwMutex.Unlock()
	}

}


func main() {

	rand.Seed(time.Now().UnixNano())


	ch := make(chan int)

	for i := 0; i < 5; i++ {
		go readGo(ch,i+1)
	}

	for i := 0; i < 5; i++ {
		go writeGo(ch,i+100)
	}

	for{
		;
	}
}
