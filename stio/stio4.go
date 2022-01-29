package main

import "fmt"


func stio4(){
	/*
	표준입력

	Scan() : 표준 입력에서 값을 입력받는다.

	Scan() :  값을 채워넣을 변수들의 메모리 주소를 인수로 받는다(&)
	반환값은 성공적으로 입력한 개수와 입력 실패시 에러를 반환한다.

	Scanf() : 표준 입력에서 서식 형태로 값을 입력받는다

	Scanln() : 표준 입력에서 한줄을 읽어서 값을 입력받는다.
	*/

	var a int
	var b int

	n,err := fmt.Scan(&a,&b)
	// nil : 정의되지않은 메모리 주소라는 의미

	/*
	go 에서는 에러를 반환하지 않는 이상, nil값을 에러로 반환한다

	if variable != nil 형태로 예외처리를 해주는것이 좋다.

	따로 try catch와 같은 예외처리 구문은 없다.
	
	*/
	if err != nil{
		fmt.Println(n,err)
	}else{
		fmt.Println(n,a,b)
	}

	// fmt.Scanf() : 서식에 맞춘 입력을 받는다
	// Scanf("%d %d\n",&a,&b)라고 한다면 숫자 하나와 공란 그리고 숫자 하나와 줄바꿈에 맞춰진 입력을 받는다.
	// 서식에 맞춰 입력하기 힘들기 때문에 Scan() 혹은 Scanln()을 사용하는것이 좋다. 
	// 반환값은 Scan()과 동일하다

	var c int
	var d int

	m,err1 := fmt.Scanf("%d %d",&c,&d)
	if err1 != nil{
		fmt.Println(m,err)
	}else{
		fmt.Println(m,c,d)
	}

	// 입력창에 3 4 5 6과 같이 입력하면 a,b에 3 4 가 입력되고 c,d에 5 6이 입력된다.

	// Scanln() : 한줄을 입력받아 인수로 들어온 변수 메모리 주소에 값을 채워준다
	// Scan()과 다른점은 무조건 enter키로 종료해야한다는 점이다.

	var e int
	var f int

	o, err2 := fmt.Scanln(&e,&f)
	if err2 != nil{
		fmt.Println(o,err2)
	}else{
		fmt.Println(n,e,f)
	}
}