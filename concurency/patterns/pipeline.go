/*
Pipeline - данные обрабатываются цепочкой. Producer -> Producer/Consumer -> Consumer.
Стадий обработки может быть сколько угодно.

Реальный сценарий: Система обработки изображений,
где изображение проходит через этапы масштабирования,
фильтрации и кодирования.

producer - это паттерн генератор
*/
package main

import "fmt"

func producer() <-chan int {
	c := make(chan int)

	go func() {
		for i := 1; i <= 10; i++ {
			c <- i
		}
		close(c)
	}()

	return c
}

func producerConsumer(c <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		for v := range c {
			out <- v * v
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

func main() {
	input := producer()
	processed := producerConsumer(input)
	consumer(processed)
}
