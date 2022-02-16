package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

/*
숫자게임

- 랜덤 값을 하나 정한다
	- 랜덤함수를 받기 위해서는 math/rand 패키지와 현재시각을 지원하는 time패키지를 사용해야한다
	- maht/rand를 임포트 한다. 특정 범위에서 int타입 랜덤값을 생성하는 rand.Intn(range)를 사용한다. rand.Intn(100)이라고 하면 0~99사잇값이 생성됩니다.
		- 랜덤값은 유사 랜덤 알고리즘으로 값을 만든다 그렇기 때문에 초기값이 같으면 항상 동일한 값을 반환한다. 이 초기값을 랜덤 시드라고 하는데 랜덤이 산출되려면 매번 랜덤 시드를 다른 값으로 설정해야한다. 랜덤 시드는 rand.Seed()함수를 이용해 설정할 수 있다.


- 사용자 입력을 받습니다.
- 사용자가 입력한 숫자와 랜덤숫자가 동일할때까지 반복합니다
*/

var stdin = bufio.NewReader(os.Stdin)

func InputIntValue() (int, error) {
	var n int
	_, err := fmt.Scanln(&n)
	if err != nil {
		stdin.ReadString('\n')
	}
	return n, err
}

func main() {
	// time.Now() : Time객체를 반환한다.
	// func (t Time) UnixNano() int64 : 1970년 1월 1일부터 Time객체가 나타내는 시각까지 경과한 시간을 나노초 단위로 나타내 반환한다.
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(100)
	fmt.Println("랜덤 숫자 : ", r)
	cnt := 1
	for {
		fmt.Print("숫자값을 입력하세요 : ")
		n, err := InputIntValue()
		if err != nil {
			fmt.Println("숫자만 입력하세요")
		} else {
			if n > r {
				fmt.Println("입력하신 숫자가 더 크다")
			} else if n < r {
				fmt.Println("입력하신 숫자가 더 작습니다")
			} else {
				fmt.Println("숫자를 맞췄습니다. 시도한 횟수 : ", cnt)
				break
			}
			cnt++
		}
	}
}
