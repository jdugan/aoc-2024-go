package day16

import (
	"fmt"

	// "github.com/elliotchance/pie/v2"

	"aoc/2024/pkg/reader"
)

// ========== PUBLIC FNS ==================================

func Both() {
	fmt.Println(" ")
	fmt.Println("DAY", 16)
	fmt.Println("  Puzzle 1", "=>", Puzzle1())
	fmt.Println("  Puzzle 2", "=>", Puzzle2())
	fmt.Println(" ")
}

// 123417 - too low
func Puzzle1() int {
	maze := data()
	_, cost := maze.FindShortestPath()
	return cost
}

func Puzzle2() int {
	return -2
}

// ========== PRIVATE FNS =================================

func data() Maze {
	lines := reader.Lines("./data/day16/input.txt")
	points := make(map[string]Point)

	for y, line := range lines {
		for x, col := range line {
			if string(col) != "#" {
				p := Point{x: x, y: y, value: string(col)}
				points[p.Id()] = p
			}
		}
	}

	return Maze{points: points}
}
