package main

import "fmt"

type OurDB struct {
	Name string
}

// Question2 : OurDB구조체의 모든 공개된 메소드를 이용하는 인터페이스 구현하기
type ADB interface {
	GetData() string
	WriteData() (data string)
	Close() error
}

// Question3 : 어떤 경우에도 Runtime에러가 발생하지 않도록 CheckAndRun()함수 구현하기
func CheckAndRun(stringer Stringer) interface{} {
	if r, ok := stringer.(Reader); ok {
		r.Read()
		return true
	} else {
		return fmt.Sprint("Can't convert to Reader interface type")
	}
}
