package main

import (
	"io/ioutil"
	"log"
	"path/filepath"
	"sync"
)

/*
 concurrentFileWalkWaitHandler function closes the fileCh once all concurrentFileWalk
 goroutines are completed

 Input parameters ->

 fileCh chan<- map[string]int - recieve only channel

 inputDirPath string - directory path

 Return value -> nil
*/
func concurrentFileWalkWaitHandler(fileCh chan<- map[string]int, inputDirPath string) {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go concurrentFileWalk(wg, fileCh, inputDirPath)
	wg.Wait()
	close(fileCh)
}

/*
 concurrentFileWalk function recursively and concurrently walks over the directories
 and puts all the file extensions with their frequencies in a map. This map is put into a
 recieve-only channel.

 Input parameters ->

 wg *sync.WaitGroup	- wait group for first time call of function

 fileCh chan<- map[string]int - recieve only channel

 inputDirPath string - directory path

 Return value -> nil
*/
func concurrentFileWalk(wg *sync.WaitGroup, fileCh chan<- map[string]int, inputDirPath string) {
	defer wg.Done()

	extMap := make(map[string]int)
	files, err := ioutil.ReadDir(inputDirPath)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		if file.IsDir() {
			wg.Add(1)
			go concurrentFileWalk(wg, fileCh, inputDirPath+"/"+file.Name())
		} else {
			fileExt := filepath.Ext(file.Name())
			if fileExt == "" {
				fileExt = "hidden"
			}
			extMap[fileExt] += 1
		}

	}
	fileCh <- extMap
}
