package main

import "fmt"

func main() {

	//長さは3、キャパシティは5である、このままだと000が入る
	n := make([]int, 3, 5)

	//長さはlen、キャパシティはcap
	fmt.Printf("len=%d cap=%d value=%v", len(n), cap(n), n)

	c = make([]int, 0, 5)
	for i := 0; i < 5; i++ {
		c = append(c, i)
		fmt.Println(c)
	}
	fmt.Println(c)
}
