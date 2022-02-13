package main

import "fmt"

func arr2(){
	// 배열 -> 같은 타입 데이터로 이루이진 타입이다.
	// 선언 -> var (varname) [elementcount](type)
	var t [5]float64 = [5]float64{24.0,25.9,27.8,26.8,26.2}
	for i:= 0;i < 5 ; i++{
		fmt.Println(t[i])
	}
} 