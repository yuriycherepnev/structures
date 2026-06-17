// Fan-in - собирает результаты из нескольких каналов в один

package main

import (
	"fmt"
	"sync"
)

func fanIn(channels []chan int) chan int {
	var wg sync.WaitGroup
	results := make(chan int)

	for _, channel := range channels {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for v := range channel {
				results <- v
			}
		}()
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	return results
}

func main() {
	chnCount := 30
	channels := make([]chan int, chnCount)
	for i := 0; i < chnCount; i++ {
		channels[i] = make(chan int)
	}

	for index, channel := range channels {
		go func() {
			defer close(channel)
			channel <- index
		}()
	}

	results := fanIn(channels)

	for msg := range results {
		fmt.Println(msg)
	}
}

/*
select ждёт, когда хотя бы один из каналов станет готов к чтению

когда мы так читаем канал:
for v := range input {
мы блокируем текущую горутину до тех пор, пока канал не будет готов к чтению
Во время ожидания следующего сообщения цикл не потребляет процессор
это блокировка на уровне runtime
цикл не "крутится" и не "пингует". Он эффективно ждёт, используя механизмы внутреннего планировщика Go

range будет читать значения из канала до тех пор, пока канал не будет закрыт

если канал никогда не закрыть, цикл range будет вечно ждать новых значений — возникнет deadlock

После закрытия канала:
Все оставшиеся сообщения останутся в канале и будут доступны для чтения.
Цикл range (или обычное чтение <-ch) сначала прочитает все сообщения,
которые были в канале на момент закрытия.
После того как последнее сообщение будет прочитано, последующие попытки чтения будут возвращать нулевое значение и
false (при использовании val, ok := <-ch), а цикл range завершится.

Нельзя отправлять в закрытый канал — паника.
Можно читать из закрытого канала до опустошения буфера.

*/
