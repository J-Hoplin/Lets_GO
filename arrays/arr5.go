package main

import "fmt"

func arr5(){
	//다중배열(다차원 배열)

	a := [2][5]int{
		{1,2,3,4,5},
		{6,7,8,9,10},
	}

	for _, arr := range a{
		for _, v := range arr{
			fmt.Print(v," ")
		}
		fmt.Println()
	}
}