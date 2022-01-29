package main

import "fmt"

// fmt.Scan() 다중입력 테스트
func stio5(){
	var g int
	var h int

	// 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20
	// 위와 같이 입력하면 한번에 입력이 된다. 그리고 입력이 끝나고 총 10번 출력이 이루어지는것을 볼 수 있다.
	for i := 0 ; i < 10; i++{
		fmt.Scan(&g,&h)
		fmt.Println(g,h)	
	}

}