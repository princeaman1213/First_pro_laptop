package main

import (
	"fmt"
	//"time"
	"time"
)

func main() {
	c :=make(chan int)
	done :=make(chan bool)
	n:=2

		go func() {
			for i := 0; i < 5; i++ {
				c <- i
				time.Sleep(time.Millisecond)
			}
			close(c)
		}()


	for i:=0;i<n;i++{
		go func() {
			for n:=range c{
				fmt.Println(n)
			}
			done<-true
		}()
	}

	for i:=0;i<n;i++{
		<-done
	}

}


/*
package main

import (
	"fmt"
	//"time"
	"time"
)

func main() {
	c :=make(chan int)
	//done :=make(chan bool)
	//n:=2

		go func() {
			for i := 0; i < 5; i++ {
				c <- i
				time.Sleep(time.Millisecond)
			}
			close(c)
		}()

arr :=make([]int,1)
	//for i:=0;i<n;i++{
		//go func() {
			for n:=range c{
				//fmt.Println(n)
				arr=append(arr, n)
			}
			//done<-true
		//}()
	//}
fmt.Println(arr)
	//for i:=0;i<n;i++{
	//	<-done
	//}

}

 */