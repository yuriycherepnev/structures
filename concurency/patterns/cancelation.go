// Cancellation - Способ прерывания горутин. Необходим, чтобы избегать висящих горутин,
// останавливать слишком долгие операции.

package main

import (
	"context"
	"fmt"
	"time"
)

func cancelWorker(ctx context.Context) {
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
	ctx, cancel := context.WithCancel(context.Background())
	go cancelWorker(ctx)

	time.Sleep(1 * time.Second)
	cancel()
	time.Sleep(1 * time.Second)
}
