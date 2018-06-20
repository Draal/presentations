package main
import ( "sync"; "fmt" )
func main() {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	ch := make(chan int)
	var wg sync.WaitGroup; wg.Add(5)
	for p := 0; p < 5; p++ {
		go func() {
			defer wg.Done()
			for {
				v, ok := <-ch
				if !ok { return }
				fmt.Println(v)
			}
		}()
	}
	for _, v := range data {
		ch <- v
	}
	close(ch)
	wg.Wait()
}
