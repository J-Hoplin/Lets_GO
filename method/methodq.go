package main

import (
	"fmt"
	"time"
)

// Quesiton2 : 주어진 함수를 메소드로 변경하시오

type ParkingLot struct {
	LotSize int
}

func (p *ParkingLot) ParkCar(carSize int) {
	p.LotSize -= carSize
}
func (p *ParkingLot) printLotSize() {
	fmt.Println(p.LotSize)
}

// Question4 문제 구현하기
type Courier struct {
	Name string
}

type Product struct {
	Name  string
	Price int
	ID    int
}

type Parcel struct {
	Pdt           *Product
	ShippedTime   time.Time
	DeliveredTime time.Time
}

func (c *Courier) SendProduct(p *Product) *Parcel {
	// 메소드가 호출되면 Parcel객체를 생성한다. ShippedTime을 현재시간으로 설정한다
	// Parcel의 product변수는 메소드 인수로 들어온 값을 사용한다
	new_product := &Parcel{ShippedTime: time.Now(), Pdt: p}
	// 생성된 Parcel반환
	return new_product
}

func (p *Parcel) Delivered() *Product {
	// 메소드 호출시 리시버의 Delivered Time을 현재 시각으로 변경
	// 리시버 pdt를 함수 결과로 반환한다. pdt : Product 포인터
	p.DeliveredTime = time.Now()
	return p.Pdt
}

func methodq() {
	//Question2
	lot := &ParkingLot{100}
	lot.ParkCar(10)
	lot.printLotSize()
}
