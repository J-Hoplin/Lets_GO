package main

import "fmt"

func slice5() {
	/*
		슬라이싱
		슬라이싱은 배열 일부를 집어내는 기능을 말한다
		파이썬과 매우 유사
	*/
	// 기본적인 형태 : array[startIndex : endIndex]
	// 범위 : startIndex ~ endIndex - 1인덱스 까지 불러옴

	// 슬라이싱으로 배열 일부 가리키는 슬라이스 만들기
	array := [5]int{1, 2, 3, 4, 5}
	slice := array[1:3]
	// 슬라이싱의 값 반환은 슬라이스가 반환된다.
	// 슬라이스 최대 길이를 나타내는 cap같은 경우에는 총 길이에서 시작인덱스를 뺀 값으로 결정된다
	// 5 - 1 -> 4인것이다.
	// 슬라이스는 결론적으로 배열의 2,3번째 요소를 가리키고 있다 그렇기 때문에 array의 2,3번째 값이 바뀌면 slice도 같이 바뀐다
	fmt.Println(array)
	fmt.Println(slice)
	array[1] = 100
	array[2] = 200
	fmt.Println(array)
	fmt.Println(slice)
	// 배열 뿐만 아니라 슬라이스도 슬라이싱 할 수 있다.
	// cap의 값 결정과 방법은 똑같다
	// 그리고 슬라이싱해 반환받은 슬라이스는 기존 슬라이스의 해당 범위 요소들을 가리키고 있다는 사실도 동일하다
	sli1 := []int{1, 2, 3, 4, 5}
	sli2 := sli1[1:3]
	sli2 = sli1[0:3]
	sli2 = sli1[:3]  // sli1[0:3]과 동일한 의미
	sli3 := sli1[3:] // sli1의 인덱스 3부터 끝까지 의미
	sli4 := sli1[:]  // 전체 슬라이싱
	fmt.Println(sli2)
	fmt.Println(sli3)
	fmt.Println(sli4)
}
