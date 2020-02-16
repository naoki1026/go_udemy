package main

import "fmt"

// func main() {
// 	x := 0

// 	//変数にfunctionを定義することができる
// 	increment := func() int {
// 		x++
// 		return x
// 	}
// 	fmt.Println(increment())
// 	fmt.Println(increment())
// 	fmt.Println(increment())
// }

//返り値にfunctionを返すことができる
func incrementGenerator() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

func circleArea(pi float64) func(radius float64) float64 {
	return func(radius float64) float64 {
		return pi * radius * radius
	}
}

func main() {

	//変数にfunctionを入れることができる
	//int型を返すことができる
	counter := incrementGenerator()
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())

	c1 := circleArea(3.14)
	fmt.Println(c1(3))

	c2 := circleArea(3)
	fmt.Println(c2(3))
}
