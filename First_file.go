package main

import(
	"fmt"
	"io/ioutil"
    "errors"
)


func main() {
	err :=errors.New("dfdsf fds")
	a,err := ioutil.ReadFile("name.txt")     //This program needs to be understood and learned
	if err !=nil{
		panic("panic panic")
	}

        fmt.Println("err",err)
		fmt.Println(string(a))

}
