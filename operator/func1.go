package main

import "fmt"

/*
함수

함수의 첫글자가 대문자인 경우에는 패키지 외부로 공개되는함수임을 의미한다.

함수 선언 형태는 아래와 같다

func (name)(parameters)(returntype){

}

*/

func Add(a int, b int)int{
	return a + b
}

func PrintAvgScore(name string,math int,eng int, history int){
	total := math + eng + history
	avg :=  total / 3
	fmt.Println(name,"님의 평균점수는 ",avg,"점입니다")
}

func func1(){
	c := Add(3,10)
	fmt.Println(c)
	PrintAvgScore("hoplin",90,90,100)
}