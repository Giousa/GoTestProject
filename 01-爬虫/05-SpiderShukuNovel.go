package main

import (
	"fmt"
	"io/ioutil"
	"mahonia"
	"net/http"
	"regexp"
	"strconv"
	"testProject/01-爬虫/db"
	"testProject/01-爬虫/models"
	"time"
)

//http://www.dybz2.net/shuku/7-monthvisit-0-1.html
//http://www.dybz2.net/0/30/
//http://www.dybz2.net/0/30/705.html
var basUrl5 = "http://www.dybz2.net"
var engin5 = db.InitMysqlEngin()
var enc5 mahonia.Decoder

func main() {

	enc5 = mahonia.NewDecoder("gbk")

	for i := 1; i <= 673; i++ {
		url := basUrl5 + "/shuku/7-monthvisit-0-" + strconv.Itoa(i) + ".html"
		fmt.Println("开始爬取：",url)
		requestNovelInfo5(url)
	}
	fmt.Println("---爬虫结束---")
}


func requestNovelInfo5(url string){
	resp,_ := http.Get(url)

	defer resp.Body.Close()

	bys,_ := ioutil.ReadAll(resp.Body)


	result := string(bys)

	//fmt.Println(result)

	//<a class="name" href="/2/2011/">大陆演艺圈艳史</a>
	//<a class="name" href="/3/3611/">少年的欲望</a>
	ret := regexp.MustCompile(`<a class="name" href="(.*?)">(.*?)</a>`)
	ktdmList := ret.FindAllStringSubmatch(result,-1)
	for _,v := range ktdmList{
		//fmt.Println("page = ",page+1)//页码
		//fmt.Println(v[1])//url
		//fmt.Println(v[2])//title
		url := basUrl5+v[1]
		//将标题转为gbk格式
		title := enc5.ConvertString(v[2])
		fmt.Printf("【标题】%s【url】%s\n",title,url)
		novel := models.DasNovelInfo{
			Url: url,
			Title: title,
			Type: "其他类别",
			CreateTime: time.Now(),
		}
		engin5.Insert(novel)

	}
}