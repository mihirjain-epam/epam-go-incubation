package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Input arguments -
 *                d int32 - rotate index
 * 				  arr []int32 - input slice
 *
 * Logic - 1. Take mod of 'd' with size of slice to avoid unnescessary repetitions
		   2. Use reverse utility function to rearrange the elements in the right position
 * Time complexity - O(size of arr)
 * Space complexity - O(1)
*/

func rotateLeft(d int32, arr []int32) []int32 {
	// Write your code here
	if d <= 0 {
		return arr
	}
	k := int(d) % len(arr)
	end := len(arr) - 1
	reverse(arr, 0, end)
	reverse(arr, 0, end-k)
	reverse(arr, end-k+1, end)
	return arr
}

/*
 * Input arguments -
 *                d int32 - rotate index
 * 				  arr []int32 - input slice
 *
 * Logic - 1. Take mod of 'd' with size of slice to avoid unnescessary repetitions
		   2. Use reverse utility function to rearrange the elements in the right position
 * Time complexity - O(size of arr)
 * Space complexity - O(size of arr)
*/
func rotateLeftAlt(d int32, arr []int32) []int32 {
	k := int(d) % len(arr)
	arr = append(arr[k:], arr[:k]...)
	return arr
}

/*
 * Input arguments -
 * 				  arr []int32   - input slice
 * 				  start int - starting index
 *                end int   - ending index
 * Logic - 1. Take mod of 'd' with size of slice to avoid unnescessary repetitions
		   2. Use reverse utility function to rearrange the elements in the right position
 * Time complexity - O(size of arr)
 * Space complexity - O(size of arr)
*/
func reverse(arr []int32, start, end int) {
	for start < end {
		arr[start], arr[end] = arr[end], arr[start]
		start, end = start+1, end-1
	}
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	nTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	dTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	checkError(err)
	d := int32(dTemp)

	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var arr []int32

	for i := 0; i < int(n); i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	result := rotateLeft(d, arr)

	for i, resultItem := range result {
		fmt.Fprintf(writer, "%d", resultItem)

		if i != len(result)-1 {
			fmt.Fprintf(writer, " ")
		}
	}

	fmt.Fprintf(writer, "\n")

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
