package main

import "fmt"

// < 포인터를 사용한 경우 >
type test struct{
	value int
	data [200]int
}

func changeData(arg *test){
	arg.value = 999
	arg.data[100] = 200
}

func ptr6(){
	var a test
	changeData(&a)
	fmt.Printf("value : %d\n", a.value)
	fmt.Printf("value data[100] : %d\n",a.data[100])
}