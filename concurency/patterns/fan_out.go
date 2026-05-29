/*
Паттерн Fan-Out (веерный выход) в Go используется для распараллеливания обработки данных:
один входной канал читается несколькими горутинами, каждая из которых выполняет свою работу независимо.
Это позволяет утилизировать все ядра процессора и увеличить пропускную способность.
*/
package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

func worker(workerId int, taskChan <-chan int, resultChan chan<- [2]int, wg *sync.WaitGroup) {
	defer wg.Done()
	for taskId := range taskChan {
		// some code
		time.Sleep(time.Second * 1)
		resultChan <- [2]int{workerId, taskId}
	}
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	taskChan := make(chan int, len(numbers))
	resultChan := make(chan [2]int, len(numbers))

	var wg sync.WaitGroup
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, taskChan, resultChan, &wg)
	}

	for _, number := range numbers {
		taskChan <- number
	}
	close(taskChan)

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	for result := range resultChan {
		workerId := strconv.Itoa(result[0])
		taskId := strconv.Itoa(result[1])
		fmt.Println("workerId " + workerId + " taskId " + taskId)
	}

}
