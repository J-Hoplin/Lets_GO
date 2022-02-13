package main

import "fmt"

func str4(){
	// 문자열 순회
	// 문자열에 들어있는 글자들을 순회하자. 방법에는 세가지가 있다

	// 1. 인덱스를 사용한 바이트 단위 순회(결국 string도 문자(rune)가 모여있는 배열과 마찬가지이다)
	// 2. []rune으로 타입 변환 후 한 글자씩 순회
	// 3. range키워드를 이용한 한글자씩 순회

	s1 := "Hello 월드!"

	// 인덱스를 사용한 바이트 단위 순회
	for i := 0; i < len(s1);i++{
		fmt.Printf("Type : %T Value : %d Character Value : %c\n",s1[i],s1[i],s1[i])
	}
	// 위 결과를 보면 영문과 공백문자, 특수기호 ! 는 잘 출력되지만 한글은 출력되는것을 볼 수 있다.
	// 인덱스를 통해 접근하면 요소의 타입은 uint8, 즉 byte가 된다. 그렇기 때문에 3바이트인 한글을 담지 못하는 것이다
	// 한글 순회를 위해서는 []rune(rune은 int32타입과 비슷한 타입, 즉 4바이트이다.)타입으로 변환 후 range키워드를 이용해 순회한다
	// 또한 string에 대해 len을 하면 string의 크기를 반환하므로 총 문자열 길이인 9가 아닌 13만큼 순회를 한다.
	fmt.Println()


	arr := []rune(s1)
	for i := 0;i < len(arr);i++{
		fmt.Printf("타입 : %T Value : %d Character Value : %c\n",arr[i],arr[i],arr[i])
	}
	// 이렇게 []rune으로 변환한 다음 순회하면 한 글자씩 순회할 수 있지만 []rune으로 변환하는 과정에서 별도의 배열 할당이 이루어진다
	// 이는 불필요한 메모리 사용을 야기한다.
	// range키워드를 이용해 한글자씩 순회해보자  range는 기본적으로 인덱스 ,값 총 두개의 값을 반환한다.

	fmt.Println() // 인덱스 값은 사용하지 않을것이므로 _를 통해 무시
	for _, v := range s1{
		fmt.Printf("타입 : %T Value : %d Character Value : %c\n",v,v,v,)
	}
	// 이와 같이 range를 통해 순회하면 불필요한 메모리 낭비를 줄일 수 있다.
}