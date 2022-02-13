package main

import (
	"fmt"
	"unsafe"
)

type test struct{
	Age int // 4바이트
	Score float64 // 8바이트
}


type test2 struct{
	a int8
	b int
	c int8
	d int
	e int8
}

type test3 struct{
	a int
	b int
	c int8
	d int8
	e int8
}


func struct5(){
	user := test{10,99.9}
	fmt.Println(unsafe.Sizeof(user)) // unsafe.Sizeof() : 변수의 메모리 공간 크기를 반환한다. 16바이트

	// 구조체는 총 12바이트여야할것같지만 왜 16바이트일까? 메모리 정렬때문이다.
	/*
	메모리 정렬이란? 컴퓨터가 데이터에 효과적으로 접근하고자 메모리를 일정 크기 간격으로 정렬하는것을 의미한다.
	데이터가 레지스터 크기와 똑같은 크기로 정렬되어있으면 더욱 효과적으로 데이터를 읽어올 수 있다.
	32비트 컴퓨터 -> 레지스터가 4바이트
	64비트 컴퓨터 -> 레지스터가 8바이트

	예를 들어 64비트 컴퓨터에서 int64데이터 시작주소가 100번일 경우 100은 8의 배수가 아니기 때문에 레지스터 크기 8에 맞춰져있지 않다.
	이 경우 메모리를 읽어올 때 성능을 손해보기 때문에 프로그램 언어에서 데이터를 만들때 8의 배수인 메모리 주소에 데이터를 할당한다
	예를 들면 이 경우에는 100이 아닌 104에 할당된다.

	type VIPUser1 struct{
		VIPLevel int32
		Score float64
	}

	var user VIPUser1
	
	위와 같은 예시가 있다고 하자. user의 시작 주소가 240번지 인 경우에 VIPLevel값 또한 시작 주소가 240이 된다. 단 VIPLeveldms 4바이트 공간을 
	차지하기 때문에 바로 붙여서 Score를 할당하면 244번지가 되고 이는 8의 배수가 아니므로 성능적 손해를 본다.
	그렇기 때문에 프로그램 언어에서는 VIPLevel과 Score사이에 4바이트를 띄워 할당함으로서 8의 배수에 맞춘다. 이와같이 공간을 띄우는 것을 메모리 패딩이라고 부른다. 
	

	*/
	//< 메모리 패딩을 고려한 필드 배치 >
	// test2구조체 사이즈 측정

	b := test2{1,2,3,4,5}
	fmt.Println(unsafe.Sizeof(b)) // 40
	// 40이 되는 이유는 간단하다. int8, 즉 1바이트 짜리 변수는 7바이트가 패딩된다. 그렇기 때문에 총 3개의 int8타입이 패딩이 되어 40바이트가 나오는 것이다. 이를 test3과 같이 개선할 수 있다.
	c := test3{1,2,3,4,5}
	fmt.Println(unsafe.Sizeof(c))//24
	// 24가 되는 이유는 처음 2개의 int가 각각 8바이트로 정렬된 후 이후 세 int8타입에 대해 3바이트가 채워진다. 그 후 8의 배수를 맞추기 위해 5바이트만 패딩이 되어 총 24바이트가 되는것이다.

}