package main

import (
	"fmt"
	"strings"
)

func main() {

	//将字符串转为map，并统计每个单词出现的次数
	str := "I love my work and I love my family too"

	mRet := workCountFun(str)

	fmt.Println(mRet)

}

func workCountFun(str string) map[string]int {

	m := make(map[string]int)

	//将字符串，拆分成字符串切片
	s := strings.Fields(str)
	fmt.Println(s)

	//遍历
	for k,v := range s{
		fmt.Println(k)
		//fmt.Println(v)
		//m[v] = k+1
		if _,has := m[v] ; has{
			m[v] = m[v]+1
		}else{
			m[v] = 1
		}
	}

	return m
}
