package main

import "fmt"

type User1 struct{
	Name string
	ID string
	Age int
}

type VIPUser1 struct{
	User1 // 필드명을 생략하고 구조체 타입만 선언
	VIPLevel int
	Price int
}

func struct3(){
	// 포함된 필드 방식
	// 위 예시에서 구조체 내 다른 구조체 접근을 위해서 구조체.구조체.값 이렇게 접근을 하였다. 
	// 구조체 내 다른 구조체 선언할 때 필드명을 생략하면 .을 한번만 찍어서 접근할 수 있다.

	user := User1{"hoplin","hp1",24}
	vip := VIPUser1{
		user,
		1,
		10000,
	}
	// 아래와 같이 vip 안에 다른 구조체에 대해 바로 요소처럼 접근하는것을 볼 수 있다.
	fmt.Printf("VIP 유저 : %v ID : %s AGE : %v VIPLevel : %d VIP Price : %d", vip.Name,vip.ID,vip.Age,vip.VIPLevel,vip.Price)

}