package main

import "fmt"

var g int = 10 // 패키지 전역변수이다. 동일한 패키지 내에서 어디서든 접근이 가능하다.

//g1 := 20 // 패키지 전역변수 선언시에는 선언대입문을 사용하면 안된다. 일반 변수 리터럴을 사용해야한다.

func var5(){
	var m int = 20
	{
		var s = 50
		fmt.Println(m,s,g)
	}
	//m = s + 20 // s는 내부 스코프 안에서 소멸되므로 사용할 수 없다.
}