package main

import "fmt"

func str5(){
	str1 := "Hello"
	str2 := "World"

	// 문자열 연산은 +와 += 를 통해 문자열을 이을 수 있다.

	str3 := str1 + str2
	fmt.Println(str3)
	str4 := str1 + " " + str2
	fmt.Println(str4)
}