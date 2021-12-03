package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type Position struct {
	Horizontal int
	Depth      int
	Aim        int
}

func (p *Position) MoveForward(distance int) {
	p.Horizontal = p.Horizontal + distance
	p.Depth = p.Depth + (p.Aim * distance)
}

func (p *Position) MoveDown(distance int) {
	p.Aim = p.Aim + distance
}

func (p *Position) MoveUp(distance int) {
	p.Aim = p.Aim - distance
}

const ForwardPrefix = "forward "
const DownPrefix = "down "
const UpPrefix = "up "

func main() {
	file, err := ioutil.ReadFile("input.txt")

	if err != nil {
		fmt.Println("Could not read input file: ", err)
		os.Exit(1)
	}

	directions := strings.Split(string(file), "\n")

	position := Position{}

	for i := 0; i < len(directions); i++ {
		if strings.HasPrefix(directions[i], ForwardPrefix) {
			distance := strings.TrimPrefix(directions[i], ForwardPrefix)
			position.MoveForward(convertToInt(distance))
			continue
		}

		if strings.HasPrefix(directions[i], DownPrefix) {
			distance := strings.TrimPrefix(directions[i], DownPrefix)
			position.MoveDown(convertToInt(distance))
			continue
		}

		if strings.HasPrefix(directions[i], UpPrefix) {
			distance := strings.TrimPrefix(directions[i], UpPrefix)
			position.MoveUp(convertToInt(distance))
			continue
		}
	}

	fmt.Printf("Position: %d", position.Horizontal*position.Depth)
}

func convertToInt(s string) int {
	i, err := strconv.Atoi(s)

	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	return i
}
