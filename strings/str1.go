package main

import "fmt"

func str1(){
	// 문자열 : 말 그대로 문자열이다. 자료형 이름은 string이다. "" 혹은 ``로 묶어서 표시한다

	// "" 로 묶으면 특수문자가 작동한다
	s1 := "Hello\tworld\n"
	//``로 묶으면 특수문자가 작동하지 않는다.
	s2 := `Go language string\ttype\n`
	fmt.Println(s2)
	fmt.Printf("%s",s1) // 문자열의 formatter는 s1
	
	// 문자 하나를 표현하는데 rune타입을 사용한다(사실상 char타입이다.)
	// UTF-8은 GO의 기본 문자코드이다.
	//UTF-8은 한글자가 1~3바이트 크기이기 때문에 UTF-8문자값을 가지기 위해서 3바이트가 필요하다 하지만 GO언어기본타입에서 3바이트는 제공하지 않는다
	// 그렇기 때문에 rune타입은 정수타입은 int32타입의 별칭 타입이다

	var char rune = '가' // rune타입을 넣기 위해서는 ""가 아닌 ''로 묶어야한다.

	fmt.Printf("%T\n",char) //int32
	fmt.Println(char)// 44032 char값을 출력한다  char타입이 int32이므로 숫자로 출력된다
	fmt.Printf("%c\n",char) // %c포맷을 이용해 문자 하나를 출력한다

	
}