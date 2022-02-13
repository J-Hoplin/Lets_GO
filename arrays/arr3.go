package main

import "fmt"


func arr3(){
	var num [3]int // 기본적으로 모 0으로 초기화
	var strs [3]string // nil값으로 초기화(출력하면 아무것도 안나옴)
	y := [...] int{10,20,30} // ...으로 배열 개수를 생략할 수 있다. 이 경우 초기화되는 배열의 요소 개수와 동일해준다.
	//배열 크기값은 상수여야한다.
	const P int = 10
	z := [P]int{1,2,3,4,5,6,7,8,9,10}
	for i:=0 ; i<len(num);i++{
		fmt.Println(num[i])
		fmt.Println(strs[i])
		fmt.Println(y[i])
	}
	for j := 0 ; j < len(z);j++{
		fmt.Println(z[j])
	}
	
	
}