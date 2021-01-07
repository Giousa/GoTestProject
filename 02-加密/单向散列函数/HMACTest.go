package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
)

/**
这个方法是用于认证的。
场景：A 向 B发送了一条重要消息，生成散列函数，例如MD5加密后，将加密后消息生成认证码发过去。B 通过发来的加密消息，配合认证码和秘钥，判断这条消息的真实性
主要是判断消息的真实性，而不是用来对消息进行加密和解密

消息认证码是对消息进行认证并确认其完整性的技术。通过使用发送者和接收者之间共享的密钥，就可以识别
出是否存在伪装和篡改行为。
消息认证码可以使用单向散列函数HMAC， 对称加密也可以实现， 这里不再进行介绍。
消息认证码中，由于发送者和接收者共享相同的密钥，因此会产生无法对第三方证明以及无法防止否认等问题。
解决这些问题：数字签名。

一般，
 */
func main() {

	key := []byte("我是消息认证码秘钥")
	src := []byte("我是消息认证码测试数据")

	//生成认证码
	result := GenerateHMAC(src, key)

	//验证
	final := VerifyHMAC(result, src, key)
	if final {
		fmt.Println("消息认证码认证成功!!!")
	} else {
		fmt.Println("消息认证码认证失败 ......")
	}
}

// 生成消息认证码
func GenerateHMAC(src, key []byte) []byte {
	// 1. 创建一个底层采用sha256算法的 hash.Hash 接口
	myHmac := hmac.New(sha256.New, key)
	// 2. 添加测试数据
	myHmac.Write(src)
	// 3. 计算结果
	result := myHmac.Sum(nil)

	fmt.Println("生成消息认证码字节数组：",result)
	fmt.Println("生成消息认证码字符串：",base64.StdEncoding.EncodeToString(result))

	return result
}

//验证消息认证码
func VerifyHMAC(res, src, key []byte) bool {

	// 1. 创建一个底层采用sha256算法的 hash.Hash 接口
	myHmac := hmac.New(sha256.New, key)
	// 2. 添加测试数据
	myHmac.Write(src)
	// 3. 计算结果
	result := myHmac.Sum(nil)
	// 4. 比较结果
	return hmac.Equal(res, result)
}

