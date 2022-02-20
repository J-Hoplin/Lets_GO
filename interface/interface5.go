package main

import "fmt"

type Stringer1 interface {
	String() string
}

type Student1 struct {
	Age int
}

func (s *Student1) String() string {
	return fmt.Sprintf("Studnet's Age : %d\n", s.Age)
}

func PrintAge(stringer Stringer1) {
	s := stringer.(*Student1)
	fmt.Printf("Age : %d\n", s.Age)
}

func interface5() {
	s := &Student1{15}
	PrintAge(s)
}
