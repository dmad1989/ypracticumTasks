package main

import (
	"fmt"
	"sync"
)

func Foo(m sync.Mutex) {
	m.Lock()
	defer m.Unlock()
	// какие-то действия
}

func main() {
	var wg sync.WaitGroup
	for _, v := range []int{0, 1, 2, 3} {
		wg.Add(1)
		go func() {
			fmt.Print(v)
			wg.Done()
		}()
	}
	wg.Wait()
}
