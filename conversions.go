package main

import (
	"fmt"
	"strconv"
)

func main() {
	a :=strconv.Itoa(44)
	fmt.Printf("%T \n",a)

	fmt.Println(a)


	var name interface{}=33
	var name1 interface{}="Aman"

	fmt.Println(name1.(string))
	fmt.Println(name.(int))       //assertion
	fmt.Printf("%T \n%T \n",name1,name)

	fmt.Println(name.(int)+12)

}
