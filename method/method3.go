package main

import "fmt"

type myInt int

func (a myInt) add(b int) int {
	return int(a) + b
}

func method3() {
	var a myInt = 10
	fmt.Println(a.add(20))
}
