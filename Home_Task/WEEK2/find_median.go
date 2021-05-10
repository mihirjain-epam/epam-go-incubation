package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
 * input args - array to find median
 * return median value
 *
 * Logic - Sort array using builtin sort by passing a anonymous function as less function,
 *		   median is the middle element after sorting
 */

func findMedian(arr []int32) int32 {
	sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })
	return arr[int(len(arr)/2)]
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var arr []int32

	for i := 0; i < int(n); i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	result := findMedian(arr)

	fmt.Fprintf(writer, "%d\n", result)

	writer.Flush()
}

/*
 * generated method @hackerrank
 */
func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

/*
 * generated method @hackerrank
 */
func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
