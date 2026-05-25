package main

import (
	"fmt"
	"time"
	"unsafe"
)

func main() {
	// здесь на каждой итерации цикла создается новая переменная
	// в го версии < 1.22 переменная в горутины передавалась из замыкания при создании
	for i := 0; i < 5; i++ {
		fmt.Println(unsafe.Pointer(&i))
		go func() {
			fmt.Println(i)
		}()
	}

	time.Sleep(2 * time.Second)
}
