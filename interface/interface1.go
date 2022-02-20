package main

import "fmt"

// 인터페이스도 구조체처럼 타입중 하나이므로 type을 써준다.
// 즉 인터페이스도 변수 선언이 가능하고 변수의 값으로 사용할 수 있다는 뜻이다.
type Stringer interface {
	String() string // 메소드 명 , 매개변수
}

type student struct {
	Name string
	Age  int
}

func (s *student) String() string {
	return fmt.Sprintf("My name is %s and my age is %d", s.Name, s.Age)
}

func interface1() {
	student1 := &student{"hoplin", 24}
	var stringer Stringer
	stringer = student1

	fmt.Println(stringer.String())
}
