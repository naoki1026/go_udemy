package main

import "fmt"

//可変長引数
//関数の定義時に取り決めた引数を増やすことができないため、...intを使用する
func foo(params ...int) {

	//lenは長さ
	fmt.Println(len(params), params)
	for _, param := range params {
		fmt.Println(param)
	}
}

func main() {
	foo()
	foo(10, 20)
	foo(10, 20, 30)

	s := []int{1, 2, 3}
	fmt.Println(s)

	foo(s...)
}
