// Fan-in - собирает результаты из нескольких каналов в один.

package main

import (
	"fmt"
	"strconv"
	"sync"
)

func fanIn(channels []chan string) <-chan string {
	inChan := make(chan string)
	var wg sync.WaitGroup

	read := func(input <-chan string) {
		defer wg.Done()
		for v := range input {
			inChan <- v
		}
	}

	wg.Add(len(channels))

	for _, channel := range channels {
		go read(channel)
	}

	go func() {
		wg.Wait()
		close(inChan)
	}()

	return inChan
}

func main() {
	chnCount := 30
	channels := make([]chan string, chnCount)
	for i := 0; i < chnCount; i++ {
		channels[i] = make(chan string)
	}

	for index, channel := range channels {
		go func() {
			defer close(channel)
			number := strconv.Itoa(index)
			channel <- "message " + number + " from channel " + number
		}()
	}

	out := fanIn(channels)

	for msg := range out {
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
