package main

import "fmt"

func str6(){
	// 문자열 비교하기
	// == , !=를 이용해서 문자열이 같은지 다른지 비교한다.

	str1 := "Hello"
	str2 := "hllo"
	str3 := "Hello"

	fmt.Printf("%s == %s : %v\n",str1,str3,str1 == str3)
	fmt.Printf("%s == %s : %v\n",str1,str2,str1 == str2)
	fmt.Printf("%s == %s : %v\n",str2,str3,str2 == str3)
}