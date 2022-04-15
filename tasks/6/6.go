package main

import (
	"context"
	"fmt"
	"time"
)

// Реализовать все возможные способы остановки выполнения горутины.
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	someWork := make(chan int, 2)
	someWork2 := make(chan int, 2)
	close := make(chan bool)

	t := time.NewTicker(500 * time.Millisecond)

	go func() {
		for _ = range t.C {
			someWork <- 4
			someWork2 <- 10
		}
	}()

	go func() {
		time.Sleep(2 * time.Second)
		cancel()
		time.Sleep(4 * time.Second)
		close <- true
	}()

	// context
	go func(ctx context.Context) {
	L:
		for {
			select {
			case v := <-someWork:
				fmt.Println(v)
			case <-ctx.Done():
				fmt.Println("Context goroutine is canceled")
				break L
			}
		}
	}(ctx)

	// channel
	go func(ctx context.Context) {
		for {
			select {
			case <-someWork2:
				fmt.Println("Work")
			case <-close:
				fmt.Println("Channel goroutine is close")
				return
			}
		}
	}(ctx)

	<-close
}
