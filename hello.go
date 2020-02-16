package main

import "fmt"

// func main() {
// 	var n int = 100
// 	fmt.Println(n)

// 	//100を入れたメモリのアドレスが表示される
// 	fmt.Println(&n)

// 	//*intはintのポイント型という意味である
// 	var p *int = &n
// 	fmt.Println(p)
// 	fmt.Println(*p)
// }

// func one(x int){
// 	x = 1
// }

// func main(){
// 	var n int = 100

// 	//nという値だけをoneという関数に渡している
// 	//oneにはx=1と定義してあるが、影響していない
// 	one(n)
// 	fmt.Println(n)
// }


//1.ポイント型にはアスタリスクをつける
//3.xのままだとアドレスを入れることができないため、
//最終的にはデリファンレンス、実体の中に1を入れる
//xはnのアドレスなので、その中に1を入れている
func one(x *int){
	*x = 1
}

func main(){
	var n int = 100

	//2.ポイント型にはアドレスを渡す必要があるため&をつける
	//100のアドレスをここで渡している
	one(&n)
	fmt.Println(n)
	fmt.Println(&*&n)
}