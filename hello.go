package main

import "fmt"

func main() {
	for i := 1; i < 10; i++ {
		fmt.Println(i)

		//続ける
		if i == 3 {
			fmt.Println("continue")
			continue
		}

		//ここでfor文がストップする
		if i == 6 {
			fmt.Println("break")
			break
		}

	}
	sum := 1
	for sum < 10 {
		sum += sum
		fmt.Println(sum)
	}
}
