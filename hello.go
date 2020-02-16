package main

import "fmt"

func by2 (num int) string {
	if num % 2 == 0 {
		return "ok"
	} else {
		return "n0"
	}
}

func main() {
	num := 5
	if num%2 == 0 {
		fmt.Println("by 2")
	} else {
		fmt.Println("else")
	}

	//以下のように１つの式にまとめることができる
	if result2 := by2(50); result2 == "ok"{
		fmt.Println("great2")
	}
}
