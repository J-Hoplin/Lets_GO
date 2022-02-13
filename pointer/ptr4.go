package main

import "fmt"

func ptr4(){
	// 포인터 변수를 초기화하지 않을경우 기본값은 ni이다. 
	var a *int
	fmt.Printf("%v",a) // <nil>
}