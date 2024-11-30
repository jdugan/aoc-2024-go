package day14

import (
	"fmt"

	// "github.com/elliotchance/pie/v2"

	"aoc/2024/pkg/reader"
)

// ========== PUBLIC FNS ==================================

func Both() {
	fmt.Println(" ")
	fmt.Println("DAY", 14)
	fmt.Println("  Puzzle 1", "=>", Puzzle1())
	fmt.Println("  Puzzle 2", "=>", Puzzle2())
	fmt.Println(" ")
}

func Puzzle1() int {
	return -1
}

func Puzzle2() int {
	return -2
}

// ========== PRIVATE FNS =================================

func data() []string {
	lines := reader.Lines("./data/day14/input.txt")

	return lines
}
