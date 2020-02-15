package main

import "fmt"

func main() {
	var a [2]int
	a[0] = 100
	a[1] = 200
	fmt.Println(a)

	//配列は最初に定義したサイズを変更することができない
	var b = [2]int{100, 200}
	fmt.Println(b)

	//スライスを使用することで追加することができる
	var c = []int{100, 200}
	c = append(c, 300)
	fmt.Println(c)
}
