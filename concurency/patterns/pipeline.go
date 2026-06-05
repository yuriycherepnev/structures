/*
Pipeline - данные обрабатываются цепочкой. Producer -> Producer/Consumer -> Consumer.
Стадий обработки может быть сколько угодно.

Реальный сценарий: Система обработки изображений,
где изображение проходит через этапы масштабирования,
фильтрации и кодирования.
*/

package main

import "fmt"

func producer() <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i <= 10; i++ {
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
	input := producer()                  // 1-й этап: генерируем числа
	processed := producerConsumer(input) // 2-й этап: обрабатываем
	consumer(processed)                  // 3-й этап: потребляем результат
}
