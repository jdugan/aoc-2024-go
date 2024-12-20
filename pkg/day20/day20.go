package day20

import (
	"aoc/2024/pkg/reader"
	"fmt"
	"sort"
	"strconv"
)

// ========== PUBLIC FNS ==================================

func Both() {
	fmt.Println(" ")
	fmt.Println("DAY", 20)
	fmt.Println("  Puzzle 1", "=>", Puzzle1())
	fmt.Println("  Puzzle 2", "=>", Puzzle2())
	fmt.Println(" ")
}

func Puzzle1() int {
	maze := data()
	shortcuts := maze.FindShortcuts(2, 2)
	print(shortcuts)
	return len(shortcuts)
}

// 1043543 - too high
// 554470 - too low
func Puzzle2() int {
	maze := data()
	shortcuts := maze.FindShortcuts(20, 50)
	print(shortcuts)
	return len(shortcuts)
}

// ========== PRIVATE FNS =================================

func data() Maze {
	lines := reader.Lines("./data/day20/input-test.txt")
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

func print(shortcuts []Shortcut) {
	dmap := make(map[int]int)
	for _, sc := range shortcuts {
		_, ok := dmap[sc.steps]
		if !ok {
			dmap[sc.steps] = 0
		}
		dmap[sc.steps] += 1
	}
	nums := make([]int, 0)
	for dist, count := range dmap {
		nums = append(nums, dist*1000000+count)
	}
	sort.Ints(nums)
	for _, num := range nums {
		dist := num / 1000000
		count := num % 1000000
		sdist := strconv.Itoa(dist)
		scount := strconv.Itoa(count)
		fmt.Println("There are " + scount + " cheats that save " + sdist + " picoseconds.")
	}
}
