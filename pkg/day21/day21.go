package day21

import (
	"aoc/2024/pkg/reader"
	"fmt"
)

// ========== PUBLIC FNS ==================================

func Both() {
	fmt.Println(" ")
	fmt.Println("DAY", 21)
	fmt.Println("  Puzzle 1", "=>", Puzzle1())
	fmt.Println("  Puzzle 2", "=>", Puzzle2())
	fmt.Println(" ")
}

// 126886 - too high
// 127936
func Puzzle1() int {
	cmds := data()
	sum := 0
	b1 := DoorBot{}.Initialize()
	b2 := ControlBot{}.Initialize()
	b3 := ControlBot{}.Initialize()
	for _, cmd := range cmds {
		// fmt.Println("---------------------")
		// fmt.Println(cmd.code)
		codes := b1.InstructionsForCode(cmd.code)
		// for _, code := range codes {
		// 	fmt.Println(" ", code)
		// }
		codes = b2.InstructionsForCodes(codes)
		// for _, code := range codes {
		// 	fmt.Println("  ", code)
		// }
		codes = b3.InstructionsForCodes(codes)
		// for _, code := range codes {
		// 	fmt.Println("   ", code)
		// }
		codes, length := b3.ShortestCodes(codes)
		// for _, code := range codes {
		// 	fmt.Println("   ", code)
		// }
		fmt.Println(cmd.code, cmd.Multiplier(), length)
		sum += cmd.Multiplier() * length
	}
	return sum
}

// 72
// 68
// 72
// 70
// 72
func Puzzle2() int {
	return -2
}

// ========== PRIVATE FNS =================================

func data() []Command {
	lines := reader.Lines("./data/day21/input.txt")
	cmds := make([]Command, 0)
	for _, line := range lines {
		cmds = append(cmds, Command{code: line})
	}
	return cmds
}
