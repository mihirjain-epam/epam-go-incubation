package main

import (
	"fmt"
	"time"
)

func main() {
	dirPath, err := getUserInput()
	if err != nil {
		fmt.Print(err.Error())
		return
	}
	start := time.Now()
	fileCh := make(chan map[string]int, 1)
	combinedMap := make(map[string]int)
	go concurrentFileWalkWaitHandler(fileCh, dirPath)
	mergeResults(fileCh, &combinedMap)
	printCombinedResult(combinedMap, dirPath)
	duration := time.Since(start)
	fmt.Println(duration)
}
