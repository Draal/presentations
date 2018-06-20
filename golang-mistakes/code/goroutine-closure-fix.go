package main

import (
	"fmt"
	"sync"
)

var slice = []int{1, 2, 3, 4, 5}

func main() {
	var wg sync.WaitGroup
	wg.Add(len(slice))
	for i := range slice {
		go func(v int) {
			defer wg.Done()
			fmt.Println(v)
		}(i)
	}
	wg.Wait()
}
