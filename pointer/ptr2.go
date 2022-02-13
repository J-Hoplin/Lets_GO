package main

import "fmt"

func ptr2(){
	var a int = 500
	var p *int

	p = &a
	fmt.Printf("Value of P : %p\ns",p)
	fmt.Printf("Value that p points : %d\n", *p)

	*p = 100
	fmt.Printf("Value that p points : %d\n", *p)
}