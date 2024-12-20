package day20

import (
	"aoc/2024/pkg/reader"
	"fmt"
	"strconv"
	"strings"
)

// ========== PUBLIC FNS ==================================

func Both() {
	fmt.Println(" ")
	fmt.Println("DAY", 20)
	fmt.Println("  Puzzle 1", "=>", Puzzle1())
	fmt.Println("  Puzzle 2", "=>", Puzzle2())
	fmt.Println(" ")
}

// 6949 - too high
func Puzzle1() int {
	maze := data()
	return maze.ShortcutCount()
}

func Puzzle2() int {
	return -2
}

// ========== PRIVATE FNS =================================

func data() Maze {
	lines := reader.Lines("./data/day20/input.txt")
	smin := strings.Replace(lines[0], "Minimum: ", "", 1)
	min, _ := strconv.Atoi(smin)

	points := make(map[string]Point)
	for y, line := range lines[2:] {
		for x, col := range line {
			if string(col) != "#" {
				p := Point{x: x, y: y, value: string(col)}
				points[p.Id()] = p
			}
		}
	}

	return Maze{points: points, minimum: min}
}
