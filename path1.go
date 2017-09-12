package main

import (
	"fmt"
	"math"
	"sync"
)

var check sync.WaitGroup

func main() {
	check.Add(3)

		arrx:=[]float64{0,185,212,265,273,279,292,367,392,426,434,437,477,491,499}
		arry:=[]float64{0,78,99,146,188,56,243,69,321,324,115,132,364,225,121}    //coordinates of person 1 on route 1

		arrx1:=[]float64{0,123,156,190,223,235,278,302,306,387,404,458,480,499}
		arry1:=[]float64{0,35,63,355,255,144,35,366,465,234,134,234,465,324,211}    //coordinates of person 1 on route 2

		arrx2:=[]float64{0,142,156,178,234,278,333,387,390,397,406,453,487,499}
		arry2:=[]float64{0,24,57,167,388,299,120,268,444,211,100,98,167,22,110}    //coordinates of person 1 on route 3

		go calc(1,arrx,arry)
	    go calc(2,arrx1,arry1)
	    go calc(3,arrx2,arry2)

	check.Wait()
}

func calc(n int,arrx,arry []float64){

	var i2 int
	var d2 float64
	var v2 float64
	for i2=1;i2<len(arrx);i2++{
		d2 += math.Sqrt(math.Pow((arrx[i2-1] - arrx[i2]), 2) + math.Pow((arry[i2-1] - arry[i2]), 2))
	}

	d2+=arry[i2-1]
	v2=d2/float64(i2*5)
	fmt.Printf("\ndistance of route %v is          : %.2f metres \naverage velocity on this route is : %.2f m/s",n,d2,v2)
	check.Done()
}




