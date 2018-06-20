package main

import "sync"

var unprotectedMap = map[int]int{}

func main() {
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(v int) {
			defer wg.Done()
			unprotectedMap[v] = v
		}(i)
	}
	wg.Wait()
}
