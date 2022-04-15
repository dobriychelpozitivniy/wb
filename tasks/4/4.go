package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// Реализовать постоянную запись данных в канал (главный поток). Реализовать набор из N воркеров,
// которые читают произвольные данные из канала и выводят в stdout.
// Необходима возможность выбора количества воркеров при старте.
// Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ завершения работы всех воркеров.

func main() {
	countWorkers := flag.Int("count", 2, "worker count")
	flag.Parse()

	values := make(chan string, 5)
	quit := make(chan bool)
	// WaitGroup чтобы дождаться что все значения прочитаны
	wg := sync.WaitGroup{}

	go runWriter(&wg, &values, &quit)
	fmt.Println("writer run")

	runWorkers(&wg, &values, &quit, *countWorkers)
	fmt.Println("workers run")

	termChan := make(chan os.Signal)
	signal.Notify(termChan, syscall.SIGINT, syscall.SIGTERM)

	<-termChan
	quit <- true
	fmt.Println("Ждем завершния работы...")
	wg.Wait()
	fmt.Println("Работа программы завершена")
}

func runWriter(wg *sync.WaitGroup, ch *chan string, quit *chan bool) {
	t := time.NewTicker(50 * time.Millisecond)

L:
	for {
		select {
		case <-t.C:
			wg.Add(1)
			s := fmt.Sprintf("awdjawudhawuieh %s", time.Now().String())
			*ch <- s
		case <-*quit:
			break L
		}
	}

	fmt.Println("BREAK")

	t.Stop()
	close(*ch)
}

func runWorkers(wg *sync.WaitGroup, ch *chan string, quit *chan bool, count int) {
	for i := 0; i < count; i++ {
		go func() {
			for {
				select {
				case v := <-*ch:
					time.Sleep(50 * time.Millisecond)
					fmt.Println(v)
					wg.Done()
				}
			}
		}()
	}
}
