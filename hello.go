package main

import (
	"fmt"
	"strings"
)

func main() {
	var s = "Hello World!"
	//○番目のアルファベットを取り出すことができる
	fmt.Println(string("HelloWorld!"[1]))

	//条件に合致する初めの１回だけ文字を置き換えることができる
	//実際のs自体は置き換わらないので注意
	fmt.Println(strings.Replace(s, "H", "K", 1))

	//直接sの中に入れることで文字を置き換えることができる
	s = strings.Replace(s, "H", "X", 1)
	fmt.Println(s)
	fmt.Println("test\n" +
		"test")
	fmt.Println(`test
test`)

}
