package main

import "fmt"
// func main(){

// 	//値がない状態でメモリだけ確保したい
// 	//メモリの領域が確保される
// 	var p *int = new(int)
// 	fmt.Println(p)

// 	//メモリの領域が確保されていない
// 	var p2 *int
// 	fmt.Println(p2)
// }

func main(){
	s := make([]int, 0)
	fmt.Printf("%T\n", s)

	m := make(map[string]int)
	fmt.Printf("%T\n", m)

	var p *int= new(int)
	fmt.Printf("%T\n", p)
}

// 	var p *int = new(int)
// 	fmt.Println(*p)
// 	*p++
// 	fmt.Println(*p)


// 	var p2 *int
// 	fmt.Println(p2)
// }