package main

import (
	"fmt"
	"strconv"
	"sync"
)

// Реализовать конкурентную запись данных в map.
func main() {
	cm := ConcurencyMap{
		m:  make(map[string]int),
		mu: &sync.Mutex{},
	}

	wg := sync.WaitGroup{}
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(v int) {
			cm.Add(strconv.Itoa(v), v)
			wg.Done()
		}(i)
	}

	wg.Wait()
	for _, v := range cm.m {
		fmt.Println(v)
	}

	cmrw := ConcurencyMapRW{
		m:  make(map[string]int),
		mu: &sync.RWMutex{},
	}

	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(v int) {
			cmrw.Add(strconv.Itoa(v), v)
			wg.Done()
		}(i)
	}

	wg.Wait()
	for _, v := range cm.m {
		fmt.Println(v)
	}

	cmsync := ConcurencyMapSync{
		m: sync.Map{},
	}

	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(v int) {
			cmsync.Add(strconv.Itoa(v), v)
			wg.Done()
		}(i)
	}

	wg.Wait()
	for _, v := range cm.m {
		fmt.Println(v)
	}
}

// Mutex
type ConcurencyMap struct {
	mu *sync.Mutex
	m  map[string]int
}

func (c *ConcurencyMap) Add(key string, value int) {
	c.mu.Lock()
	c.m[key] = value
	c.mu.Unlock()
}

// RWMutex
type ConcurencyMapRW struct {
	mu *sync.RWMutex
	m  map[string]int
}

func (c *ConcurencyMapRW) Add(key string, value int) {
	c.mu.Lock()
	c.m[key] = value
	c.mu.Unlock()
}

// sync.map
type ConcurencyMapSync struct {
	m sync.Map
}

func (c *ConcurencyMapSync) Add(key string, value int) {
	c.m.Store(key, value)
}
