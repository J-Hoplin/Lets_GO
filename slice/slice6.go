package main

import "fmt"

//sli1을 복제한 새로운 sli2를 반환하는 함수들

// 일반적인 반복문 순회 이용하기
func method1(sli1 []int) []int {
	sli2 := make([]int, len(sli1))
	for i, v := range sli1 {
		sli2[i] = v
	}
	return sli2
}

//append()를 이용해 코드 개선
func method2(sli []int) []int {
	//append()함수를 이용해 slice1의 모든 값을 복제한 새로운 슬라이스를 만들어서 반환한다
	//배열이나 슬라이스 뒤에 ... 를 붙이면 모든 요솟값을 넣어준것과 동일하게된다.
	// append
	return append([]int{}, sli...)
}

func method3() {
	// copy()를 사용해서 개선하기
	// func copy(dst, src []Type) int
	// 첫번째 인자로 복사한 결과가 저장되는 슬라이스가,두번째 인수로 복사 대상이 되는 슬라이스 변수를 넣는다.
	// 반환값은 실제로 복사된 요소의 개수이다.
	// 목적지 슬라이스, 도착지 슬라이스중 길이가 작은 개수만큼 복사된다.
	sli := []int{1, 2, 3, 4, 5}
	sli2 := make([]int, 3, 5) // len 3, cap 5인 슬라이스
	sli3 := make([]int, 10)   //len,cap 모두 10인 슬라이스

	copy(sli2, sli)
	fmt.Println("sli2 : ", sli2, len(sli2), cap(sli2))
	copy(sli3, sli)
	fmt.Println("sli3 : ", sli3, len(sli3), cap(sli3))
}

func method4() {
	// method2, method3을 혼합해 아래와 같이 사용할 수 도 있다
	sli := []int{1, 2, 3, 4, 5}
	sli2 := make([]int, len(sli))
	copy(sli2, sli)
	// 사실 이 방법은 method2의 append([]int{}, sli1...)과 동일한 방법이다.
	fmt.Println(sli2)
}

func slice6() {
	// 두 슬라이스가 같은 배열을 가리키면서 발생하는 문제점은 크다

	// 1. 하나의 값이 바뀌면 다른 하나의 값이 바뀐다
	// 2. append()연산에 의해 중간에 cap값보다 더 큰 배열이 필요한 경우 Data의 배열 포인터가 다른 포인터로 변경된다.
	//그렇기 때문에 이런 문제가 발생하지 않도록 배열을 복제해 주는 작업이 필요할 수 도 있다.
	sli1 := []int{1, 2, 3, 4, 5}
	sli2 := method1(sli1)
	// sli1의 값이 바뀌었지만 sli2의 값은 바뀌지않음
	sli1[1] = 100
	fmt.Println(sli1)
	fmt.Println(sli2)

	sli4 := method2(sli1)
	sli1[2] = 200
	fmt.Println(sli1)
	fmt.Println(sli4)
	method3()
}
