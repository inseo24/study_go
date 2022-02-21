package main

// main package -> main func
// compile이 필요하지 않다면 main이 필요하지 않음 -> main은 컴파일을 위해 필요한 파일

import (
	"fmt"
	"strings"

	"github.com/inseo24/learngo/something"
)

func multiply(a int, b int) int {
	return a * b
}

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func repeatMe(words ...string) {
	fmt.Println(words)
}

func main() {
	fmt.Println("한글 출력 되나???")

	// upper-case로 작성하면 다른 패키지에서 import 가능
	something.SayBye()

	// 상수
	const name string = "seoin"
	fmt.Println(name)

	// 변수
	var naming string = "Camel-case"
	naming = "Pascal"
	fmt.Println(naming)

	// 다른 변수 선언 방식 - type은 go에서 알아서 찾아줌
	// 단, 이 축약형 방식은 func 안에서만 작성 가능
	man := "wook"
	man = "hyoung"
	fmt.Println(man)

	// go의 type
	fmt.Println(multiply(2, 3))

	// func는 리턴을 2가지 타입 이상을 받을 수 있음
	totalLength, upperName := lenAndUpper("seoin")
	fmt.Println(totalLength, upperName)

	// 아래처럼 일부 return을 무시할 수도 있음
	totalLen, _ := lenAndUpper("seoin")
	fmt.Println(totalLen)

	// 여러 매개변수를 받는 방법
	repeatMe("seo", "in", "wook")

}
