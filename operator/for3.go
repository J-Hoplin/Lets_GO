package main

import (
	"bufio"
	"fmt"
	"os"
)


func for3(){
	stdin := bufio.NewReader(os.Stdin)
	for{
		fmt.Print("입력하세요 >> ")
		var number int
		_,err := fmt.Scanln(&number)
		if err != nil{
			fmt.Println("숫자를 입력하세요")
			stdin.ReadString('\n')
			continue
		}
		fmt.Printf("입력하신 숫자는 %d입니다", number)
		if number % 2 == 0{
			break
		}
	}
	fmt.Println("For문이 종료되었습니다")
}