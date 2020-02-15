package main

import "fmt"

func main() {
	n := []int{1, 2, 3, 4, 5}
	fmt.Println(n)
	fmt.Println(n[1])

	//カンマの数を○番目を指定
	fmt.Println(n[2:4])
	fmt.Println(n[2:])
	fmt.Println(n[:4])
	fmt.Println(n[:])

	//複数定義することができる
	var board = [][]int{
		[]int{1, 2, 3},
		[]int{4, 5, 6},
		[]int{7, 8, 9},
	}
	fmt.Println(board)
}
