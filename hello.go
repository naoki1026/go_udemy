package main

import "fmt"

func main() {
	m := map[string]int{"apple": 100, "orange": 200}
	fmt.Println(m["apple"])
	m["new"] = 400
	fmt.Println(m)

	//値が存在するかを確認できる
	v, ok := m["apple"]
	fmt.Println(v, ok)

	//値が存在するかを確認できる
	v2, ok2 := m["nothing"]
	fmt.Println(v2, ok2)

	//mapの初期化
	var m2 = make(map[string]int)
	m2["banana"] = 200
	fmt.Println(m2)
}
