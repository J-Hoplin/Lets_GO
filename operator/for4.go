package main

import "fmt"

//중첩for문
func for4(){
	for i := 0 ; i < 3; i++{
		for j:=0 ; j < 5; j++{
			fmt.Print("*")
		}
		fmt.Println()
	}
}