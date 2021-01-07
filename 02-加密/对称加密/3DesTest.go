package main

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"encoding/base64"
	"fmt"
)

/**
三重DES（triple-DES）是为了增加DES的强度，==将DES重复3次所得到的一种密码算法==，通常缩写为3DES。
明文经过三次DES处理才能变成最后的密文，由于DES密钥的长度实质上是56比特，
因此三重DES的密钥长度就是56×3=168比特,加上用于错误检测的标志位8x3, 共192bit。
三重DES并不是进行三次DES加密（加密-->加密-->加密），而是加密-->解密-->加密的过程。
在加密算法中加人解密操作让人感觉很不可思议，实际上这个方法是IBM公司设计出来的，目的是为了让三重DES能够兼容普通的DES。
当三重DES中所有的密钥都相同时，三重DES也就等同于普通的DES了。
这是因为在前两步加密-->解密之后，得到的就是最初的明文。因此，以前用DES加密的密文，就可以通过这种方式用三重DES来进行解密。
也就是说，三重DES对DES具备向下兼容性。
 */
func main() {
	TripleDESText()
}

// 3DES加密
func TripleDESEncrypt(src, key []byte) []byte {
	// 1. 创建并返回一个使用3DES算法的cipher.Block接口
	block, err := des.NewTripleDESCipher(key)
	if err != nil{
		panic(err)
	}
	// 2. 对最后一组明文进行填充
	src = PKCS5Padding2(src, block.BlockSize())
	// 3. 创建一个密码分组为链接模式, 底层使用3DES加密的BlockMode模型
	blockMode := cipher.NewCBCEncrypter(block, key[:8])
	// 4. 加密数据
	dst := src
	blockMode.CryptBlocks(dst, src)
	return dst
}

// 3DES解密
func TripleDESDecrypt(src, key []byte) []byte {
	// 1. 创建3DES算法的Block接口对象
	block, err := des.NewTripleDESCipher(key)
	if err != nil{
		panic(err)
	}
	// 2. 创建密码分组为链接模式, 底层使用3DES解密的BlockMode模型
	blockMode := cipher.NewCBCDecrypter(block, key[:8])
	// 3. 解密
	dst := src
	blockMode.CryptBlocks(dst, src)
	// 4. 去掉尾部填充的数据
	dst = PKCS5UnPadding2(dst)
	return dst
}


// 使用pks5的方式填充
func PKCS5Padding2(ciphertext []byte, blockSize int) []byte{
	// 1. 计算最后一个分组缺多少个字节
	padding := blockSize - (len(ciphertext)%blockSize)
	// 2. 创建一个大小为padding的切片, 每个字节的值为padding
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	// 3. 将padText添加到原始数据的后边, 将最后一个分组缺少的字节数补齐
	newText := append(ciphertext, padText...)
	return newText
}

// 删除pks5填充的尾部数据
func PKCS5UnPadding2(origData []byte) []byte{
	// 1. 计算数据的总长度
	length := len(origData)
	// 2. 根据填充的字节值得到填充的次数
	number := int(origData[length-1])
	// 3. 将尾部填充的number个字节去掉
	return origData[:(length-number)]
}

func TripleDESText() {
	// 加密  24字节
	key := []byte("111111111111111111111111")
	result := TripleDESEncrypt([]byte("床前明月光, 疑是地上霜. 举头望明月, 低头思故乡."), key)
	fmt.Println("加密后展示数据：")
	//Bc6UUUEr7C2SoRrK/uSXPiAOcB/IdED7H9JjPowH3NkM3IDityTy94Gpc2J+cHAM0w8udPo2m0BN1gvPQSRE7DErddIWJIY6
	fmt.Println(base64.StdEncoding.EncodeToString(result))
	// 解密
	result = TripleDESDecrypt(result, key)
	fmt.Println("解密之后的数据: ", string(result))
}