// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

// ---------------------------------------------------------
// CHALLENGE
//  Add error handling to the feet to meters program.
//
// EXPECTED OUTPUT
//  go run main.go hello
//    error: 'hello' is not a number.
//
//  go run main.go what
//    error: 'what' is not a number.
//
//  go run main.go 100
//    100 feet is 30.48 meters.
// ---------------------------------------------------------

const usage = `
Feet to Meters
--------------
This program converts feet into meters.
Usage:
feet [feetsToConvert]`

/*
 * Checks number of arguments
 * Input parameters - args []string  -> slice of arguments containing feet value to be converted
 * Return values - string, error -> string to denote args[1], error to denote unwanted number of arguments
 */
func checkArguments(args []string) (string, error) {
	if len(args) < 2 {
		return "", errors.New("ERROR: Too less arguments!")
	} else if len(args) > 2 {
		return "", errors.New("ERROR: Too many arguments!")
	}
	return args[1], nil
}

/*
 * Converts Feet to meters
 * Input parameters - feetValue string  -> string value which might contain feet representation
 * Return values - string, error -> string to denote formatted output, error to denote any exception while parsing input
 */
func convertFeetToMeter(feetValue string) (string, error) {
	if feet, err := strconv.ParseFloat(feetValue, 64); err == nil {
		meters := feet * 0.3048
		conversionMsg := fmt.Sprintf("%g feet is %g meters.", feet, meters)
		return conversionMsg, nil
	}
	errString := fmt.Sprintf("ERROR: '%s' is not a number.", feetValue)
	return "", errors.New(errString)
}

func main() {
	if arg, errArgSize := checkArguments(os.Args); errArgSize != nil {
		fmt.Println(errArgSize)
	} else if meterValue, errConversion := convertFeetToMeter(arg); errConversion != nil {
		fmt.Println(errConversion)
	} else {
		fmt.Println(meterValue)
	}
}
