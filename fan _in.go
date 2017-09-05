package main

import "fmt"

func main() {
	c :=fanin(str("ice"),str("fire"))  //c:=one(c,c)

	for i:=0;i<13;i++{
		fmt.Println(<-c)
	}

}

func str(s string) chan string{
	c :=make(chan string)
	go func(){
		for i:=0;i<8;i++{
			c<-fmt.Sprintln(s,"is good",i)
		}
		close(c)
	}()
	return c
}

func fanin(c1,c2 chan string) chan string{
	c3 :=make(chan string)
	go func(){
		for {
			c3<-<-c1
		}

	}()

	go func(){
		for {
			c3<-<-c2
		}
	}()

	return c3

}

/*
package main

import "fmt"

func main() {
	c :=fanin(str("ice"),str("fire"))  //c:=one(c,c)

	for r:=range c{
		fmt.Println(r)
	}

}

func str(s string) chan string{
	c :=make(chan string)
	go func(){
		for i:=0;i<8;i++{
			c<-fmt.Sprintln(s,"is good",i)
		}
		close(c)
	}()
	return c
}

func fanin(c1,c2 chan string) chan string{
	c3 :=make(chan string)
	sem :=make(chan bool)
	go func(){
		for r:=range c1{
			c3<-r
		}
		sem<-true
	}()

	go func(){
		for r1:=range c2{
			c3<-r1
		}
		sem<-true
	}()

	go func(){
		<-sem
		<-sem
		close(c3)
	}()

	return c3

}
 */