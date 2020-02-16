package main

import "fmt"

func main() {
	defer fmt.Println("world")
	fmt.Println("HelloWorld")
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
}
