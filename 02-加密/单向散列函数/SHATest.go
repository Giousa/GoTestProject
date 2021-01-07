package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

func main() {
	//7ff373660c1beb229220e93bdc3d5c14435607f9  文件大小：40   加密后字节：40字节
	getSha1("file/encryption/text.txt")
	//2feade020b834e8ffbb6e62b4c82b95410b3d946  文件大小：497  加密后字节： 40字节
	getSha1("file/encryption/private.pem")
}

//使用sha1计算文件指纹
func getSha1(src string) string {
	// 1. 打开文件
	fp, err := os.Open(src)
	if err != nil {
		fmt.Println("文件打开失败,err = ",err)
		return "文件打开失败"
	}
	// 2. 创建基于sha1算法的Hash对象
	myHash := sha1.New()
	// 3. 将文件数据拷贝给哈希对象
	num, err := io.Copy(myHash, fp)
	if err != nil {
		fmt.Println("拷贝文件失败,err = ",err)
		return "拷贝文件失败"
	}
	fmt.Println("文件大小: ", num)
	// 4. 计算文件的哈希值
	tmp1 := myHash.Sum(nil)
	// 5. 数据格式转换
	result := hex.EncodeToString(tmp1)
	fmt.Println("sha1: ", result)

	return result
}