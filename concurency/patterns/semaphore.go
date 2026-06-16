/*
Семафор ограничивает количество горутин, которые могут одновременно обращаться к ресурсу.
Полезен для управления конкурентностью и предотвращения перегрузки ресурсов.

Реальный сценарий: Пул подключений к базе данных,
где одновременно допускается ограниченное количество подключений.
*/

package main

import (
	"fmt"
	"sync"
	"time"
)

func sWorker(id int, sem chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	sem <- struct{}{}
	fmt.Println(id)
	time.Sleep(time.Millisecond * 300)
	<-sem
}

func main() {
	sem := make(chan struct{}, 2)
	var wg sync.WaitGroup
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go sWorker(i, sem, &wg)
	}
	wg.Wait()
}

/*
алгоритм создания паттерна:
main
1. channels init (sem)
2. waitGroup init
3. for init воркеров, привязка их к каналам и добавление групп wg.Add
4. wg wait

worker
1. закрытие группы wg через defer
2. удержание канала
3. обработка данных
4. освобождение канала
*/
