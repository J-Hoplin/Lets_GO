package main

// Reader 인터페이스
type Reader interface {
	Read() (n int, err error) // error : Go기본 인터페이스
	Close() error
}

// Writer인터페이스
type Writer interface {
	Write() (n int, err error)
	Close() error
}

// 인터페이스 선언시 메소드 이름이 겹치면 안된다
// 이 경우 Close() error가 같은 메소드 형식이므로 합쳐져서 하나의 Close() error메소드만 ReadWriter인터페이스에 포함된다.
type ReadWriter interface {
	Reader
	Writer
}

func interface3() {

}
