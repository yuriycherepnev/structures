package main

import (
	"fmt"
	"sync"
	"time"
)

func scWorker(id int, ready *bool, cond *sync.Cond, wg *sync.WaitGroup) {
	defer wg.Done()
	cond.L.Lock()
	for !*ready {
		fmt.Println(id, "wait")
		cond.Wait()
	}
	cond.L.Unlock()
	fmt.Println(id, "done")
}

func main() {
	var mu sync.Mutex
	cond := sync.NewCond(&mu)

	ready := false

	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go scWorker(i, &ready, cond, &wg)
	}

	time.Sleep(time.Second)

	for i := 0; i < 5; i++ {
		mu.Lock()
		ready = true
		cond.Signal()
		mu.Unlock()
		time.Sleep(time.Millisecond * 1000)
	}

	wg.Wait()
}
