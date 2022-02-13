package main

import "fmt"

func arr4(){
	var t [5]int = [5]int{1,2,3,4,5}
	for i,v := range t{
		fmt.Println(i, v)
	}
	/*
	range : 배열 요소를 순회한다

	첫번째 반환값은 인덱스 ,두번째 반환값은 요소값을 대입한다.
	만약 인덱스가 필요없고 요소값만 필요한 경우 _ 를 이용해 무효화할 수 있다,
	*/

	for _, v := range t{
		fmt.Println(v)
	}
}