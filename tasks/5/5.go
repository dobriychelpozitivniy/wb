package main

import (
	"flag"
	"fmt"
	"sync"
	"time"
)

// Разработать программу, которая будет последовательно отправлять значения в канал, а с другой стороны канала — читать.
// По истечению N секунд программа должна завершаться.

func main() {
	seconds := flag.Int("seconds", 10, "seconds work app")
	flag.Parse()

	wg := sync.WaitGroup{}
	ch := make(chan int)
	quit := make(chan bool)

	go func() {
		time.Sleep(time.Duration(*seconds) * time.Second)
		quit <- true
	}()

	runWrite(&wg, &ch, &quit)

	go func() {
		for {
			select {
			case <-quit:
				return
			case v := <-ch:
				fmt.Println(v)
				wg.Done()
			}
		}
	}()

	<-quit
	wg.Wait()
}

func runWrite(wg *sync.WaitGroup, ch *chan int, quit *chan bool) {
	t := time.NewTicker(500 * time.Millisecond)

	go func() {
		for {
			select {
			case <-*quit:
				t.Stop()
				return
			case <-t.C:
				wg.Add(1)
				*ch <- 10
			}
		}
	}()
}
