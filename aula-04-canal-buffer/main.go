package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Car struct {
	Name string
	Year int
}

func Loja(name string, tempo int, canal <-chan *Car) {
	go func() {
		for {
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(tempo)))
			i := <-canal
			fmt.Printf("Vendido %s ano %d na loja %s, %d carros no patio\n", i.Name, i.Year, name, len(canal))
		}
	}()
}

func main() {

	c := make(chan *Car, 3)
	models := []string{"Fusca", "Gol", "Lambo", "Porsche", "Ferrary"}

	// produtor
	go func() {
		i := 0
		for {
			i++
			r := rand.Intn(150)
			modelName := models[rand.Intn(len(models))]
			time.Sleep(time.Millisecond * time.Duration(r))
			c <- &Car{Name: modelName, Year: 1950 + i}
		}
	}()

	// consumidor
	Loja("A", 300, c)
	Loja("B", 300, c)
	time.Sleep(time.Second * 10)
}
