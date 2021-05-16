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
 *                s string - input string
 *
 * Return value - int32 -> distinct chars
 *
 * Logic - If we only count the unique chars with cost 1 each, all their repetitions can be appended as free operations
 */

func stringConstruction(s string) int32 {
	distinctChars := make(map[byte]int32, 0)
	for i := range s {
		distinctChars[s[i]] = 1
	}
	return int32(len(distinctChars))
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	qTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	q := int32(qTemp)

	for qItr := 0; qItr < int(q); qItr++ {
		s := readLine(reader)

		result := stringConstruction(s)

		fmt.Fprintf(writer, "%d\n", result)
	}

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
