package main

import "fmt"

type Vertex struct {
	//小文字にしてしまうと外部からアクセスできなくなってしまう
	// X int
	// Y int
	//並べて記述することもできる
	X, Y int
	Z string
}

//構造体
func main(){
	v := Vertex{X:1, Y:2}
	fmt.Println(v)
	fmt.Println(v.X, v.Y)

	v1 := Vertex{1, 2, "test"}
	fmt.Println(v1)
}