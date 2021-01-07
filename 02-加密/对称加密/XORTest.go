package main

import (
	"fmt"
)

/**
按位异或
一般是针对整型，使用^符号，string等类型无法使用
 */
func main() {

	var x,y string

	//明文
	x = "hello world"
	//密钥
	y = "111ppp555"

	//加密后：
	m := []byte(x)
	n := []byte(y)
	fmt.Println("m = ",m)
	fmt.Println("n = ",n)
	//i := m^n
	//i := 12^11233333333333

	 b := 12
	 c := 100
	 d := b^c
	 fmt.Println("解密int类型：",d)
	 fmt.Println("加密int类型：",d^c)


	//解密后：


}
