package main

import "fmt"

// interface{}는 메소드를 가지고 있지 않은 빈 인터페이스이다.
// 가지고있어야할 메소드가 하나도 없기때문에, 모든 타입이 빈 인터페이스로 쓰일 수 있다.
// 빈 인터페이스는 어떤 값이든 받을 수 있는 함수, 메소드 변숫값을만들때 사용한다.

func PrintVal(v interface{}) {
	switch t := v.(type) {
	case int:
		fmt.Printf("v is int %d\n", t)
	case float64:
		fmt.Printf("v is float %f\n", t)
	case string:
		fmt.Printf("v is string %s\n", t)
	default:
		fmt.Printf("Not supported type : %v\n", t)
	}
}

type Student struct {
	Age int
}

func interface4() {
	PrintVal(10)
	PrintVal(3.14)
	PrintVal("Hello")
	PrintVal(Student{10})
}
