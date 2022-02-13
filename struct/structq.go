package main

import (
	"fmt"
	"unsafe"
)

//q1

type Product struct{
	Name string
	Price int
	ReviewScore float64
}
//q3 : 다음 구조체(Padding) 패딩을 최대한 줄이고 구조체 크기를 적으세요

type Padding struct{
	A int8 //1
	B int // 8
	C float64 // 8
	D uint16 //2
	E int //8
	F float32 // 4
	G int8 // 1
}


// q3 Answer
// 8바이트를 앞에 두고 4바이트 하나 2바이트 하나 1바이트 하나이므로 하나의 8바이트로 합쳐 패딩이 없게끔 만든다.
// 예상 총 32바이트

type ImprovedPadding struct{
	B int
	C float64
	E int
	F float32
	D uint16
	A int8
	G int8
}

func structq(){
	a := ImprovedPadding{
		1,1.0,2,3,3.0,4,5,
	}
	fmt.Println(unsafe.Sizeof(a))
}