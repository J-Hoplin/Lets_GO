package main

import "fmt"

func ptr1(){
	// 포인터 : 메모리 주소를 값으로 갖는 타입

	// 메모리 주솟값을 변숫값으로 가질 수 있는 변수를 포인터 변수라고 한다.
	// 포인터 변수는 자료형 앞에 *를 붙여 선언한다. 구조체든 원시타입이든, 참조타입이든 모두 똑같다.
	
	var p *int
	// 포인터는 메모리 주소값을 갖는다. 메모리 주소값을 알아오기 위해서는 데이터 앞에 &를 붙이면 된다.
	var a int
	p  = &a // 포인터 변수 p에 int형 변수 a의 주소값을 넣는다.
	// p는 a의 메모리 주소를 접근할 수 있으므로 아래와 같이 값을 변경할 수 있다.
	*p = 20 // 포인터 변수 앞에 *를 붙여 해당 주소에 대해 접근할 수 있다. * : 역참조 연산자(C언어 기준)

	// %p : 주소값을 나타내는 formatter
	fmt.Printf("Value of P : %p \n", p)
	fmt.Printf("Value that P point : %v\n",*p)
	*p = 100
	fmt.Printf("Value that P point : %v\n",*p)
}