package main

import "fmt"

func stio1(){
	a := 10
	var b int = 20
	var f float64 = 32799438743.8297

	/*
	
	fmt.Print() : 함수 매개변수로 들어온 데이터를 출력한다
	fmt.Println() : Print()함수에 개행문자를 추가한 형태
	fmt.Printf() : 포맷에 맞게 출력을 할 수 있게 한다.
	
	*/

	fmt.Print("a : ", a , "b : ",b)
	fmt.Println("a : ", a , "b : ",b)
	fmt.Printf("a : %d b : %d f : %f",a,b,f)

	/*
	형식지정자는 모두 몰라도 적어도 많이 쓰는 아래 세가지만 알아도 된다

	%d : 정수형
	%f : 실수형
	%s : 문자열

	만약 어떤 서식 문자를 써야할지 모르는 경우에는 서식문자 %v를 사용하면 된다

	%v : 데이터 타입에 맞춰서 기본 형태로 출력한다.
	
	*/
	fmt.Printf("a : %v b : %v f : %f",a,b,f)
}