package main

import (
	"fmt"
)

func main() {
	db :=map[string]map[int]string{
		"a" : map[int]string{
			1 : "abate",
			2 : "alarm",
			3 : "algorithm",
			4 : "analogy",
			5 : "admire",
			6 : "alarity",
			7 : "abscond",
			8 : "apathy",
		},
		"b" : map[int]string{
			1 : "bad",
			2 : "ball",
			3 : "base",
		},
	}
	var l string
	var index int
	fmt.Println("Enter the first letter :")
	fmt.Scanf("%v",&l)
	fmt.Println("Suggestions are listed as follows...")
	for i,v:=range db[l]{
		fmt.Println(i," : ",v)
	}

	fmt.Println("Enter the index of the word to be printed :")
	fmt.Scanf("%v",&index)

	if index<=len(db[l]){
		fmt.Println("Selected word is...\n",db[l][index])
	}else{
		fmt.Printf("No words found for provided index. Enter the correct index ")
	}

}
