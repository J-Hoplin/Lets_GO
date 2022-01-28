package main

import "fmt"

/*
타입별 기본값

모든 정수타입(int8,int16,int32,uint8,uint16,uint32,int,uint, byte,rune....etc)
-> 0
모든 실수타입(float32,float64,complex64)
-> 0.0
Boolean
-> False
string
-> 빈 문자열

위 타입 외
-> nil(정의되지 않은 메모리 주소를 나타내는 go 키워드이다.)
*/

func var3(){
	var a int = 3 // 기본 형태의 변수선언이다.
	var b int // 초기화 과정을 생략해줄 수 있다.
	var c = 4 // 타입을 생략할 수 있다. 변수타입은 우변의 값의 타입에 따라 결정된다. 정수형은 기본 자료형이 int, 실수형은 기본 자료형이 float64이다.
	d := 5 // 선언대입문이다. 선언과 대입을 한꺼번에 하는 구문이다. 선언 대입문을 사용하면 var키워드와 타입지정 과정을 생략해 선언할 수 있다.
	fmt.Println(a,b,c,d)
}