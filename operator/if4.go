package main

import "fmt"

// 중첩함수

func splitPay()bool{
	return true
}

func if4(){
	price := 3600000
	
	if price > 5000000{
		if splitPay(){
			fmt.Println("Split pay")
		}else{
			fmt.Println("Pay once")
		}
	}else if price >= 3000000 && price <= 5000000{
		fmt.Println("Pay Once")
	}
}

func getMyAge()(int,bool){
	return 33, true
}
// if 초기문;조건문
//초기문 자리에 하나의 구문이 올 수 있으므로 ; 로 구분해준다.
// 초기문은 검사에 사용할 변수를 초기화할때 주로 사용한다.
func if5(){
	if age,ok := getMyAge(); ok && age < 20{
		fmt.Println("You are young", age)
	}else if ok && age < 30{
		fmt.Println("Nice age",age)
	}else if ok{
		fmt.Println("Good age",age)
	}else{
		fmt.Println("Error")
	}
	// 초기문에 선언한 변수의 범위로는 if문 안으로 한정된다.
	// if문 블록 외부에서는 사용할 수 없다
	//fmt.Println("You're age is ",age) // Error
}