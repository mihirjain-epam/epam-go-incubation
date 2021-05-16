package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode"
)

/*
 * Input arguments -
 *                s string - camelCase input string
 *
 * Return value - int32 -> count of words
 */
func camelcase(s string) int32 {
	var count int32 = 1
	for _, r := range s {
		if unicode.IsLetter(r) && unicode.IsUpper(r) { // s[i]>=65 && s[i]<=90
			count++
		}
	}
	return count
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	s := readLine(reader)

	result := camelcase(s)

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
