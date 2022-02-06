package main

/*
함수 이름은 Multiple입니다

입력으로 int타입 2개를 받고 반환을 1개 합니다
두 입력값을 반환하는 함수작성
*/

func Multiple(a int, b int)int{
	return a * b
}

//피보나치 수열 재귀적 구현

func Fibonacci(a int)int{
	if a < 2{
		return a
	}
	return Fibonacci(a - 1) + Fibonacci(a - 2)
}