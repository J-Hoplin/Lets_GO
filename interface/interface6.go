package main

import "fmt"

type Reader1 interface {
	Read()
}

type Closer interface {
	Close()
}

type File struct {
}

func (f *File) Read() {

}

func ReadFile(reader Reader1) interface{} {
	// 타입 변환 성공 여부를 반환할 수 있다. 두번째 매개변수에 반환을 성공했는지, 실패했는지를 받을 수 있다.
	a, ok := reader.(Closer)
	if ok {
		a.Close()
		return true // boolean반환
	} else {
		return fmt.Sprintf("Unable to Convert Interface type") // string반환
	}
}

func interface6() {
	file := &File{}
	ReadFile(file) // 문법상 틀린것은 아니다 하지만 file은 File구조체 타입이고, File구조체 타입은 Reader1인터페이스의 Read()만 구현했으므로, Closer인터페이스 타입이 될 수 없다.
}
