package main

import (
	"fmt"
	"time"
)

func main() {


	fmt.Println("time : ",time.Now())
	myTicker := time.NewTicker(time.Second*2)

  	go func() {

  		for{
			nowTime := <- myTicker.C
			fmt.Println("nowTime = ",nowTime)
		}

	}()

	for{
		;
	}

}
