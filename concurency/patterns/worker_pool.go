// Worker pool - каждый воркер берёт задачу, делает работу и отправляет результат в канал,
// другая горутина, в нашем случае main, читает результат из канала.

package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for j := range jobs {
		time.Sleep(1 * time.Second)
		fmt.Println("job", j)
		results <- j * j
	}
}

func main() {
	jobs := make(chan int)
	results := make(chan int)

	var wg sync.WaitGroup

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go worker(jobs, results, &wg)
	}

	go func() {
		for i := 0; i < 10; i++ {
			jobs <- i
		}
		close(jobs)
	}()

	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		fmt.Println(result)
	}
}

/*
схема воркера:
main
1. channels init (jobs results)
2. waitGroup init
3. for запуск воркеров в горутинах, привязка их к каналам и добавление групп wg.Add
4. пуш данных через горутину в канал jobs и закрытие его после отправки сообщений
5. запуск в горутине wg.Wait() и ожидание закрытия канала results
6. получение данных из канала results

worker
1. закрытие группы wg.Done через defer
2. получение данных из jobs
3. обработка
4. пуш данных в results
*/
