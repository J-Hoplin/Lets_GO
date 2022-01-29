package main

import "fmt"

func stio6(){
	var a int
	var b int
	// 1 2 3 4 5 6 7 8 9 10
	// 위와같이 입력하면 Scan()과 달리 연속적으로 값을 받지 못한다. Scanln()의 기본 정의와 같이 각 입력별로 enter값을 입력해 주어야한다.
	for i := 1; i <= 5; i++{
		n,err := fmt.Scanln(&a, &b)
		if err != nil{
			fmt.Println(n,err)
		}else{
			fmt.Println(a,b)
		}
	}
	fmt.Println(a,b)
}