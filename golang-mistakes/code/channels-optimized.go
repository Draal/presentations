package main

import (
	"fmt"
	"sync"
)

func main() {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	var wg sync.WaitGroup
	wg.Add(5)
	for p := 0; p < 5; p++ {
		go func(producerID int) {
			defer wg.Done()
			for idx, v := range data {
				if (idx % 5) == producerID {
					fmt.Println(v)
				}
			}
		}(p)
	}
	wg.Wait()
}
