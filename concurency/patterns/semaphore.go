// Семафор ограничивает количество горутин, которые могут одновременно обращаться к ресурсу.
// Полезен для управления конкурентностью и предотвращения перегрузки ресурсов.

package main

import (
	"fmt"
	"sync"
	"time"
)

func sWorker(id int, sem chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	sem <- struct{}{} // Захват семафора
	fmt.Printf("Воркер %d начал\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Воркер %d завершил\n", id)
	<-sem // Освобождение семафора
}

func main() {
	const numWorkers = 5
	const maxConcurrent = 2
	sem := make(chan struct{}, maxConcurrent)
	var wg sync.WaitGroup

	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go sWorker(i, sem, &wg)
	}

	wg.Wait()
}

/*
Реальный сценарий: Пул подключений к базе данных,
где одновременно допускается ограниченное количество подключений.
*/
