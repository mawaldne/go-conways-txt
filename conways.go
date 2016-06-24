package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type Environment struct {
	cells     [][]int
	cellsNext [][]int
	rows      int
	cols      int
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
	//environment.printCells()
	environment.runRules()
	environment.copyNextGeneration()
	environment.printCells()
}

func initializeEnvironment(content string) Environment {
	lines := strings.Split(content, "\n")
	s := strings.Split(lines[0], " ")
	if len(s) != 2 {
		fmt.Fprintf(os.Stderr, "Problem with size format values")
		os.Exit(1)
	}

	rows, _ := strconv.Atoi(s[0])
	cols, _ := strconv.Atoi(s[1])
	cells := make([][]int, rows)
	cellsNext := make([][]int, rows)

	for i, line := range lines[1 : len(lines)-1] {
		cells[i] = make([]int, cols)
		cellsNext[i] = make([]int, cols)
		for j := 0; j < cols; j++ {
			char := line[j]
			if char == '.' {
				cells[i][j] = 0
				cellsNext[i][j] = 0
			}
			if char == '*' {
				cells[i][j] = 1
				cellsNext[i][j] = 1
			}
		}
	}
	environment := Environment{cells, cellsNext, rows, cols}
	return environment
}

//Conways 4 rules of life
//1. Any live cell with fewer than two live neighbours dies, as if caused by underpopulation.
//2. Any live cell with more than three live neighbours dies, as if by overcrowding.
//3. Any live cell with two or three live neighbours lives on to the next generation.
//4. Any dead cell with exactly three live neighbours becomes a live cell.
func (e *Environment) runRules() {
	for x := 0; x < e.rows; x++ {
		for y := 0; y < e.cols; y++ {
			liveCells := e.surroundingLiveCells(x, y)
			//fmt.Println("x:", x, "y:", y, "LiveCells:", liveCells)
			if e.isAlive(x, y) {
				if liveCells < 2 {
					e.cellsNext[x][y] = 0
				}
				if liveCells > 3 {
					e.cellsNext[x][y] = 0
				}
			} else {
				if liveCells == 3 {
					e.cellsNext[x][y] = 1
				}
			}
		}
	}
}

func (e *Environment) surroundingLiveCells(x int, y int) int {
	liveCells := 0
	for i := x - 1; i <= x+1; i++ {
		for j := y - 1; j <= y+1; j++ {
			if e.inEnvironment(i, j) && e.isAlive(i, j) {
				liveCells++
			}
		}
	}
	return liveCells
}

func (e *Environment) printCells() {
	for x := 0; x < e.rows; x++ {
		for y := 0; y < e.cols; y++ {
			fmt.Print(e.cells[x][y])
		}
		fmt.Println()
	}
}

func (e *Environment) copyNextGeneration() {
	for x := 0; x < e.rows; x++ {
		for y := 0; y < e.cols; y++ {
			e.cells[x][y] = e.cellsNext[x][y]
		}
	}
}

func (e *Environment) inEnvironment(x int, y int) bool {
	return x > 0 && y > 0 && x < e.rows && y < e.cols
}

func (e *Environment) isAlive(x int, y int) bool {
	return e.cells[x][y] == 1
}

func (e *Environment) isDead(x int, y int) bool {
	return e.cells[x][y] == 0
}
