package main

//iota로 간편하게 열거값 사용하기
//iota는 0부터 시작해 1씩 증가한다. iota는 소괄호를 벗어나면 다시 초기화된다.
// 상수목록을 const와 ()로 묶고 iota를 사용하면 0부터 1씩 증가하며 초기화된다

const (
	Red int = iota//0
	Blue int = iota//1
	Green int = iota//2
)
// 첫번째 값과 똑같은 규칙이 적용되면 타입과 iota를 생략할 수 있다.

const (
	C1 int = iota + 1 // 0 + 1 = 1
	C2 //2 1 + 1 = 2
	C3 //3 2 + 1 = 3
)