package main

import "fmt"

// 일반고객
type User struct{
	Name string
	ID string
	Age int
}
// VIP고객
type VIPUser struct{
	UserInfo User
	VIPLevel int
	Price int
}

func struct2(){
	// 구조체를 포함하는 구조체
	// 구조체 필드로 다른 구조체를 포함시킬 수 있다. 두가지의 방법이 있다
	// 1. 일반적인 내장타입처럼 포함하는 방법
	// 2. 포함된 필드 방법


	// < 내장 타입 형태로 포함하는 방식 >
	user := User{"hoplin", "hp",24}
	vip1 := VIPUser{
		// 이렇게 바로 구조체 초기화를 진행하면서 넣을수 있다.
		User{"hoplin2","hp2",24},
		1,
		100000,
	}
	vip2 := VIPUser{
		// 아니면 기존 선언 구조체를 대입할 수 도 있다.
		user,
		1,
		100000,
	}

	fmt.Printf("VIP 유저 : %v ID : %s AGE : %v VIPLevel : %d VIP Price : %d", vip1.UserInfo.Name,vip1.UserInfo.ID,vip1.UserInfo.Age,vip1.VIPLevel,vip1.Price)
	fmt.Println()
	fmt.Printf("VIP 유저 : %v ID : %s AGE : %v VIPLevel : %d VIP Price : %d", vip2.UserInfo.Name,vip2.UserInfo.ID,vip2.UserInfo.Age,vip2.VIPLevel,vip2.Price)
	
}