/*
Worker pool - каждый воркер берёт задачу, делает работу и отправляет результат в канал,
другая горутина, в нашем случае main, читает результат из канала.

Реальный сценарий: Веб-сервер, обрабатывающий входящие HTTP-запросы,
где каждый запрос обрабатывается воркером из пула.
*/

package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		time.Sleep(1 * time.Second)
		results <- job
	}
}

func main() {
	jobs := make(chan int, 10)
	results := make(chan int, 10)
	wg := &sync.WaitGroup{}

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go worker(jobs, results, wg)
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
алгоритм создания паттерна:
main
1. channels init (jobs results)
2. waitGroup init
3. for init воркеров, привязка их к каналам и добавление групп wg.Add
4. пуш данных через горутину в канал jobs и закрытие его после отправки сообщений
5. запуск в горутине wg.Wait() и ожидание закрытия канала results
6. получение данных из канала results

worker
1. закрытие группы wg через defer
2. получение данных из jobs
3. обработка
4. пуш данных в results
*/
