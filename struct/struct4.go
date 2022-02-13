package main

import "fmt"

type User2 struct{
	Name string
	ID string
	Age int
	Level int // 겹치는 필드
}

type VIP2 struct{
	User2
	Price int
	Level int // 겹치는 필드
}

func struct4(){
	// 필드가 중복되는 경우
	/*
	앞에서 했던 방법을 통해서는 구조체 필드가 동일한 경우 충돌 할 가능성이있다.
	그런 경우에는 각 구조체 내 구조체 타입.요소 형태로 접근해야한다.
	*/
	a := User2{"Hoplin","hp",24,1}
	vip := VIP2{
		a,
		10000,
		10,
	}

	fmt.Printf("VIP 유저 : %v ID : %s AGE : %v Normal Level : %d VIPLevel : %d VIP Price : %d",
	vip.Name,
	vip.ID,
	vip.Age,
	vip.User2.Level, // 이와 같이 충돌하는 필드 Level에 대해서 구조체 타입.필드 형태로 접근을 한다.
	vip.Level,
	vip.Price)

}