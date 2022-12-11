package main

import "fmt"

func rangeAddition () {
	a := 10
	b:= 1
	var sum int
	var add int

	if a == b {
		fmt.Println(a)

	} else if a > b {
		 for ;b <= a ;b++ {
			sum += b
			fmt.Println(sum) 
    } 
	}else {
	 for ;a <= b;a++{
	 add += b
	 fmt.Println(add)
		 }

	}
	}