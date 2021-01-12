package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"testProject/01-爬虫/db"
	"testProject/01-爬虫/models"
	"time"
)

//http://aa036.space/das/18/1.htm

var basUrl2 = "http://aa036.space"
var engin = db.InitMysqlEngin()

func main() {

	//isQuit := make(chan bool)
	//
	////1---371
	//go func() {
	//	for i := 1; i <= 371; i++ {
	//		url := basUrl2 + "das/18/" + strconv.Itoa(i) + ".htm"
	//		fmt.Println("开始爬取：",url)
	//		requestNovelInfo(url)
	//	}
	//
	//	isQuit <- true
	//}()
	//
	//
	//<-isQuit

	//371
	for i := 1; i <= 371; i++ {
		url := basUrl2 + "/das/18/" + strconv.Itoa(i) + ".htm"
		fmt.Println("开始爬取：",url)
		requestNovelInfo2(url)
	}
	fmt.Println("---爬虫结束---")
}


func requestNovelInfo2(url string){
	resp,_ := http.Get(url)

	defer resp.Body.Close()

	bys,_ := ioutil.ReadAll(resp.Body)
	result := string(bys)

	//fmt.Println(result)

	//<li><a href=/fe/933516.htm target=_blank title=【绝情剑】（完）><span>11-04</span>【绝情剑】（完）</a></li>
	ret := regexp.MustCompile(`<li><a href=(.*?) target=_blank title=.*><span>.*</span>(.*?)</a></li>`)
	ktdmList := ret.FindAllStringSubmatch(result,-1)
	for _,v := range ktdmList{
		//fmt.Println("page = ",page+1)//页码
		//fmt.Println(v[1])//url
		//fmt.Println(v[2])//title
		url := basUrl2+v[1]
		title := v[2]
		fmt.Printf("【标题】%s【url】%s\n",title,url)
		novel := models.DasNovelInfo{
			Url: url,
			Title: v[2],
			Type: "武侠古典",
			CreateTime: time.Now(),
		}

		//fmt.Println(novel)

		//录入数据库
		//fmt.Println("开始录入数据库::")
		engin.Insert(novel)

	}
}