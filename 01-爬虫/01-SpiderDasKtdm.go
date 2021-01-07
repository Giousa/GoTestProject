package main

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strconv"
	"testProject/01-爬虫/db"
	"testProject/01-爬虫/models"
	"time"
)

//http://aa036.space/das/12/1.htm
//1----371

func main() {

	baseUrl := "http://aa036.space/das/12/"

	isQuit := make(chan bool)

	go func() {
		for i := 359; i <= 371; i++ {
			url := baseUrl+strconv.Itoa(i)+".htm"
			requestKtdmUrl(url)
		}

		isQuit <- true
	}()

	<-isQuit
	close(isQuit)

	fmt.Println("main over")
}

func requestKtdmUrl(url string)  {

	fmt.Println("go程请求url = ",url)
	resp,err := http.Get(url)
	if err != nil{
		fmt.Println("Get err:",err)
		return
	}
	defer resp.Body.Close()

	buf := make([]byte,1024*4)
	var result string
	for{
		n,err := resp.Body.Read(buf)
		if n == 0{
			fmt.Println("读取失败")
			break
		}
		if err != nil && err != io.EOF{
			fmt.Println("Read err:",err)
			break
		}
		//fmt.Println(string(buf[:n]))
		result += string(buf[:n])
	}


	parseResult(result)

}

func parseResult(result string)  {
	if result == ""{
		return
	}
	//fmt.Println(result)
	// <li><a href=/fe/954842.htm target=_blank title=嘘嘘图[30P]><span>12-24</span>嘘嘘图[30P]</a></li>
	ret := regexp.MustCompile(`<li><a href=(.*?) target=_blank title=.*><span>.*</span>(.*?)</a></li>`)
	ktdmList := ret.FindAllStringSubmatch(result,-1)

	for _,v := range ktdmList{
		//fmt.Println(v)
		url := "http://aa036.space"+v[1]
		title := v[2]
		//fmt.Printf("url = %v,title = %v \n",url,title)

		ktmlInfo := models.DasKtdmInfo{
			Url: url,
			Title: title,
			CreateTime: time.Now(),
		}

		engin := db.InitMysqlEngin()
		engin.Insert(ktmlInfo)

	}
}
