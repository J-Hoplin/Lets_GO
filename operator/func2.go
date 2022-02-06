package main

// 함수는 여러개의 반환값을 반환할 수 있다
// 반환타입 명시시 소괄호로 묶어 표현해준다.

func Divide(a int,b int)(int,bool){
	if b == 0{
		return 0,false
	}else{
		return a / b, true
	}
}
