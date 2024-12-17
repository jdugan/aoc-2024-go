package day17

import (
	"fmt"
	"math"
	"strings"

	"github.com/elliotchance/pie/v2"
)

// ========== DEFINITION ==================================

type Computer struct {
	memory  map[string]int
	output  []int
	pointer int
}

// ========== RECEIVERS ===================================

func (c *Computer) Run(program []int) string {
	c.RunWithTarget(program, make([]int, 0))
	return c.Result()
}

func (c *Computer) RunWithTarget(program []int, target []int) string {
	seed := c.memory["A"]
loop:
	for c.pointer < len(program) {
		opscode := program[c.pointer]
		operand := program[c.pointer+1]
		c.pointer += 2

		switch opscode {
		// adv
		case 0:
			combo, _ := c.ComboValue(operand)
			denom := int(math.Pow(2, float64(combo)))
			c.memory["A"] = c.memory["A"] / denom
		//bxl
		case 1:
			c.memory["B"] = c.memory["B"] ^ operand
		// bst
		case 2:
			combo, _ := c.ComboValue(operand)
			c.memory["B"] = combo % 8
		// jnz
		case 3:
			if c.memory["A"] != 0 {
				c.pointer = operand
			}
		// bxc
		case 4:
			c.memory["B"] = c.memory["B"] ^ c.memory["C"]
		// out
		case 5:
			combo, _ := c.ComboValue(operand)
			c.output = append(c.output, combo%8)
			if len(target) > 0 {
				olen := len(c.output)
				ostr := c.Result()
				tstr := strings.Join(pie.Strings(target[:olen]), ",")
				if ostr != tstr {
					break loop
				} else {
					if len(c.output) >= 10 {
						fmt.Println(seed, c.output)
					}
				}
			}
		// bdv
		case 6:
			combo, _ := c.ComboValue(operand)
			denom := int(math.Pow(2, float64(combo)))
			c.memory["B"] = c.memory["A"] / denom
		// cdv
		case 7:
			combo, _ := c.ComboValue(operand)
			denom := int(math.Pow(2, float64(combo)))
			c.memory["C"] = c.memory["A"] / denom
		}
	}
	return c.Result()
}

//	1 - 6
//	2 - 14
//	3 - 332
//	4 - 31157
//	5 - 1079733
//	6 - 2977469
//	7 - 7171773
//	8 - 23948989
//	9 - 23948989
//
// 10 -
// 11 -
// 12 -
// 13 -
// 14 -
// 15 -
// 16 -
func (c *Computer) Search(program []int) int {
	target := strings.Join(pie.Strings(program), ",")
	result := ""
	n := 0
	for result != target {
		n += 1
		c.Reset(n)
		result = c.RunWithTarget(program, program)
	}
	return n
}

// ---------- UTILITIES -----------------------------------

func (c Computer) ComboValue(operand int) (int, bool) {
	switch operand {
	case 0, 1, 2, 3:
		return operand, true
	case 4:
		return c.memory["A"], true
	case 5:
		return c.memory["B"], true
	case 6:
		return c.memory["C"], true
	default:
		return -1, false
	}
}

func (c *Computer) Reset(seed int) {
	c.memory["A"] = seed
	c.memory["B"] = 0
	c.memory["C"] = 0
	c.pointer = 0
	c.output = make([]int, 0)
}

func (c Computer) Result() string {
	return strings.Join(pie.Strings(c.output), ",")
}
