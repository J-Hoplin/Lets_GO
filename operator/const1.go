package main

import "fmt"

func const1(){
	// 상수는 한번 정해지면 다시는 바꿀수 없는 수이다
	// const키워드로 선언을 해주면 된다
	// 상수는 & 연산자를 이용해서 주소값에 접근할 수 없다
	const C int = 10
	b := C * 20
	fmt.Println(b)
}