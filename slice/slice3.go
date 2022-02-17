package main

import "fmt"

func changeArray(arr [5]int) {
	arr[2] = 200
}

func changeSlice(sli []int) {
	sli[2] = 200
}

func slice3() {
	array := [5]int{1, 2, 3, 4, 5}
	slice := []int{1, 2, 3, 4, 5}
	changeArray(array)
	changeSlice(slice)
	fmt.Println(array)
	fmt.Println(slice)
}
