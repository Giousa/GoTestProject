package main

import (
	"fmt"
	"unicode"
)

func main() {

	str := "LoginInfo"
	fmt.Println(Lcfirst(str))

}


/**
	首字母大写
 */
func Ucfirst(str string) string {
	for i, v := range str {
		return string(unicode.ToUpper(v)) + str[i+1:]
	}
	return ""
}

/**
	首字母小写
*/
func Lcfirst(str string) string {
	for i, v := range str {
		return string(unicode.ToLower(v)) + str[i+1:]
	}
	return ""
}