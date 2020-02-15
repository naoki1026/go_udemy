package main

import "fmt"

// func main() {
// 	var i int = 1
// 	var f64 float64 = 1.2
// 	var s string = "text"
// 	var t bool = true
// 	var f bool = false
// 	fmt.Println(i, f64, s, t, f)
// }

//カッコでひとまとめにすることができる
// func main() {
// 	var (
// 		i   int     = 1
// 		f64 float64 = 1.2
// 		s   string  = "text"
// 		t   bool    = true
// 		f   bool    = false
// 	)
// 	fmt.Println(i, f64, s, t, f)
// }

//初期化する場合
func main() {
	var (
		i   int
		f64 float64
		s   string
		t   bool
		f   bool
	)
	fmt.Println(i, f64, s, t, f)

	//関数内でしか宣言することができない
	xi := 1
	xf64 := 1.2
	xs := "test"
	xt := true
	xf := false
	fmt.Println(xi, xf64, xs, xt, xf)

	//変数の型を調べることができる
	fmt.Printf("%T\n", xf64)
}
