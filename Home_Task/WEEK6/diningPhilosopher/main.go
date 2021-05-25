package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"sync"
	"time"
)

const commands = `
Welcome to dining-philosphers restaurant. Choose your pick of philosphers:

	> 1   : 2 philosophers next to each other(check for race condition)
	> 2   : 2 philosophers not next to each other(check for efficiency in non blocking case)
	> 3   : randomised run
	> 4   : quits

`

type fork struct {
	mutex  *sync.Mutex
	isFree bool
}

type philosopher struct {
	id                  int
	leftFork, rightFork *fork
}

var randPh = rand.New(rand.NewSource(time.Now().UnixNano()))

func main() {
	wg := &sync.WaitGroup{}
	forks := make([]*fork, 5)
	for i := 0; i < 5; i++ {
		forks[i] = &fork{&sync.Mutex{}, true}
	}
	philosophers := make([]*philosopher, 5)
	for i := 0; i < 5; i++ {
		philosophers[i] = &philosopher{i, forks[(i)%5], forks[(i+1)%5]}
	}
	runByUsersChoice(philosophers, forks, wg)
	wg.Wait()
}

/*
 eat method allows philosopher to simulate eat activity concurrently

 Input arguments -> wg *sync.WaitGroup - wait group

 Method pointer reciever -> p *philosopher - bind to philosopher struct

 Return value -> nil
*/
func (p *philosopher) eat(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("%d Not yet obtained locks :%d \n", time.Now().Unix(), p.id)
	p.leftFork.mutex.Lock()
	p.rightFork.mutex.Lock()
	p.leftFork.isFree = false
	p.rightFork.isFree = false

	fmt.Printf("%d Starting to eat:%d \n", time.Now().Unix(), p.id)
	time.Sleep(2 * time.Second)
	fmt.Printf("%d finishing eating:%d \n", time.Now().Unix(), p.id)

	p.leftFork.isFree = true
	p.rightFork.isFree = true
	p.rightFork.mutex.Unlock()
	p.leftFork.mutex.Unlock()
	fmt.Printf("%d released locks :%d \n", time.Now().Unix(), p.id)

}

/*
 runByUsersChoice function allows user to choose which condition needs to be checked

 Input arguments ->

 philosophers []*philosopher - slice containing all philosophers

 forks []*fork - slice containing all forks

 wg *sync.WaitGroup - wait group

 Return value -> nil
*/
func runByUsersChoice(philosophers []*philosopher, forks []*fork, wg *sync.WaitGroup) {
	scanner := bufio.NewScanner(os.Stdin)
	exitFlag := false
	for !exitFlag {
		fmt.Print(commands)
		scanner.Scan()
		command := strings.Trim(scanner.Text(), " ")
		if len(command) != 0 {
			switch command {
			case "1":
				p1 := randPh.Intn(5)
				p2 := (p1 + 1) % 5
				wg.Add(2)
				go philosophers[p1].eat(wg)
				go philosophers[p2].eat(wg)
				wg.Wait()
			case "2":
				p1 := randPh.Intn(5)
				p2 := (p1 + 2) % 5
				wg.Add(2)
				go philosophers[p1].eat(wg)
				go philosophers[p2].eat(wg)
				wg.Wait()
			case "3":
				for i := 0; i < 5; i++ {
					id := randPh.Intn(5)
					wg.Add(1)
					go philosophers[id].eat(wg)
				}
				wg.Wait()
			case "4":
				exitFlag = true
			default:
				fmt.Print("\nPlease select from the provided choice:\n")
				continue
			}
		}
	}
}
