package main

import (
	"fmt"
	"time"
)

func main() {
	// var c chan string
	c := make(chan string)

	// enviar em uma goroutine
	go func() {
		i := 1
		for {
			time.Sleep(time.Second * 1)
			i++
			msg := fmt.Sprintf("Ola %d", i)
			c <- msg
		}
	}()

	// escutar do canal
	go func() {
		for {
			time.Sleep(time.Second * 2)
			msg := <-c
			fmt.Println(msg)
		}
	}()
	time.Sleep(time.Second * 10)
}
