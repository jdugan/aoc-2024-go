package day25

import (
	"fmt"

	// "github.com/elliotchance/pie/v2"

	"aoc/2024/pkg/reader"
)

// ========== PUBLIC FNS ==================================

func Both() {
	fmt.Println(" ")
	fmt.Println("DAY", 25)
	fmt.Println("  Puzzle 1", "=>", Puzzle1())
	fmt.Println("  Puzzle 2", "=>", Puzzle2())
	fmt.Println(" ")
}

func Puzzle1() int {
	return -1
}

func Puzzle2() int {
	return 50
}

// ========== PRIVATE FNS =================================

func data() []string {
	lines := reader.Lines("./data/day25/input.txt")

	return lines
}
