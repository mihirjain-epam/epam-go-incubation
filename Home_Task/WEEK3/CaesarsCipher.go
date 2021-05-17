package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"unicode"
)

const (
	numberOfAlphabets              int32 = 26
	lowerCaseAsciiAlphabetMinValue int32 = 97
	lowerCaseAsciiAlphabetMaxValue int32 = 122
	upperCaseAsciiAlphabetMinValue int32 = 65
	upperCaseAsciiAlphabetMaxValue int32 = 90
)

/*
 * Input arguments -
 *                s string - plain text
 * 				  k int32 - number of times to left rotate english alphabet
 *
 * Return value -
 * 				  string - cipher text/message
 */

func caesarCipher(s string, k int32) string {
	k = k % numberOfAlphabets
	cipheredMessage := make([]rune, 0)
	for _, r := range s {
		cipheredLetter := r
		if unicode.IsLetter(r) {
			cipheredLetter = r + k
			if unicode.IsUpper(r) && cipheredLetter > upperCaseAsciiAlphabetMaxValue {
				cipheredLetter = upperCaseAsciiAlphabetMinValue + cipheredLetter - upperCaseAsciiAlphabetMaxValue - 1
			} else if unicode.IsLower(r) && cipheredLetter > lowerCaseAsciiAlphabetMaxValue {
				cipheredLetter = lowerCaseAsciiAlphabetMinValue + cipheredLetter - lowerCaseAsciiAlphabetMaxValue - 1
			}
		}
		cipheredMessage = append(cipheredMessage, cipheredLetter)
	}
	return string(cipheredMessage)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	_ = int32(nTemp)

	s := readLine(reader)

	kTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	k := int32(kTemp)

	result := caesarCipher(s, k)

	fmt.Fprintf(writer, "%s\n", result)

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
