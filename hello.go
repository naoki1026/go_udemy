package main

import "fmt"

func main() {
	l := []string{"phthon", "go", "java"}
	for i := 0; i < len(l); i++ {
		fmt.Println(i, l[i])
	}

	//lの中身を取ってきてindexと値を表示する
	for i, v := range l {
		fmt.Println(i, v)
	}

	//forをアンスコにすることでindexを表示させないこともできる
	for _, v := range l {
		fmt.Println(v)
	}

	m := map[string]int{"apple": 100, "banana" : 200}
	for k, v := range m {
		fmt.Println(k, v)
	}
	for k := range m {
		fmt.Println(k)
	}
	for _, v := range m {
		fmt.Println(v)
	}
}
