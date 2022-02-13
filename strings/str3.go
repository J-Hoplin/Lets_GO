package main

import "fmt"

func str3(){
	// string타입, rune슬라이스 타입인 []rune타입은 상호 타입 변환이 가능하다.
	// string타입, []rune타입 모두 문자들의 집합을 나타내기 때문에 상호 타입 변환이 가능하다.
	// 이렇게 변환할 경우 rune배열 각 요소에는 문자열의 각 글자가 대입된다.

	s1 := "Hello 월드"
	runes := []rune(s1)
	
	// len은 아래와 같이 string타입 변수 길이는 문자열 바이트 길이가 반환된ㄷ.
	// 반면 배열을 len의 인수로 넣으면 요소 개수가 반환된다

	// 오버라이딩 : 상위클래스의 메소드를 하위 클래스에서 재정의
	// 내부적으로 len함수가 오버로딩(동일한 이름의 메소드이지만 다른 매개변수 타입, 개수를 지님)이 되어있지 않을까 라는 추측
	fmt.Printf("len(s1) : %d\n",len(s1)) // 12
	fmt.Printf("len(runes) : %d\n",len(runes)) //8	
}