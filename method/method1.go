package main

import "fmt"

type account1 struct {
	balance int
}

func (a *account1) withdrawPointer(amount int) {
	a.balance -= amount
}

func withdrawFunc(a *account1, amount int) {
	a.balance -= amount
}

func method1() {
	a := &account1{100} // account1타입의 인스턴스 생성
	withdrawFunc(a, 30)
	a.withdrawPointer(10)
	fmt.Println(a.balance)
}
