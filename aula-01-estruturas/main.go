package main

import "fmt"

type Car struct {
	Name    string
	Age     int
	ModelNo int
}

func (c Car) Drive() {
	fmt.Println("driving...")
}

func Anything(value interface{}) {
	fmt.Println(value)
}

func main() {
	var c Car
	c = Car{
		Name:    "chevy",
		Age:     12,
		ModelNo: 33,
	}
	mymap := make(map[string]interface{})
	mymap["name"] = "Oshiro"
	mymap["age"] = 40

	fmt.Println(c)
	c.Drive()

	Anything(mymap)
}
