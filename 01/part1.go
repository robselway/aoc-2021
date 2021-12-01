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
		if i == 0 {
			continue
		}

		if convertToInt(measurements[i]) > convertToInt(measurements[i-1]) {
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
