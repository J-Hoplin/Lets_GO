package main

// 함수 선언부에 반환타입을 적을때 변수명까지 지정해주면
//return 문으로 해당 변수를 명시적으로 반환하지 않아도 값을 반환할 수 있다.

func func3(a int,b int)(result int, success bool){
	if b == 0{
		result = 0
		success = false
		return
	}
	result = a / b
	success = true
	return
}