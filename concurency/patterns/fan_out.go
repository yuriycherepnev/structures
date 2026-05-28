/*
Паттерн Fan-Out (веерный выход) в Go используется для распараллеливания обработки данных:
один входной канал читается несколькими горутинами, каждая из которых выполняет свою работу независимо.
Это позволяет утилизировать все ядра процессора и увеличить пропускную способность.
*/

package main

import (
	"fmt"
	"net/http"
	"sync"
)

func worker(id int, tasks <-chan string, results chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	for url := range tasks {
		resp, err := http.Get(url)
		if err != nil {
			results <- string('0')
			continue
		}
		err = resp.Body.Close()
		if err != nil {
			return
		}
		fmt.Println(id)

		results <- resp.Status
	}
}

func main() {
	urls := []string{
		"https://google.com",
		"https://yandex.ru",
		"https://github.com",
		"https://stackoverflow.com",
		"https://golang.org",
	}

	tasks := make(chan string, len(urls))
	results := make(chan string, len(urls))

	numWorkers := 3
	var wg sync.WaitGroup
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, tasks, results, &wg)
	}

	for _, url := range urls {
		tasks <- url
	}
	close(tasks)

	go func() {
		wg.Wait()
		close(results)
	}()

	for length := range results {
		fmt.Println(length)
	}

}
