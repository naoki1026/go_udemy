package main

import "fmt"

func thirdPartyConnectDB(){
	//panicは自分で何を修正すれば良いのかわからない時なので、
	//そのようなことがないようにしておくことが大事である
	panic("Unable to conenct to Database")
}

func save(){
	defer func(){

		//panicでシステムが強制終了しないようにrecoverしている
		s := recover()
		fmt.Println(s)
		}()
		thirdPartyConnectDB()
	}

func main() {
	save()
	fmt.Println("OK?")
}
