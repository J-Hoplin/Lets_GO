package main

import "fmt"

func method8_1(sli []int, val int) []int {
	idx := 2
	sli = append(sli, 0)

	for i := len(sli) - 2; i >= idx; i-- {
		sli[i+1] = sli[i]
	}
	sli[idx] = val
	return sli
}

func method8_2(sli []int, val int) []int {
	idx := 2
	sli = append(sli[:idx], append([]int{val}, sli[idx:]...)...)
	return sli
}

func method8_3(sli []int, val int) []int {
	idx := 2
	sli = append(sli, 0) // 0추가
	copy(sli[idx+1:], sli[idx:])
	fmt.Println(sli)
	sli[idx] = val
	return sli
}

func slice8() {
	// 요소 추가하기
	// 슬라이스 중간에 요소를 추가하려면 다음 과정을 따라야 한다
	// 슬라이스 맨 뒤 요소를 하나 추가한다
	// 맨 뒤 값부터 삽입하려는 위치까지 한칸씩 뒤로 밀어준다
	// 삽입하는 위치의 값을 바꿔준다

	// method8_1 : 반복문 순회를 통한 방법
	// method8_2 : append()를 사용한 개선
	// method8_3 : method8_2에서 발생하는 불필요한 메모리(중첩 append()에서 생기는 임시 슬라이스) 를 copy로 개선하기

	slice := []int{1, 2, 3, 4, 5}
	slice = method8_1(slice, 100)
	fmt.Println(slice)
	slice = method8_2(slice, 200)
	fmt.Println(slice)
	slice = method8_3(slice, 300)
	fmt.Println(slice)
}
