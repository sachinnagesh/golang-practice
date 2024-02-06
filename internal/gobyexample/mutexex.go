package gobyexample

import (
	"fmt"
	"sync"
)

type Container struct{
	mu sync.Mutex
	counters map[string]int
}

func (c *Container) inc(name string){
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counters[name]++
}

func MutexExamp(){

	c := Container{
		counters: map[string]int{},
	}

	var wg sync.WaitGroup

	incrementFunc := func (name string, n int)  {
		for i := 0 ; i < n;i++ {
			c.inc(name)
		}
		wg.Done()
	}

	wg.Add(4)
	go incrementFunc("a",2000)
	go incrementFunc("a",2000)
	go incrementFunc("b",2000)
	go incrementFunc("b",2000)
	wg.Wait()
	fmt.Println(c.counters)

}