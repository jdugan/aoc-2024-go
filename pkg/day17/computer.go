package day17

import "math"

// ========== DEFINITION ==================================

type Computer struct {
	memory  map[string]int
	output  []int
	pointer int
}

// ========== RECEIVERS ===================================

func (c *Computer) Run(program []int) {
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
