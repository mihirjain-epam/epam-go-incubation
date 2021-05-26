package main

import "fmt"

/*
 Input parameters ->

 combinedMap map[string]int	- final result map

 Return value -> nil
*/
func printCombinedResult(combinedMap map[string]int, directory string) {
	totalFileCount := 0
	fmt.Printf("Directory: %s has below files grouped by extension\n", directory)
	for k, v := range combinedMap {
		fmt.Printf("(file extension, frequency):(%v, %v)\n", k, v)
		totalFileCount += v
	}
	fmt.Printf("Total file count:%d for folder:%s\n", totalFileCount, directory)
}
