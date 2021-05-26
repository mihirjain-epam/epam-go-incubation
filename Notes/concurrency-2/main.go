package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	ch := make(chan int)
	wg.Add(2)

	go func(ch <-chan int, wg *sync.WaitGroup) {
		defer wg.Done()
		// if fromChannel, ok := <-ch; ok {
		// 	fmt.Println(fromChannel)
		// }
		// for i := 0; i < 10; i++ {
		// 	fmt.Println(<-ch)
		// }
		for msg := range ch {
			fmt.Println(msg)
		}
	}(ch, wg)
	go func(ch chan<- int, wg *sync.WaitGroup) {
		defer wg.Done()

		for i := 0; i < 10; i++ {
			ch <- i
		}
		// close(ch) // ch <- 0
		// toChannel := 2
		// ch <- toChannel
		// ch <- 0
	}(ch, wg)
	wg.Wait()
}
