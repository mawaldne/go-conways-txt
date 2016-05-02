package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

//Conways 4 rules of life
//1. Any live cell with fewer than two live neighbours dies, as if caused by underpopulation.
//2. Any live cell with more than three live neighbours dies, as if by overcrowding.
//3. Any live cell with two or three live neighbours lives on to the next generation.
//4. Any dead cell with exactly three live neighbours becomes a live cell.

type Environment struct {
	cells [][]int
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Not enough command line arguments\n")
		os.Exit(1)
	}

	content, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Problem with file: %v\n", err)
		os.Exit(1)
	}

	environment := initializeEnvironment(string(content))
	fmt.Println("2d: ", environment.cells)
}

func initializeEnvironment(content string) Environment {
	lines := strings.Split(content, "\n")
	s := strings.Split(lines[0], " ")
	if len(s) != 2 {
		fmt.Fprintf(os.Stderr, "Problem with size format values")
		os.Exit(1)
	}

	x, _ := strconv.Atoi(s[0])
	y, _ := strconv.Atoi(s[1])
	cells := make([][]int, x)

	for i, line := range lines[1 : len(lines)-1] {
		cells[i] = make([]int, y)
		for j := 0; j < y; j++ {
			char := line[j]
			if char == '.' {
				cells[i][j] = 0
			}
			if char == '*' {
				cells[i][j] = 1
			}
		}
	}
	return Environment{cells}
}
