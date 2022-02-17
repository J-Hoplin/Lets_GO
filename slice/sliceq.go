package main

import (
	"fmt"
	"sort"
)

type Player struct {
	name  string
	age   int
	score int
	pass  float64
}

type Players []Player

func (s Players) Len() int {
	return len(s)
}

func (s Players) Less(i, j int) bool {
	return s[i].score < s[j].score
}
func (s Players) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func sliceqq() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// 다음 슬라이스가 있을때 처음부터 마지막 두번째 전까지 잘라내는 구문을 쓰세요

	slice = slice[:len(slice)-1-1]
	fmt.Println(slice)

	a := []Player{
		{"a", 20, 30, 10.5},
		{"b", 19, 20, 30.5},
		{"c", 27, 25, 50.1},
	}
	// 선수 데이터를 표현하는 구조체를 만든 뒤 높은 득점부터 낮은 순으로 정렬하세요
	// 기본 정렬은 오름차순으로 정렬된다 하지만 높은 득점부터 낮은 순으로 정렬하는것이므로 내림차순이다.
	sort.Sort(sort.Reverse(Players(a)))
	fmt.Println(a)
}
