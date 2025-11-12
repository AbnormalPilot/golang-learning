package main

import "fmt"

func hello() {
	fmt.Println("Hello World")
	var name = "Jash"
	var age int = 18
	var no float64 = 64.3234
	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(no)

	var hello string = "hello"
	hi := "hi"
	fmt.Println(hello)
	fmt.Println(hi)

	if(age>=18){
		fmt.Println("Allowed")
	}
	

}