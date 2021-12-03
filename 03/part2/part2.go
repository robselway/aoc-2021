package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strings"
)

func main() {
	file, err := ioutil.ReadFile("input.txt")

	if err != nil {
		fmt.Println("Could not read input file: ", err)
		os.Exit(1)
	}

	data := strings.Split(string(file), "\n")

	oxygenMatch := findMatch(data, false, 0)[0]
	carbonMatch := findMatch(data, true, 0)[0]

	fmt.Printf("Life support rating: %d", convertBinaryToInt(oxygenMatch)*convertBinaryToInt(carbonMatch))
}

func findMatch(data []string, negate bool, index int) []string {
	remaining := []string{}

	bitCount := 0.0

	for _, number := range data {
		if number[index] == '1' {
			bitCount++
		}
	}

	shouldBePositive := bitCount >= float64(len(data))/2.0

	for i := 0; i < len(data); i++ {
		for j, bit := range data[i] {
			if j != index {
				continue
			}

			if negate {
				if !shouldBePositive && bit == '1' || shouldBePositive && bit == '0' {
					remaining = append(remaining, data[i])
				}
			} else {
				if shouldBePositive && bit == '1' || !shouldBePositive && bit == '0' {
					remaining = append(remaining, data[i])
				}
			}

		}

	}

	if len(remaining) == 1 {
		return remaining
	}

	return findMatch(remaining, negate, index+1)
}

func convertBinaryToInt(binary string) int {
	result := 0

	for i, bit := range binary {
		if bit == '1' {
			result = result + int(math.Pow(2, float64(len(binary)-i-1)))
		}
	}

	return result
}
