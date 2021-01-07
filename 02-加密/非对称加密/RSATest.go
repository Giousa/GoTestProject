package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"os"
)

func main() {

	//生成公钥私钥
	//RsaGenKey(512)

	//加密 倘若需要加密的字符串过长，会报错  rsa: message too long for RSA public key size
	//每次加密的“位”数，不能超过密钥的长度值减去11。例如：密钥长度为512（2的9次方）64字节，那么最大名为长度为53字节即（（53+11）*8）
	res := RSAEncrypt([]byte("今天的天气非常不错，在学习规划的程度上，你还有待加强"),[]byte("file/encryption/public.pem"))
	fmt.Println("加密后字节数组：",res)
	fmt.Println("加密后字符串：",base64.StdEncoding.EncodeToString(res))
	//解密
	resEnd := RSADecrypt(res,[]byte("file/encryption/private.pem"))
	fmt.Println("解密后字节数组：",resEnd)
	fmt.Println("解密后字符串：",string(resEnd))


}

//生成公钥和私钥
// 参数bits: 指定生成的秘钥的长度, 单位: bit
func RsaGenKey(bits int) error{
	// 1. 生成私钥文件
	// GenerateKey函数使用随机数据生成器random生成一对具有指定字位数的RSA密钥
	// 参数1: Reader是一个全局、共享的密码用强随机数生成器
	// 参数2: 秘钥的位数 - bit
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil{
		return err
	}
	// 2. MarshalPKCS1PrivateKey将rsa私钥序列化为ASN.1 PKCS#1 DER编码
	derStream := x509.MarshalPKCS1PrivateKey(privateKey)
	// 3. Block代表PEM编码的结构, 对其进行设置
	block := pem.Block{
		Type: "RSA PRIVATE KEY",//"RSA PRIVATE KEY",
		Bytes: derStream,
	}
	// 4. 创建文件
	privFile, err := os.Create("file/encryption/private.pem")
	if err != nil{
		return err
	}
	// 5. 使用pem编码, 并将数据写入文件中
	err = pem.Encode(privFile, &block)
	if err != nil{
		return err
	}
	// 6. 最后的时候关闭文件
	defer privFile.Close()

	// 7. 生成公钥文件
	publicKey := privateKey.PublicKey
	derPkix, err := x509.MarshalPKIXPublicKey(&publicKey)
	if err != nil{
		return err
	}
	block = pem.Block{
		Type: "RSA PUBLIC KEY",//"PUBLIC KEY",
		Bytes: derPkix,
	}
	//../file/encryption/
	pubFile, err := os.Create("file/encryption/public.pem")
	if err != nil{
		return err
	}
	// 8. 编码公钥, 写入文件
	err = pem.Encode(pubFile, &block)
	if err != nil{
		panic(err)
		return err
	}
	defer pubFile.Close()

	return nil

}

//RSA公钥加密
func RSAEncrypt(src, filename []byte) []byte {
	// 1. 根据文件名将文件内容从文件中读出
	fmt.Println("加密文件地址：",string(filename))
	file, err := os.Open(string(filename))
	if err != nil {
		return nil
	}
	// 2. 读文件
	info, _ := file.Stat()
	allText := make([]byte, info.Size())
	file.Read(allText)
	// 3. 关闭文件
	file.Close()

	// 4. 从数据中查找到下一个PEM格式的块
	block, _ := pem.Decode(allText)
	if block == nil {
		return nil
	}
	// 5. 解析一个DER编码的公钥
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil
	}
	pubKey := pubInterface.(*rsa.PublicKey)

	// 6. 公钥加密
	result, errRes := rsa.EncryptPKCS1v15(rand.Reader, pubKey, src)
	if errRes != nil{
		fmt.Println("加密失败：errRes = ",errRes)
	}
	fmt.Println("result = ",result)
	return result
}


//RSA私钥解密
func RSADecrypt(src, filename []byte) []byte {
	// 1. 根据文件名将文件内容从文件中读出
	fmt.Println("解密文件地址：",string(filename))
	file, err := os.Open(string(filename))
	if err != nil {
		return nil
	}
	// 2. 读文件
	info, _ := file.Stat()
	allText := make([]byte, info.Size())
	file.Read(allText)
	// 3. 关闭文件
	file.Close()
	// 4. 从数据中查找到下一个PEM格式的块
	block, _ := pem.Decode(allText)
	// 5. 解析一个pem格式的私钥
	privateKey , err := x509.ParsePKCS1PrivateKey(block.Bytes)
	// 6. 私钥解密
	result, _ := rsa.DecryptPKCS1v15(rand.Reader, privateKey, src)

	return result
}