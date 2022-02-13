package main

import "strings"

func ToUpperCase(str string) string{
	var result string
	for _,c := range str{
		if c >= 'a' && c <= 'z'{
			result += string('A' + (c - 'a'))
		}else{
			result += string(c)
		}
	}
	return result
}

// string 합연산을 빈번하게 하면 메모리가 낭비된다. 그렇기 때문에 string합 연산을 빈번히 사용하는 경우 strings패키지의 Builder를 이용해 메모리 낭비를 줄일 수 있다.
// string은 합연산을 할때마다 새로운 공간을 할당해 문자열을 더한다 -> 메모리 공간이 버려진다. + 성능문제까지

// strings패키지의 Builder를 사용하면 메모리를 매번 생성하지 않고 기존 메모리 공간에 더하게 된다.
// GO Document : https://go.dev/src/strings/builder.go
func ToUpperCase2(str string) string{
	var builder strings.Builder

	for _,c := range str{
		if c >= 'a' && c <= 'z'{
			builder.WriteRune('A' + (c - 'a'))
		}else{
			builder.WriteRune('A' + (c - 'a'))
		}
	}
	return builder.String() // Return builder to String also in documentation
}

func str7(){
	// 소문자를 대문자로 만들기 (toUpperCase)
}