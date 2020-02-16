package main

import "fmt"

// func add1(x int, y int) {
// 	fmt.Println("add function")
// 	fmt.Println(x + y)
// }

// //関数addを呼び出している
// func main() {
// 	add1(10, 20)
// }

// //返り値が１つの場合
// func add2(x int, y int) int {
// 	fmt.Println("add function")
// 	return x + y
// }

// //関数addを呼び出している
// func main() {
// 	r := add2(10, 20)
// 	fmt.Println(r)
// }

//返り値が2つの場合
//返り値の型に対してカッコが必要になる
func add3(x, y int) (int, int) {
	fmt.Println("add function")
	return x + y, x * y
}

//変数名を関数内で宣言して使用することができ、
//returnの後のresultも省略することができる
func cal(price, item int) (result int) {
	result = price * item
	return result
}

//関数addを呼び出している
func main() {
	r1, r2 := add3(10, 20)
	fmt.Println(r1, r2)

	r3 := cal(200, 3)
	fmt.Println(r3)

	//無名関数のようなもの
	f := func() {
		fmt.Println("inner func")
	}
	f()

	//無名関数のようなもの2
	func(num int) {
		fmt.Println(num)
	}(1)
}
