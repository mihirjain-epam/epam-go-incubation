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
 *                s string - input string
 *
 * Return value - string -> "pangram" or "not pangram"
 */

func pangrams(s string) string {
	occuredLetters := make(map[rune]int)
	for _, c := range s {
		if unicode.IsLetter(c) {
			if unicode.IsLower(c) {
				c -= 'a' - 'A'
			}
			occuredLetters[c] = 1
		}
	}
	if len(occuredLetters) < 26 {
		return "not pangram"
	}
	return "pangram"
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	s := readLine(reader)

	result := pangrams(s)

	fmt.Fprintf(writer, "%s\n", result)

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
