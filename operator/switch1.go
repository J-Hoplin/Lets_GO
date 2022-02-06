package main

import "fmt"

//go 언어의 switch case에는 break를 붙여주지 않아도 된다.

func switch1(){
	day := 3
	
	switch day{
	case 1:
		fmt.Println("1st day")
	case 2:
		fmt.Println("2nd day")
	case 3:
		fmt.Println("3rd day")
	case 4:
		fmt.Println("4th day")
	default:
		fmt.Println("default")
	}
}

//조건문 비교하기
func switch2(){
	temp := 18
	switch true{
	case temp < 10 || temp > 30:
		fmt.Println("Not good day for outdoor")
	case temp >= 10 && temp < 20:
		fmt.Println("Migth be some cold")
	case temp >= 15 && temp < 25:
		fmt.Println("Good day for outdoor")
	default:
		fmt.Println("Warm")
	}
}