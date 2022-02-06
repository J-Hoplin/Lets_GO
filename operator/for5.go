package main

import "fmt"

// 레이블을 이용하는 방법
// break할때 앞서 정의한 레이블을 적어주면 그 레이블에서 가장 먼저 속한 for문까지 모두 종료한다.
// 레이블은 혼동을 일으킬 수 있으므로, 꼭 필요한 경우가 아니면 사용하지 말자
// 클린코드를 위해서는 중첩된 내부 로직을 함수로 묶어 중첩을 줄이고 플래그 변수나 레이블 사용을 최소화해야한다.

func for5(){
	a := 1
	b := 2

OuterFor:
	for ; a <= 9 ;a++{
		for b = 1;b <=9;b++{
			if a * b == 45{
				break OuterFor
			}
		}
	}
	fmt.Printf("%d * %d = %d\n",a,b,a * b)
}
