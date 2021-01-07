package main

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"encoding/base64"
	"fmt"
)

/**
DES是一种将64比特的明文加密成64比特的密文的对称密码算法，
==它的密钥长度是56比特==。
尽管从规格上来说，DES的密钥长度是64比特，但由于每隔7比特会设置一个用于错误检查的比特，因此实质上其密钥长度是56比特。
DES是以64比特的明文（比特序列）为一个单位来进行加密的，这个64比特的单位称为分组。
一般来说，以分组为单位进行处理的密码算法称为分组密码（blockcipher），DES就是分组密码的一种。
DES每次只能加密64比特的数据，如果要加密的明文比较长，就需要对DES加密进行迭代（反复），而迭代的具体方式就称为模式（mode）。
大B -> bit
小b -> byte
秘钥长度(56bit + 8bit)/8 = 8byte 12345678
 */
func main() {
	DESText()
}

//加密
// src -> 要加密的明文
// key -> 秘钥, 大小为: 8byte
func DesEncrypt_CBC(src, key []byte) []byte{
	// 1. 创建并返回一个使用DES算法的cipher.Block接口 block, err := des.NewCipher(key)
	block, err := des.NewCipher(key)
	// 2. 判断是否创建成功
	if err != nil{
		panic(err) }
	// 3. 对最后一个明文分组进行数据填充
	src = PKCS5Padding(src, block.BlockSize())
	// 4. 创建一个密码分组为链接模式的, 底层使用DES加密的BlockMode接口 // 参数iv的长度, 必须等于b的块尺寸
	tmp := []byte("helloAAA")
	blackMode := cipher.NewCBCEncrypter(block, tmp)
	// 5. 加密连续的数据块
	dst := make([]byte, len(src))
	blackMode.CryptBlocks(dst, src)
	fmt.Println("加密之后的数据: ", dst)
	// 6. 将加密数据返回
	return dst
}


//解密
// src -> 要解密的密文
// key -> 秘钥, 和加密秘钥相同, 大小为: 8byte
func DesDecrypt_CBC(src, key []byte) []byte {
	// 1. 创建并返回一个使用DES算法的cipher.Block接口
	block, err := des.NewCipher(key)
	// 2. 判断是否创建成功
	if err != nil{
		panic(err)
	}
	// 3. 创建一个密码分组为链接模式的, 底层使用DES解密的BlockMode接口
	tmp := []byte("helloAAA")
	blockMode := cipher.NewCBCDecrypter(block, tmp)
	// 4. 解密数据
	dst := src
	blockMode.CryptBlocks(src, dst)
	// 5. 去掉最后一组填充的数据
	dst = PKCS5UnPadding(dst)

	// 6. 返回结果
	return dst
}

// 使用pks5的方式填充
func PKCS5Padding(ciphertext []byte, blockSize int) []byte{
	// 1. 计算最后一个分组缺多少个字节
	padding := blockSize - (len(ciphertext)%blockSize)
	// 2. 创建一个大小为padding的切片, 每个字节的值为padding
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	// 3. 将padText添加到原始数据的后边, 将最后一个分组缺少的字节数补齐
	newText := append(ciphertext, padText...)
	return newText
}

// 删除pks5填充的尾部数据
func PKCS5UnPadding(origData []byte) []byte{
	// 1. 计算数据的总长度
	length := len(origData)
	// 2. 根据填充的字节值得到填充的次数
	number := int(origData[length-1])
	// 3. 将尾部填充的number个字节去掉
	return origData[:(length-number)]
}

func DESText() {
	// 加密  8字节
	key := []byte("11111111")
	result := DesEncrypt_CBC([]byte("床前明月光, 疑是地上霜. 举头望明月, 低头思故乡."), key)
	fmt.Println("加密后展示数据：")
	//uTrwJYGpItmOo+wiRz3xRumMm6ysaE50dlfma/NQ0ycOAtozomje9j2U4akTEmtWk3eN/rLgD1d0jLR1NAwHG1/t5EcGA5me
	fmt.Println(base64.StdEncoding.EncodeToString(result))
	// 解密
	result = DesDecrypt_CBC(result, key)
	fmt.Println("解密之后的数据: ", string(result))
}
