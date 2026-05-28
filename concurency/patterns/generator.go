// Generator - микропаттерн, который наполняет канал.
package main

import "fmt"

func generator() <-chan int {
	ch := make(chan int)
	go func() {
		for i := 0; i <= 12; i++ {
			ch <- i + 1
		}
		close(ch)
	}()
	return ch
}

func main() {
	intChan := generator()

	for value := range intChan {
		fmt.Println(value)
	}
}
