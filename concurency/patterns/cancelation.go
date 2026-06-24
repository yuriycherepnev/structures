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
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Done")
			return
		default:
			fmt.Println("Working...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	wg := &sync.WaitGroup{}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	wg.Add(1)
	go cancelWorker(wg, ctx)

	wg.Wait()
}
