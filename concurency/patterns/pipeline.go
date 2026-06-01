/*
Pipeline - данные обрабатываются цепочкой. Producer -> Producer/Consumer -> Consumer.
Стадий обработки может быть сколько угодно.
*/

package main

import "fmt"

func main() {
	prod := producer()
	cons := producerConsumer(prod)
	consumer(cons)
}

func producer() <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i <= 10; i++ {
			c <- i + 1
		}
		close(c)
	}()

	return c
}

func producerConsumer(c <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		for v := range c {
			out <- v * 2
		}
		close(out)
	}()

	return out
}

func consumer(ch <-chan int) {
	for v := range ch {
		fmt.Println(v)
	}
}
