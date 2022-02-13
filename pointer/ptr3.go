package main

import "fmt"

func ptr3(){
	// == 연산을 이용해서 포인터가 같은 메모리 공간을 가리키는지 확인할 수 있다.

	var a int = 10
	var b int = 20

	var p1 *int = &a
	var p2 *int = &a
	var p3 *int = &b

	fmt.Printf("p1 == p2 : %v\n", p1 == p2) // true
	fmt.Printf("p2 == p3 : %v\n",p2 == p3) // false
}