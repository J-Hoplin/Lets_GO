package main

import "fmt"

type data struct{
	value int
	data [200]int
}

func ChangeData(arg data){
	arg.value = 999
	arg.data[100] = 200
}

func ptr5(){
	//왜 포인터를 사용하는가?
	// 변수 대입 혹은 함수 인수 전달은 항상 값을 복사하기 떄문에 많은 공간을 사용하는 문제와 큰 메모리 공간을 복사할 때 발생하는 성능문제도 있다
	// 또한 다른 공간으로 복사되기때문에(Call by Reference Call by Value) 변경사항이 적용되지 않는다.
	
	// 포인터를 사용하지 않는경우
	
	var a data
	ChangeData(a)
	// 참조타입 -> Call by Reference를 통해 값을 변경해야한다.
	fmt.Printf("value : %d\n",a.value) // 아무런 값의 변화가 없다
	fmt.Printf("data[100] = %d\n",a.data[100]) 

}