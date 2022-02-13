package main

import "fmt"

// 여러 필드를 묶어 하나의 구조체를 만든다. 구조체는 배열과 달리 다른 타입의 값들을 변수 하나로 묶어준다.
// 구조체 타입 이름이 대문자로 시작하면 외부에 공개하는 타입이라는 의미가 된다.
type Student struct{
	Name string
	Class int
	No int
	Score float64
}


func struct1(){
	// 구조체 타입 선언 : 일반 원시타입과 다르지 않다

	var student Student
	// 구조체 각 요소를 접근하기 위해서는 '.'을 이용하면 된다
	// 위와 같이 초기값을 생략하면 기본값으로 초기화된다 .str = "" int 0 float64 0.0

	student.Name = "hoplin"
	student.Class = 1
	student.No = 1
	student.Score = 100

	fmt.Println("Name ", student.Name)
	fmt.Println("Class ", student.Class)
	fmt.Println("No ", student.No)
	fmt.Println("Score ", student.Score)

	// 초기화 과정을 생략할수도 있지만 반대로 초기화를 해줄수도 있다. 초기화를 위해서는 중괄호 사이에 넣어 초기화한다.

	// 한줄초기화
	var student2 Student = Student{"test",2,2,100}
	// 여러줄 초기화
	var student3 Student = Student{
		"test2",
		3,
		3,
		100, // 여러줄 초기화할때는 마지막에 쉼표를 붙여야한다.
	}
	fmt.Println(student2.Name)
	fmt.Println(student3.Name)

	// 일부만 초기화하기 : 일부만 초기화할때는 필드명 : 값 형식으로 초기화한다.

	var student4 Student = Student{Name: "test4", Score: 90}
	fmt.Println(student4.Name, student4.Score)
}