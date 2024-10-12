package main

import "fmt"

func producer(ch chan int) {
	for i := 1; i <= 10; i++ {
		fmt.Printf("Numero enviado: %d\n", i)
		ch <- i
	}
	close(ch)
}

func main() {
	ch := make(chan int)

	go producer(ch)

	for num := range ch {
		fmt.Printf("Numero recebido: %d\n", num)
	}
}
