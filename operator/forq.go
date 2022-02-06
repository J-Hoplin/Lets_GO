package main

import "fmt"

// 1~10 역순출력
func forq(){
	for i := 10; i >0;i--{
		fmt.Printf("%d\t",i)
	}
}

// 2~9단 구구단
func forq2(){
	for i:= 2; i < 10; i++{
		if i == 3 || i == 6{
			continue
		}
		for j:=1;j<10;j++{
			fmt.Println(i," * ",j," = ", i*j)
		}
		fmt.Println()
	}
}


// 9까지의 홀수들의 제곱값
func forq3(){
	for i := 1; i < 10;i++{
		if (i % 2) != 0{
			fmt.Println(i," * ",i," = ",i * i)
		}
	}
}

//이중 for문을 사용해서 역삼각형 별 출력
func forq4(){
	for i := 5;i>0;i--{
		for j := i;j > 0;j--{
			fmt.Print("*")
		}
		fmt.Println()
	}
}