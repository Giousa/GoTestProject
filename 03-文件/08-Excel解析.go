package main

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
)

func main() {
	fmt.Println("开始Excel测试：")
	//parseExcel("/Users/zhangmengmeng/Documents/CodeResource/go_project/src/testProject/03-文件/excel_file.xlsx")
}

func parseExcel(fileName string) {

	fmt.Println("开始解析Excel文件: url = ",fileName)

	f,err := excelize.OpenFile(fileName)
	if err != nil{
		fmt.Println("文件解析失败,err = ",err)
		return
	}

	rows,err := f.GetRows("省市县区")
	if err != nil{
		fmt.Println("查询失败！！")
		return
	}
	for _,row := range rows{
		//fmt.Println("index = ",index)
		fmt.Println("row = ",row)
		length := len(row)
		if length == 3{
			fmt.Println("省 ： ",row[0])
			fmt.Println("市 ： ",row[1])
			fmt.Println("区 ： ",row[2])
			fmt.Println("-----------------------")
		}
	}


}