package main

import "fmt"

func method7_1(i []int) []int {
	idx := 2
	for j := idx + 1; j < len(i); j++ {
		i[j-1] = i[j]
	}
	i = i[:len(i)-1] // 슬라이싱을 통해 마지막 값을 잘라준다.
	return i
}
func method7_2(i []int) []int {
	idx := 2
	return append(i[:idx], i[idx+1:]...) // ... 슬라이스 / 배열 내 모든 요소 의미
}

func slice7() {
	// 슬라이스 중간 요소를 삭제하는 방법을 살펴좌
	// 하나만 삭제하는 경우
	// 1. 중간 요소 삭제
	// 2. 중간 요소 삭제 후 뒤의 값을 앞당기기
	// method1()을 통해 기본적인 방법을 보자
	// method2()를 통해 append로 코드를 개선해 보자

	slice := []int{1, 2, 3, 4, 5, 6}
	slice = method7_1(slice)
	fmt.Println(slice)
	// 이를 append를 통해 코드를 개선시켜보자
	slice = method7_2(slice)
	fmt.Println(slice)
}
