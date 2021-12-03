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

	bitCounts := make([]int, len(data[0]))

	for i := 0; i < len(data); i++ {
		for j, bit := range data[i] {
			if bit == '1' {
				bitCounts[j]++
			}
		}
	}

	gammaRate := 0
	epsilonRate := 0

	for i, count := range bitCounts {
		additional := int(math.Pow(2, float64(len(bitCounts)-i-1)))
		if count > len(data)/2 {
			gammaRate = gammaRate + additional
		} else {
			epsilonRate = epsilonRate + additional
		}

	}

	fmt.Printf("Power consumption: %d", gammaRate*epsilonRate)
}
