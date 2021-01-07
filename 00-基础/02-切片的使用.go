package main

import "fmt"

func main() {


	s1 := []int{1,2,3,4,5}
	s2 := make([]int,3)
	s3 := make([]int,10)
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println("-----------")
	copy(s2,s1)
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println("-----------")
	copy(s3,s1)
	fmt.Println(s1)
	fmt.Println(s3)

}
