package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := ioutil.ReadFile("input.txt")

	if err != nil {
		fmt.Println("Could not read input file: ", err)
		os.Exit(1)
	}

	measurements := strings.Split(string(file), "\n")

	increases := 0

	for i := 0; i < len(measurements); i++ {
		// Ignore first 3 items
		if i < 3 {
			continue
		}

		previousWindow := convertToInt(measurements[i-3]) + convertToInt(measurements[i-2]) + convertToInt(measurements[i-1])

		currentWindow := convertToInt(measurements[i-2]) + convertToInt(measurements[i-1]) + convertToInt(measurements[i])

		if currentWindow > previousWindow {
			increases++
		}
	}

	fmt.Printf("Increases: %d", increases)
}

func convertToInt(s string) int {
	i, err := strconv.Atoi(s)

	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	return i
}
