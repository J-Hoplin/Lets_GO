package main

import "fmt"

func slice2() {
	//배열과 달리 슬라이스는 요소를 추가해 길이를 늘릴 수 있다(파이썬 리스트 느낌)
	// append()를 이용해서 요소를 추가해 줄 수 있다.
	// 첫번째 인수로는 추가하고자 하는 슬라이스를 적고 그 뒤에 요소를 적어주면 해당 요소가 맨 뒤에 들어간 슬라이스를 결과로 반환한다.

	var slice = []int{1, 2, 3}
	slice2 := append(slice, 4)
	fmt.Println(slice2)

	// append()로 하나 뿐만 아니라 여러개의 요소를 한번에 추가해 줄 수 있다.
	// 첫번째 매개변수 이후 모든 값이 슬라이스에 들어가는 값들이 된다

	slice = append(slice, 3, 4, 5, 6, 7)
	fmt.Println(slice)
	for i := slice[len(slice)-1] + 1; i <= 15; i++ {
		slice = append(slice, i)
	}
	fmt.Println(slice)
}
