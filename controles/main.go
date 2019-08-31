package main

import "fmt"

func IfElse(value int) {
	if value == 1 {
		fmt.Println("um")
	} else if value == 2 {
		fmt.Println("dois")
	} else {
		fmt.Println("nem um nem dois")
	}
}

func main() {

	// if else for
	fmt.Println("Ola Mundo")
	IfElse(1)
	IfElse(2)
	IfElse(4)
}
