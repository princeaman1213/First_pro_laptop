package main

import (
	"fmt"
	"sync"
	"time"
)
//var wg sync.WaitGroup
var done chan bool
func main() {

	c :=gen()

	/*
		//Fan out
	c1:=make([]chan int,10)
	for i:=0;i<10;i++{
		c1[i]=fact(c)
	}

	//Fan in
	c11:=mearge(c1...)
	 */
	// Can also use the above approach for fan out and fan in instead of the below one to avoid using unnecessary memory due to many variables

	//Fan out
	c1 :=fact(c)
	c2 :=fact(c)
	c3 :=fact(c)
	c4 :=fact(c)
	c5 :=fact(c)
	c6 :=fact(c)
	c7 :=fact(c)
	c8 :=fact(c)
	c9 :=fact(c)
	c10:=fact(c)

	//Fan in
	c11:=mearge(c1,c2,c3,c4,c5,c6,c7,c8,c9,c10)


    var co int
	for r:=range c11{
		co++
		fmt.Println(co,"  :",r)
	}

}

func gen()chan int{
	c:=make(chan int)
	go func(){
		for ii:=0;ii<10;ii++{
			for i:=0;i<10;i++{
				c<-i+1
				time.Sleep(time.Microsecond)   //to get output in order
			}

			//time.Sleep(time.Second)
		}
		close(c)
	}()
	//fmt.Println("dfssadf",<-c)
	//fmt.Println("dfssadf",<-c)
	return c
}

func fact(n chan int) chan int{
	c2 :=make(chan int)

	go func(){
		for r:=range n{
			//f=f*(r+1)
			c2<-fact1(r)
		//	fmt.Println("Teting",fact1(r))
			time.Sleep(time.Millisecond*5)
			//fmt.Println("rrrrrr",r)
		}
		//c2<-f
		close(c2)
	}()
	//fmt.Println("fact of", n ," is :",<-c)
	return c2
	//wg.Done()
}

func fact1(n1 int) int{
	var f int=1
	for k:=n1;k>0;k--{
		f*=k
	}
	return f
}

func mearge(c5 ...chan int) chan int{
	c6:=make(chan int)
	var wg sync.WaitGroup
	//fmt.Println("lent is ",len(c5))
	wg.Add(len(c5))
	for _, r := range c5 {
		go func(r chan int) {

				for r8 := range r {
					c6 <- r8
					time.Sleep(time.Millisecond)
				}
				wg.Done()

			}(r)
		}

	go func() {
		wg.Wait()
		//time.Sleep(time.Second)
		close(c6)
	}()

	return c6
}