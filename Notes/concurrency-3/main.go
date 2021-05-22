package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var cache = map[int]Book{}
var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

// in concurrency-1 we did not have any sync between cache and db goroutines
// in this file we will use channels to bring about sync
func main() {
	wg := &sync.WaitGroup{}
	m := &sync.RWMutex{}
	cacheCh := make(chan Book)
	dbCh := make(chan Book)

	for i := 0; i < 10; i++ {
		id := rnd.Intn(10) + 1
		wg.Add(2)
		// send only channel used int cache go routine
		go func(id int, m *sync.RWMutex, wg *sync.WaitGroup, ch chan<- Book) {
			defer wg.Done()
			if b, ok := queryCache(id, m); ok {
				ch <- b
			}
		}(id, m, wg, cacheCh)
		// send only channel used in db go routine
		go func(id int, m *sync.RWMutex, wg *sync.WaitGroup, ch chan<- Book) {
			defer wg.Done()
			if b, ok := queryDatabase(id, m); ok {
				ch <- b
			}
		}(id, m, wg, dbCh)
		// third goroutine(recieve only for both db and cache) to actually moderate between the cache and db
		go func(cacheCh, dbCh <-chan Book) {
			select {
			case b := <-cacheCh:
				fmt.Println("from cache:")
				fmt.Println(b)
				<-dbCh // we put in this to make sure if cache is called we drain value from db channel, because db channel puts value in each iteration and cache puts in value in only some times
			case b := <-dbCh:
				fmt.Println("from db:")
				fmt.Println(b)
			}
		}(cacheCh, dbCh)
		time.Sleep(150 * time.Millisecond)
	}
	wg.Wait()
}

func queryCache(id int, m *sync.RWMutex) (Book, bool) {
	m.RLock()
	b, ok := cache[id]
	m.RUnlock()
	return b, ok
}

func queryDatabase(id int, m *sync.RWMutex) (Book, bool) {
	// time.Sleep(100 * time.Millisecond)
	for _, b := range books {
		if b.ID == id {
			m.Lock()
			cache[id] = b
			m.Unlock()
			return b, true
		}
	}
	return Book{}, false
}
