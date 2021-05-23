package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
 * Five silent philosophers sit at a round table with bowls of spaghetti.
 * Forks are placed between each pair of adjacent philosophers.
 * Each philosopher must alternately think and eat.
 * However, a philosopher can only eat spaghetti when they have both left and right forks.
 * Each fork can be held by only one philosopher and so a philosopher can use the fork only if it is not being used by another philosopher.
 * After an individual philosopher finishes eating, they need to put down both forks so that the forks become available to others.
 * A philosopher can only take the fork on their right or the one on their left as they become available
   and they cannot start eating before getting both forks.
 * Eating is not limited by the remaining amounts of spaghetti or stomach space; an infinite supply and an infinite demand are assumed.
 * The problem is how to design a discipline of behavior (a concurrent algorithm) such that no philosopher will starve;
 * i.e., each can forever continue to alternate between eating and thinking,
   assuming that no philosopher can know when others may want to eat or think.
 * Five silent philosophers sit at a round table with bowls of spaghetti. Forks are placed between each pair of adjacent philosophers.
 * Each philosopher must alternately think and eat.
 * However, a philosopher can only eat spaghetti when they have both left and right forks.
 * Each fork can be held by only one philosopher and so a philosopher can use the fork only if it is not being used by another philosopher.
 * After an individual philosopher finishes eating, they need to put down both forks so that the forks become available to others.
 * A philosopher can only take the fork on their right or the one on their left as they
   become available and they cannot start eating before getting both forks.
 * Eating is not limited by the remaining amounts of spaghetti or stomach space;
   an infinite supply and an infinite demand are assumed.
 * The problem is how to design a discipline of behavior (a concurrent algorithm) such that no philosopher will starve;
 * i.e., each can forever continue to alternate between eating and thinking, assuming that no philosopher can know when others
   may want to eat or think.
*/

/*
process P[i]
 while true do
   {  THINK;
      PICKUP(CHOPSTICK[i], CHOPSTICK[i+1 mod 5]);
      EAT;
      PUTDOWN(CHOPSTICK[i], CHOPSTICK[i+1 mod 5])
   }

*/

//initial state - all chopsticks are free

//two states of philospher possible - think/eat
//true - think, false - eat
var philosophers []int = []int{0, 1, 2, 3, 4}

// two states of fork possible - busy/free
// true - free, false - busy
var forks map[int]bool = map[int]bool{
	0: true,
	1: true,
	2: true,
	3: true,
	4: true,
}

var randPh = rand.New(rand.NewSource(time.Now().UnixNano()))

// release forks if any in use
func think(id int, m *sync.Mutex, wg *sync.WaitGroup, ch chan<- int) {
	time.Sleep(2 * time.Second)
	defer wg.Done()
	leftForkId := (id + 4) % 5
	rightForkId := (id + 1) % 5
	ch <- id
	close(ch)
	m.Lock()
	forks[leftForkId] = true
	forks[rightForkId] = true
	fmt.Printf("philosopher %d is thinking\n", id)
	m.Unlock()
}

// occupy 2 forks
func eat(id int, m *sync.Mutex, wg *sync.WaitGroup, ch <-chan int) {
	defer wg.Done()
	leftForkId := (id + 4) % 5
	rightForkId := (id + 1) % 5
	m.Lock()
	forks[leftForkId] = false
	forks[rightForkId] = false
	fmt.Printf("philosopher %d is eating\n", id)
	if _, ok := <-ch; ok {
		time.Sleep(2 * time.Second)
		m.Unlock()
	}
}

func main() {
	m := &sync.Mutex{}
	wg := &sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		ch := make(chan int)
		id := randPh.Intn(5)
		fmt.Println(id)
		wg.Add(1)
		go eat(id, m, wg, ch)
		wg.Add(1)
		go think(id, m, wg, ch)
	}
	wg.Wait()
}
