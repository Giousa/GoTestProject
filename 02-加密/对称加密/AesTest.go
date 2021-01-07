package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
)

/**
和DES—样，AES算法也是由多个轮所构成的，下图展示了每一轮的大致计算步骤。
DES使用Feistel网络作为其基本结构，而AES没有使用Feistel网络，而是使用了SPN Rijndael的输人分组为128比特，也就是16字节。
首先，需要逐个字节地对16字节的输入数据进行SubBytes处理。
所谓SubBytes,就是以每个字节的值（0～255中的任意值）为索引，从一张拥有256个值的替换表（S-Box）中查找出对应值的处理，
也是说，将一个1字节的值替换成另一个1字节的值。
 */
func main() {
	AESText()
}

// AES加密
func AESEncrypt(src, key []byte) []byte{
	// 1. 创建一个使用AES加密的块对象
	block, err := aes.NewCipher(key)
	if err != nil{
		panic(err)
	}
	// 2. 最后一个分组进行数据填充
	src = AESPKCS5Padding(src, block.BlockSize())
	// 3. 创建一个分组为链接模式, 底层使用AES加密的块模型对象
	blockMode := cipher.NewCBCEncrypter(block, key[:block.BlockSize()])
	// 4. 加密
	dst := src
	blockMode.CryptBlocks(dst, src)
	return dst
}

// AES解密
func AESDecrypt(src, key []byte) []byte{
	// 1. 创建一个使用AES解密的块对象
	block, err := aes.NewCipher(key)
	if err != nil{
		panic(err)
	}
	// 2. 创建分组为链接模式, 底层使用AES的解密模型对象
	blockMode := cipher.NewCBCDecrypter(block, key[:block.BlockSize()])
	// 3. 解密
	dst := src
	blockMode.CryptBlocks(dst, src)
	// 4. 去掉尾部填充的字
	dst = AESPKCS5UnPadding(dst)
	return dst
}

// 使用pks5的方式填充
func AESPKCS5Padding(ciphertext []byte, blockSize int) []byte{
	// 1. 计算最后一个分组缺多少个字节
	padding := blockSize - (len(ciphertext)%blockSize)
	// 2. 创建一个大小为padding的切片, 每个字节的值为padding
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	// 3. 将padText添加到原始数据的后边, 将最后一个分组缺少的字节数补齐
	newText := append(ciphertext, padText...)
	return newText
}

// 删除pks5填充的尾部数据
func AESPKCS5UnPadding(origData []byte) []byte{
	// 1. 计算数据的总长度
	length := len(origData)
	// 2. 根据填充的字节值得到填充的次数
	number := int(origData[length-1])
	// 3. 将尾部填充的number个字节去掉
	return origData[:(length-number)]
}

func AESText() {
	// 加密   go里面采用16字节，其他语言，可以使用24和32字节加密
	key := []byte("1111111111111111")
	//key := []byte("111111111111111111111111")
	//key := []byte("11111111111111111111111111111111")
	result := AESEncrypt([]byte("床前明月光, 疑是地上霜. 举头望明月, 低头思故乡."), key)
	fmt.Println("加密后展示数据：")
	//YecPlzKvnn644QeK8Nh8DV+KmwvZfjAXWtfsXilLN5AbfjmSESdwNsuG9TCKagJmKVxC5CQw0SdTmV5QdSdI9wQCCQALP4C5vspwEVhb9pQ=
	fmt.Println(base64.StdEncoding.EncodeToString(result))
	// 解密
	result = AESDecrypt(result, key)
	fmt.Println("解密之后的数据: ", string(result))
}
