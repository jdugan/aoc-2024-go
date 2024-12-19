package day19

import (
	"aoc/2024/pkg/reader"
	"fmt"
	"strings"
)

// ========== PUBLIC FNS ==================================

func Both() {
	fmt.Println(" ")
	fmt.Println("DAY", 19)
	fmt.Println("  Puzzle 1", "=>", Puzzle1())
	fmt.Println("  Puzzle 2", "=>", Puzzle2())
	fmt.Println(" ")
}

func Puzzle1() int {
	towels, designs := data()
	sum := 0
	for i, design := range designs {
		fmt.Println(i)
		seq := design.FindSequences(towels)
		if len(seq) > 0 {
			sum += 1
		}
	}
	return sum
}

func Puzzle2() int {
	return -2
}

// ========== PRIVATE FNS =================================

func data() ([]string, []Design) {
	lines := reader.Lines("./data/day19/input.txt")
	towels := strings.Split(lines[0], ", ")
	designs := make([]Design, 0)
	for _, pattern := range lines[2:] {
		designs = append(designs, Design{pattern: pattern})
	}
	return towels, designs
}
