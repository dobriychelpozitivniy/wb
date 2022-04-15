package main

import (
	"fmt"
	"time"
)

// Реализовать собственную функцию sleep.
func main() {
	sleep(5 * time.Second)
	fmt.Println("End")
}

func sleep(d time.Duration) {
	t := time.NewTicker(d)
	<-t.C
}
