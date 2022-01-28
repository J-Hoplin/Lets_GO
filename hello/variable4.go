package main

import "fmt"

func var4(){
	// Go언어는 강타입중에서도 타입검사를 강하게 하는 최강타입 언어이다.
	a := 3 // int
	var b float64 = 3.5 // float64

	// var c int = b; // float64변수를 int에 대입 불가
	// d := a * b // int, float64 : 다른 타입의 변수간 연산 불가

	var e int64 = 7
	// f := a * e // int, int64동일한 숫자형 타입이지만, 타입이 다르기 때문에 연산 불가.
	// 같은 숫자값이라도 타입이 다르면 연산이 안된다. 이럴때는 형변환을 해주어야한다.
	// 이러한 것을 타입 반환 이라고 한다. 타입명을 적고 ()안에 변수를 적어주면된다.

	d := a * int(b) // 실수형을 int형으로 바꾸면 정수부분만 남음
	
	f := int64(a) * e
	fmt.Println(a,d,f)

}