package main

import (
	"fmt"
	"strconv"
)

func main() {
	var x int = 1
	xx := float64(x)

	//型、値、小数
	fmt.Printf("%T %v %f\n", xx, xx, xx)

	//型、値、整数
	var y = 1.2
	yy := int(y)
	fmt.Printf("%T %v %d\n", yy, yy, yy)

	//文字型を数値型に変換
	var s string = "14"

	//アンスコにすることで、エラーが表示されなくなる
	i, _ := strconv.Atoi(s)
	fmt.Printf("%T %v\n", i, i)
}
