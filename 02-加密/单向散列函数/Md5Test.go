package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
)

func main() {

	getMD5_1([]byte("今天天气不错，适合出去钓鱼"))
	getMD5_2([]byte("今天天气不错，适合出去钓鱼"))

}

func getMD5_1(str []byte) string {
	// 1. 计算数据的md5
	result := md5.Sum(str)
	fmt.Println(result)
	fmt.Printf("%x\n", result)
	// 2. 数据格式化为16进制格式字符串
	res := fmt.Sprintf("%x", result)
	fmt.Println(res)
	// --- 这是另外一种格式化切片的方式
	res = hex.EncodeToString(result[:])
	fmt.Println("res1 = ",res)
	return  res
}

func getMD5_2(str []byte) string {
	// 1. 创建一个使用MD5校验的Hash对象`
	myHash := md5.New()
	// 2. 通过io操作将数据写入hash对象中
	io.WriteString(myHash, string(str))
	//或者
	//myHash.Write(str)
	// 3. 计算结果
	result := myHash.Sum(nil)
	fmt.Println(result)
	// 4. 将结果转换为16进制格式字符串
	res := fmt.Sprintf("%x", result)
	fmt.Println("res2 = ",res)
	// --- 这是另外一种格式化切片的方式
	res = hex.EncodeToString(result)
	fmt.Println(res)

	return res
}