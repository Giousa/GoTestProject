package main

import (
	"fmt"
	"sync"
	"time"
)

//type Mutex struct {
//	state int32
//	sema  uint32
//}
//创建一个互斥量，也就是互斥锁
//刚创建，此时state = 0  未加锁 锁只有一把(一个进程中)
var mutex sync.Mutex

func printer(str string)  {

	//上锁   在访问共享数据之前加锁
	mutex.Lock()

	for _,c := range str{
		fmt.Printf("%c",c)
		time.Sleep(300*time.Millisecond)
	}

	//访问结束，解锁
	mutex.Unlock()
}

func person1()  {
	printer("hello")
}

func person2()  {
	printer("world")
}


func main() {
	go person1()
	go person2()

	for{
		;
	}
}
