package main

import "fmt"


func for2(){
	//continue : 이후 코드 블록을 수행하지 않고 곧바로 후처리를 하고 조건문 검사부터 다시 시작
	//break : for문에서 곧바로 빠져나온다.
	for i := 0; i < 10; i++{
		if i == 3{
			continue
		}
		if i == 6{
			break
		}
		fmt.Println("6 * ",i," = ",6*i)
	}
}