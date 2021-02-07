package main

import (
	"fmt"
	"io/ioutil"
	"mahonia"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"testProject/01-爬虫/db"
	"testProject/01-爬虫/models"
	"time"
)

//http://www.dybz2.net/shuku/7-monthvisit-0-1.html
//http://www.dybz2.net/0/30/
//http://www.dybz2.net/0/30/705.html
var basUrl6 = "http://www.dybz2.net"
var engin6 = db.InitMysqlEngin()
var enc6 mahonia.Decoder


func main() {

	enc6 = mahonia.NewDecoder("gbk")

	//TODO 明天继续
	sql := "select * from das_novel_info where type = '其他类别' limit 421,1000"

	queryList,err := engin6.QueryString(sql)
	if err != nil{
		fmt.Println("query err ",err)
		return
	}

	for k,v := range queryList{
		url := v["url"]
		id,_ := strconv.Atoi(v["id"])

		fmt.Printf("正在存储第 %v 页 数据\n",k+421)
		fmt.Println(url)

		requestKtdmQtlxInfo(url,id)

	}

	fmt.Println("---爬虫结束---")
}


func requestKtdmQtlxInfo(url string,id int){
	resp,_ := http.Get(url)
	defer resp.Body.Close()

	bys,_ := ioutil.ReadAll(resp.Body)

	result := string(bys)

	//fmt.Println(result)

	//<li><a href="/2/2407/45661.html">发发发</a></li>
	ret := regexp.MustCompile(`<li><a href="(.*?)">(.*?)</a></li>`)
	ktdmList := ret.FindAllStringSubmatch(result,-1)


	for page,v := range ktdmList{

		//if page >= 1{
		//	return
		//}
		newUrl := basUrl6+v[1]
		//将标题转为gbk格式
		title := enc6.ConvertString(v[2])
		fmt.Printf("【page】%d - 【标题】%s【url】%s\n",page+1,title,newUrl)

		resultDetl := requestNovelPage(newUrl)

		if resultDetl == ""{
			continue
		}

		novelContent := ""
		novelContent += pickUpCh(resultDetl)

		//查询页码和url
		retDetl := regexp.MustCompile(`a href="(.*?)">(.*?)</a>`)
		novelDetlList := retDetl.FindAllStringSubmatch(resultDetl,-1)
		for _,v := range novelDetlList{
			subUrl := url+v[1]
			subPage := v[2]
			if strings.Contains(subPage,"【") && subPage != "【1】"{
				fmt.Printf("【页码】%s【url】%s\n",subPage,subUrl)
				novelContent += pickUpCh(requestNovelPage(subUrl))
			}
		}


		//fmt.Println(novelContent)

		novelDetl := models.DasNovelDetl{
			NovelId: id,
			SubTitle: title,
			SubUrl: url,
			Content: novelContent,
			CreateTime: time.Now(),
		}

		engin6.Insert(novelDetl)
	}

}

func requestNovelPage(url string) string {
	respDetl,_ := http.Get(url)
	defer respDetl.Body.Close()

	detlBys,_ := ioutil.ReadAll(respDetl.Body)
	resultDetl := string(detlBys)

	//fmt.Println(resultDetl)
	//<div class="page-content font-large">

	begin := strings.Index(resultDetl,"<div class=\"page-content font-large\">")
	end := strings.Index(resultDetl,"<div class=\"tuijian\"><span>")

	fmt.Println("begin : ",begin)
	fmt.Println("end : ",end)

	if begin < 0 || end < 0 || begin > end{
		return ""
	}

	//<div class="slide-baidu">

	return resultDetl[begin:end]

	//return resultDetl
}

//将所有中文汉字提取出来
func pickUpCh(content string) string {
	if content == ""{
		return ""
	}

	r := []rune(content)
	//fmt.Println("rune=", r)
	strSlice := []string{}
	cnstr := ""
	for i := 0; i < len(r); i++ {
		//if r[i] <= 40869 && r[i] >= 19968 {
		//	cnstr = cnstr + string(r[i])
		//	strSlice = append(strSlice, cnstr)
		//
		//}

		if r[i] <= 40869 && r[i] >= 19968 {
			cnstr = cnstr + string(r[i])
			strSlice = append(strSlice, cnstr)

		}

		if r[i] == 65292 || r[i] == 3002 || r[i] == 65311 || r[i] == 65281 || r[i] == 65306{
			cnstr = cnstr + string(r[i])
			strSlice = append(strSlice, cnstr)
		}
		//fmt.Println("r[", i, "]=", r[i], "string=", string(r[i]))
	}
	if 0 == len(strSlice) {
		//无中文，需要跳过，后面再找规律
	}
	//fmt.Println("提取出的中文字符串:")
	//fmt.Println(cnstr)
	return cnstr
}
