package main

import (
	"bufio"
	"fmt"
	"os"
)

func stio8(){
	/*
	사용자 표준 입력장치로 입력하면 입력 데이터는 컴퓨터 내부에 표준 입력 스트림이라는 메모리 공간에 임시 저장된다.

	Scan()함수들은 그 표준 입력 스트림에서 값을 읽은 후 입력값을 처리한다.

	가장 먼저 입력한 데이터부터 읽어오므로 데이터가 거꾸로 저장된다.
	먼저 입력된 데이터가 먼저 읽히는 데이터 구조 FIFO구조이다.
	Scan()함수는 먼저 표준 입력 스트림에서 한 글자씩 읽어온다.

	var a,b int
	fmt.Scan(&a,&b) 
	만약 위와 과 같이 되어고 입력을 Hello 10으로 되어있으면 한글자씩 읽어오므로, H,e를 가져오게 된다.
	여러번 Scan()함수를 호출해야하는 경우 이러한 문제에 부딫힐 수 있다.(입력 버퍼가 초기화 되지 않는 문제)

	이를 방지하기위해서는 표준 입력 스트림을 지워주어야한다.
	*/

	// NewReader()함수는 인수로 입력되는 입력스트림을 가지고 Reader객체를 생성해준다.
	// 표준 입력 스트림을 의미하는 os.Stdin을 이용해 Reader객체를 만든다.
	stdin := bufio.NewReader(os.Stdin)
	var a,b int
	n,err := fmt.Scanln(&a,&b)
	if err != nil{
		fmt.Println(err)
		// \n이 나올때까지 입력스트림을 읽는다. 읽는 이유는 에러가 난 경우에, 버퍼를 \n이 나올때까지 읽어 버퍼를 비우기 위함이다.
		stdin.ReadString('\n')
	}else{
		fmt.Println(n,a,b)
	}
	n,err = fmt.Scanln(&a,&b)
	if err != nil{
		fmt.Println(err)
	}else{
		fmt.Println(n,a,b)
	}
}