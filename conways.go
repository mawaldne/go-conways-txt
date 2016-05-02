package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

//Conways 4 rules of life
//1. Any live cell with fewer than two live neighbours dies, as if caused by underpopulation.
//2. Any live cell with more than three live neighbours dies, as if by overcrowding.
//3. Any live cell with two or three live neighbours lives on to the next generation.
//4. Any dead cell with exactly three live neighbours becomes a live cell.

type Environment struct {
	cells [][]int
	x     int
	y     int
}

func main() {
	if len(os.Args) != 4 {
		fmt.Fprintf(os.Stderr, "Not enough command line arguments\n")
		os.Exit(1)
	}

	//x := os.Args[1]
	//y := os.Args[2]
	content, err := ioutil.ReadFile(os.Args[3])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Problem with file: %v\n", err)
		os.Exit(1)
	}
	//lines := strings.Split(string(content), "\n")
	fmt.Println(string(content))
}

func (e *Environment) initializeCells(content) float64 {
}
