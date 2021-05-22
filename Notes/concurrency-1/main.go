package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var cache = map[int]Book{}
var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))
var wg sync.WaitGroup

func main() {
	// wg := &sync.WaitGroup{}
	m := &sync.RWMutex{}
	for i := 0; i < 20; i++ {
		id := rnd.Intn(10) + 1
		// fmt.Printf("count: %v\n ", id)
		wg.Add(2)
		go func(id int, m *sync.RWMutex) {
			defer wg.Done()
			if b, ok := queryCache(id, m); ok {
				fmt.Println("from cache:")
				fmt.Println(b)
			}
		}(id, m)
		go func(id int, m *sync.RWMutex) {
			defer wg.Done()
			if b, ok := queryDatabase(id, m); ok {
				fmt.Println("from db:")
				fmt.Println(b)
			}
		}(id, m)
		// fmt.Printf("Book not found with id: %v", id)
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
