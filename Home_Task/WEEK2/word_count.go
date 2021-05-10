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

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	lines_size := int32(nTemp)

	m := make(map[string]string)
	keys := make([]string, 0, lines_size)

	for line_count := 1; line_count <= int(lines_size); line_count++ {
		words := strings.Split(strings.TrimSpace(readLine(reader)), " ")
		word_count(words, m, &keys, line_count)
		fmt.Println(m)
		fmt.Println(keys)
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Fprintf(writer, "%s %s\n", k, m[k])
	}

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

func word_count(words []string, m map[string]string, keys *[]string, line_count int) {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	for _, v := range words {
		if val, ok := m[v]; ok {
			m[v] = val + " " + strconv.Itoa(line_count)
		} else {
			m[v] = strconv.Itoa(line_count)
			*keys = append(*keys, v)
		}
	}
}
