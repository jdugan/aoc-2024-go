package day17

import (
	"aoc/2024/pkg/reader"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/elliotchance/pie/v2"
)

// ========== PUBLIC FNS ==================================

func Both() {
	fmt.Println(" ")
	fmt.Println("DAY", 17)
	fmt.Println("  Puzzle 1", "=>", Puzzle1())
	fmt.Println("  Puzzle 2", "=>", Puzzle2())
	fmt.Println(" ")
}

func Puzzle1() string {
	memory, program := data()
	computer := Computer{memory: memory}
	result := computer.Run(program)
	return result
}

// 185449480 - too low
func Puzzle2() int {
	memory, program := data()
	computer := Computer{memory: memory}
	result := computer.Search(program)
	return result
}

// ========== PRIVATE FNS =================================

func data() (map[string]int, []int) {
	lines := reader.Lines("./data/day17/input-test2.txt")

	// parse registers
	re1 := regexp.MustCompile("^Register ([A-Z]): ([0-9]+)$")
	memory := make(map[string]int)
	for _, line := range lines[:3] {
		matches := re1.FindAllStringSubmatch(line, 1)
		match := matches[0]
		key := match[1]
		value, _ := strconv.Atoi(match[2])
		memory[key] = value
	}

	// parse program
	re2 := regexp.MustCompile("^Program: (.+)$")
	matches := re2.FindAllStringSubmatch(lines[4], 1)
	match := matches[0]
	strs := strings.Split(match[1], ",")
	program := pie.Ints(strs)

	return memory, program
}
