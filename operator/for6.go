package main

import "fmt"

//중첩 내부 로직을 함수로 빼어 플래그변수, 레이블 사용 최소화
func find45(a int)(int,bool){
	for b:=1;b<=9;b++{
		if a * b == 9{
			return b,true
		}
	}
	return 0,false
}


func for6(){
	a :=1
	b := 0
	
	for ; a <=9;a++{
		if c,found := find45(a);found{
			b = c
			break
		}
	}
	fmt.Printf("%d * %d = %d",a,b,a * b)
}