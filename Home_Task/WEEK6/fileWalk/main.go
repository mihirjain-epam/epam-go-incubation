package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

var dirPath string

const command string = `
Enter the directory path:
`

func main() {
	fileCh := make(chan map[string]int, 1)
	wg := &sync.WaitGroup{}
	wgRecursion := &sync.WaitGroup{}
	combinedMap := make(map[string]int)
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(command)
	scanner.Scan()
	dirPath = strings.Trim(scanner.Text(), " ")
	if len(dirPath) < 1 {
		fmt.Println("ERROR: no directory input found!")
		return
	}
	_, err := os.Stat(dirPath)
	if err != nil {
		fmt.Println("ERROR: no directory/file found at - ", dirPath)
		return
	}
	wg.Add(1)
	go fileConcurrentWalk(wg, wgRecursion, fileCh, dirPath, true)
	wg.Add(1)
	go mergeResults(wg, fileCh, combinedMap)
	wg.Wait()
	printCombinedResult(combinedMap)
}

/*
 fileConcurrentWalk function recursively and concurrently walks over the directories
 and puts all the file extensions with their frequencies in a map. This map is put into a
 recieve-only channel.

 Input parameters ->

 wg *sync.WaitGroup	- wait group for first time call of function

 wgRecursion *sync.WaitGroup - wait group for every non-first call of function

 fileCh chan<- map[string]int - recieve only channel

 inputDirPath string - directory path

 firstTime bool	- flag to identify first time call of function

 Return value -> nil
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
 mergeResults function merges the maps from all sub-directories into a combinedMap concurrently

 Input parameters ->

 wg *sync.WaitGroup	- wait group for first time call of function

 fileCh chan<- map[string]int - recieve only channel

 combinedMap map[string]int	- final result map

 Return value -> nil
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
 Input parameters ->

 combinedMap map[string]int	- final result map

 Return value -> nil
*/
func printCombinedResult(combinedMap map[string]int) {
	totalFileCount := 0
	for k, v := range combinedMap {
		fmt.Printf("(file extension, frequency):(%v, %v)\n", k, v)
		totalFileCount += v
	}
	fmt.Println("Total file count:", totalFileCount)
}
