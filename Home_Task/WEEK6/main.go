package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"sync"
)

/*
Write a program which reads full path to the directory and return statistic about how many files this directory have grouped by file extension
-	the initial function should start with file descriptor
-	create the buffered channel as a value you can use map[string]int
-	create wait group (sync.WaitGroup)
-	go for each child item
-	if file descriptor referring to file, increment a local counter for file extension
-	if file descriptor referring to the directory, increment wait group counter, start goroutine and pass result channel as a parameter to it,
	the function should make the same operations go over subitems and depends on if it’s directory or file increment counter or start the daemon
-	when child daemon is over decrement wait group counter

-	read and merge results from the channel, when all child item is processed return summarized result to the caller

As a result, you will get a recursive program which can concurrently scrab statistic and print it to console

*/
const dirPath = "directory"

func main() {
	fileCh := make(chan map[string]int, 1)
	wg := &sync.WaitGroup{}
	wgRecursion := &sync.WaitGroup{}
	combinedMap := make(map[string]int)
	wg.Add(1)
	go fileConcurrentWalk(wg, wgRecursion, fileCh, dirPath, true)
	wg.Add(1)
	go mergeResults(wg, fileCh, combinedMap)
	wg.Wait()
	printCombinedResult(combinedMap)
}

/*
 * Input parameters -
 * 			wg *sync.WaitGroup				- wait group for first time call of function
 * 			wgRecursion *sync.WaitGroup		- wait group for every non-first call of function
 * 			fileCh chan<- map[string]int	- recieve only channel
 * 			inputDirPath string				- directory path
 * 			firstTime bool					- flag to identify first time call of function
 *
 * Return value -
 * 				nil
 */
func fileConcurrentWalk(wg *sync.WaitGroup, wgRecursion *sync.WaitGroup, fileCh chan<- map[string]int, inputDirPath string, firstTime bool) {
	if firstTime {
		defer wg.Done()
	} else {
		defer wgRecursion.Done()
	}
	files, err := ioutil.ReadDir(inputDirPath)
	if err != nil {
		log.Fatal(err)
	}
	extMap := make(map[string]int)
	for _, file := range files {
		if !file.IsDir() {
			fileExt := filepath.Ext(file.Name())
			extMap[fileExt] += 1
		}
		if file.IsDir() {
			wgRecursion.Add(1)
			go fileConcurrentWalk(wg, wgRecursion, fileCh, inputDirPath+"\\"+file.Name(), false)
		}
	}
	fileCh <- extMap
	if firstTime {
		wgRecursion.Wait()
		close(fileCh)
	}
}

/*
 * Input parameters -
 * 			wg *sync.WaitGroup				- wait group for first time call of function
 * 			fileCh chan<- map[string]int	- recieve only channel
 * 			combinedMap map[string]int		- final result map
 *
 * Return value -
 * 				nil
 */
func mergeResults(wg *sync.WaitGroup, fileCh <-chan map[string]int, combinedMap map[string]int) {
	defer wg.Done()
	for extMap := range fileCh {
		for k, v := range extMap {
			combinedMap[k] += v
		}
	}
}

/*
 * Input parameters -
 * 			combinedMap map[string]int		- final result map
 *
 * Return value -
 * 				nil
 */
func printCombinedResult(combinedMap map[string]int) {
	for k, v := range combinedMap {
		fmt.Println(k, v)
	}
}
