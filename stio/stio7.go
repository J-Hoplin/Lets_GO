package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// string입력받기

func stio7(){
	kbReader := bufio.NewReader(os.Stdin)
	fmt.Print("이름을 입력하세요 : ")
	strName, err := kbReader.ReadString('\n')
	if err != nil{
		log.Fatal(err)
	}else{
		strName = strings.TrimSpace(strName)
		fmt.Printf("당신의 이름 : %s\n",strName)
	}
}