package main

import "fmt"

func slice4() {
	// 현재 요소 개수 3 최대 길이 5 인 슬라이스
	slice1 := make([]int, 3, 5)
	slice2 := append(slice1, 4, 5) // 4,5가 추가된 슬라이스
	// len : 요소개수 cap : 최대길이
	fmt.Println("slice1 : ", slice1, len(slice1), cap(slice1))
	fmt.Println("slice2 : ", slice2, len(slice2), cap(slice2))
	slice1[2] = 100 // slice2까지 영향을 받는다
	fmt.Println("slice1 : ", slice1, len(slice1), cap(slice1))
	fmt.Println("slice2 : ", slice2, len(slice2), cap(slice2))
	slice1 = append(slice1, 500)
	fmt.Println("slice1 : ", slice1, len(slice1), cap(slice1))
	fmt.Println("slice2 : ", slice2, len(slice2), cap(slice2))
	slice1 = append(slice1, 600, 700) // 원래 최대 길이 값이었던 5를 넘어서게됨
	// 슬라이스에서 실제 배열을 가리키고 있던 포인터가 다른 주소의 배열을 가리키게됨
	//slice1, slice2는 독립적인 성격으로 바뀜
	slice1[0] = 100 // slice1은 바뀌지만 slice2는 바뀌지않음
	fmt.Println("slice1 : ", slice1, len(slice1), cap(slice1))
	fmt.Println("slice2 : ", slice2, len(slice2), cap(slice2))
}
