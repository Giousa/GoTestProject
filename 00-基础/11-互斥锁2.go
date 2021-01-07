package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

var mutex2 sync.Mutex

func main() {

	baseUrl := "http://aa036.space/das/12/"

	//isQuit := make(chan bool)
	//go func() {
	//	for i := 1; i <= 100; i++ {
	//		url := baseUrl+strconv.Itoa(i)+".htm"
	//		parseKtdmUrl(url)
	//	}
	//
	//	isQuit <- true
	//}()
	//
	//<-isQuit

	for i := 0; i < 5; i++ {
		go buildMultyGo2(baseUrl,i)
	}

	for{
		;
	}

	fmt.Println("main over")
}

func buildMultyGo2(baseUrl string,index int)  {
	mutex2.Lock()
	for i := 1; i <= 10; i++ {
		url := baseUrl+strconv.Itoa(i+10*index)+".htm"
		parseKtdmUrl2(url)
	}
	mutex2.Unlock()
}

func parseKtdmUrl2(url string)  {

	fmt.Println("go  url = ",url)
	time.Sleep(time.Second*2)

}

