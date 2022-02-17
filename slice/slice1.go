package main

import "fmt"

func slice1() {
	// 슬라이스 선언은 배열과 비슷하지만 [] 안에 배열 개수 적지않는다

	sli := []int{1, 2, 3}
	//배열과 동일하게 인덱싱,값 변경도 된다

	for i := 0; i < len(sli); i++ {
		sli[i] += 10
	}

	// range : 인덱스, 해당 인덱스의 값 을 순서대로 반환한다.
	for i, v := range sli {
		sli[i] = v * 2
	}

	for _, v := range sli {
		fmt.Println(v)
	}
}
