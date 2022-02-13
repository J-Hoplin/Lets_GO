package main

import "fmt"

type Actor struct{
	Name string
	HP int
	Speed float64
}

// 주어진 인수를 이용해 값을 초기화해 반환하는 NewActor구성하시오
func NewActor(name string,hp int,speed float64) *Actor{
	return &Actor{name,hp,speed}
}

func ptrq(){
	var actor *Actor = NewActor("Gold",99,100)
	fmt.Println(actor.Speed)
	fmt.Println(actor.Name)
}