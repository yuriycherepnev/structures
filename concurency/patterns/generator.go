// Generator - микропаттерн, который наполняет канал. Закрываем канал, чтобы не было проблем.
// паттерн генератор не должен содержать в себе блокирующих операций

package main

import (
	"fmt"
)

func generator() <-chan int {
	ch := make(chan int)

	go func() {
		for i := 0; i <= 10; i++ {
			ch <- i
		}
		close(ch)
	}()

	return ch
}

func main() {
	ch := generator()

	for result := range ch {
		fmt.Println(result)
	}
}

/*
алгоритм создания паттерна:
generator
1. создаем канал
2. пушим данные в канал + закрываем
3. возвращаем канал

main
1. вызываем generator
2. читаем данные из канала
*/
