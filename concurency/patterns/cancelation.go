// Cancellation - Способ прерывания горутин. Необходим, чтобы избегать висящих горутин,
// останавливать слишком долгие операции.

package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func cancelWorker(wg *sync.WaitGroup, ctx context.Context) {
	defer wg.Done()
	ch := make(chan int)
	go func() {
		time.Sleep(5 * time.Second)
		close(ch)
	}()
	select {
	case <-ch:
		fmt.Println("work done")
		return
	case <-ctx.Done():
		fmt.Println("cancel context")
		return
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go cancelWorker(wg, ctx)

	wg.Wait()
}

/*
алгоритм создания паттерна:
main
1. создаем wg
2. создаем контекст и функцию отмены
3. привязываем отмену к defer
4. добавляем группу
5. запускаем воркера и передаем в него группу и контекст
6. ожидаем окончание группы

cancelWorker
1. закрытие группы через дефер
2. обрабатываем закрытие контекста через for + select
3. проверяем через промежутки времени в default
*/
