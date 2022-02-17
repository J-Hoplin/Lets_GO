package main

import (
	"fmt"
	"sort"
)

type Student struct {
	Name string
	Age  int
}

type Students []Student //[]Student의 별칭타입 Students

func (s Students) Len() int {
	return len(s)
}

func (s Students) Less(i, j int) bool {
	return s[i].Age < s[j].Age
}

func (s Students) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func slice9() {
	// 슬라이스 정렬하기
	// go언어의 sort패키지를 사용해 슬라이스를 정렬할 수 있다.
	s := []int{5, 2, 6, 3, 1, 4}
	sort.Ints(s)
	fmt.Println(s)

	// 구조체 슬라이스 정렬
	// Sort()함수 정렬을 위해서는 Len(), Less(), Swap()세 메소드를 사용해야한다.
	//https://pkg.go.dev/sort#Interface
	// 이 메소드만 구현하면 정의된 구조체도 정렬을 할 수 있다.
	a := []Student{
		{"A", 3},
		{"b", 4},
		{"c", 1},
		{"d", 2},
		{"e", 5},
	}
	sort.Sort(Students(a))
	fmt.Println(a)
}
