package main

import (
	"fmt"
	"strconv"
)

//http://aa036.space/das/18/1.htm

var basUrl = "http://aa036.space/"

func main() {

	isQuit := make(chan bool)

	//1---371
	go func() {
		for i := 1; i <= 371; i++ {
			url := basUrl + "das/18/" + strconv.Itoa(i) + ".htm"
			fmt.Println("开始爬取：",url)
			requestNovelInfo(url)
		}

		isQuit <- true
	}()


	<-isQuit
	fmt.Println("---爬虫结束---")
}


func requestNovelInfo(url string){

}