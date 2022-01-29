package main

import "fmt"


func stio3(){
	/*
	
	특수문자

	\n : 개행문자

	\\ : \문자 자체를 출력한다

	\t : 탭을 삽입한다

	\" : "를 출력한다. 큰 따옴표로 묶인 문자열 내부에 따옴표를 넣을때 사용한다.
	

	문자열 서식 지정자

	%s : 문자열을 출력한다. 기본적으로 특수문자(이스케이프 시퀸스)를 모두 적용한다

	%q : 문자열을 출력한다. 이스케이프 시퀸스를 모두 무시하고 문자 그대로 출력한다.
	*/

	str := "Hello\tGO\tWorld\n\"GO\"is Awesome!\n"

	fmt.Print(str)
	fmt.Printf("%s", str)
	fmt.Printf("%q", str)
}