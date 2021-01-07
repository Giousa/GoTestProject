package main

import "fmt"

type Person struct {

	name string
	age int
}

func main() {

	var man = Person{"曹操",45}
	fmt.Println(man)
	var p = &man
	fmt.Println(p)
	*p = Person{"孙权",11}
	fmt.Println(man)

	var p1 *Person = &Person{"刘备",56}
	fmt.Println(p1)


	
}
