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
	chn := generator()

	fmt.Println(<-chn)
	fmt.Println(<-chn)
	fmt.Println(<-chn)
	fmt.Println(<-chn)
}
