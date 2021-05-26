package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

var dirPath string

const command string = `
Enter the directory path:
`

func getUserInput() (string, error) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(command)
	scanner.Scan()
	dirPath = strings.Trim(scanner.Text(), " ")
	if len(dirPath) < 1 {
		fmt.Println()
		return "", errors.New("ERROR: no directory input found")
	}
	_, err := os.Stat(dirPath)
	if err != nil {
		msg := fmt.Sprintf("ERROR: no directory/file found at - %s", dirPath)
		return "", errors.New(msg)
	}
	return dirPath, nil
}
