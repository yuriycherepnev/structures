// Fan-in - собирает результаты из нескольких каналов в один.

package main

import (
	"fmt"
	"strconv"
	"sync"
)

func fanIn(channels []chan string) <-chan string {
	ch := make(chan string)
	var wg sync.WaitGroup

	read := func(input <-chan string) {
		defer wg.Done()
		for v := range input {
			ch <- v
		}
	}

	wg.Add(len(channels))

	for _, channel := range channels {
		go read(channel)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	return ch
}

func main() {
	chnCount := 30
	channels := make([]chan string, chnCount)
	for i := 0; i < chnCount; i++ {
		channels[i] = make(chan string)
	}

	for index, value := range channels {
		go func() {
			defer close(value)
			number := strconv.Itoa(index)
			value <- "message " + number + " from channel " + number
		}()
	}

	out := fanIn(channels)

	for msg := range out {
		fmt.Println(msg)
	}
}

// select ждёт, когда хотя бы один из каналов станет готов к чтению

// когда мы так читаем канал:
// for v := range input {
// мы блокируем текущую горутину до тех пор, пока канал не будет готов к чтению
// Во время ожидания следующего сообщения цикл не потребляет процессор
// это блокировка на уровне runtime
// цикл не "крутится" и не "пингует". Он эффективно ждёт, используя механизмы внутреннего планировщика Go
