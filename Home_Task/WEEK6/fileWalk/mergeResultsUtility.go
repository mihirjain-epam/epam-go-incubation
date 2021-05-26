package main

/*
 mergeResults function merges the maps from all sub-directories into a combinedMap concurrently

 Input parameters ->

 fileCh chan<- map[string]int - recieve only channel

 combinedMap map[string]int	- final result map

 Return value -> nil
*/
func mergeResults(fileCh <-chan map[string]int, combinedMap *map[string]int) {
	for extMap := range fileCh {
		for k, v := range extMap {
			(*combinedMap)[k] += v
		}
	}
}
