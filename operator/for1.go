package main

import "fmt"

func for1(){
	for i := 0 ; i < 10;i++{
		fmt.Print(i,", ")
	}
}

func infiniteloop(){
	for true{}
}