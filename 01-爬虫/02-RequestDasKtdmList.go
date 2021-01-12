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

func main() {

	engin := db.InitMysqlEngin()
	//已经收录100
	sql := "select * from das_ktdm_info limit 7001,1000"

	queryList,err := engin.QueryString(sql)
	if err != nil{
		fmt.Println("query err ",err)
		return
	}

	isQuit := make(chan bool)
	go func() {
		for k,v := range queryList{
			url := v["url"]
			id,_ := strconv.Atoi(v["id"])

			fmt.Printf("正在存储第 %v 页 数据\n",k+7001)
			fmt.Println(url)

			requestKtdmInfo(url,id)

		}
		isQuit <- true
	}()

	<- isQuit
	fmt.Println("---main over---")
}

func requestKtdmInfo(url string, id int) {
	resp,err := http.Get(url)
	if err != nil{
		fmt.Println("Get err ",err)
		return
	}

	defer resp.Body.Close()

	buf := make([]byte,1024*4)
	var result string
	for{
		n,err := resp.Body.Read(buf)
		if n == 0{
			//fmt.Println("读取失败")
			break
		}
		if err != nil && err != io.EOF{
			fmt.Println("Read err:",err)
			break
		}
		result += string(buf[:n])
	}
	
	parseKtdmResult(result,id)

}

func parseKtdmResult(result string,id int)  {
	if result == ""{
		return
	}
	//fmt.Println(result)
	ret := regexp.MustCompile(`<img src="(.*?)" alt=`)
	ktdmList := ret.FindAllStringSubmatch(result,-1)

	for index,v := range ktdmList{
		url := v[1]
		//fmt.Printf("url = %v,page = %v\n",url,index)

		ktdmDetl := models.DasKtdmDetl{
			KtdmId: id,
			PicUrl: url,
			Page: index+1,
			CreateTime: time.Now(),
		}

		engin := db.InitMysqlEngin()
		engin.Insert(ktdmDetl)

	}
}
