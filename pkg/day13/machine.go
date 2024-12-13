package day13

// ========== DEFINITION ==================================

type Machine struct {
	ax int
	ay int
	bx int
	by int
	px int
	py int
}

// ========== RECEIVERS ===================================

func (m Machine) FewestCoins(fa int, fb int) int {
	combos := m.WinningCombos()
	best := 0
	for _, combo := range combos {
		coins := fa*combo[0] + fb*combo[1]
		if best == 0 || coins < best {
			best = coins
		}
	}
	return best
}

// ---------- UTILITIES -----------------------------------

func (m Machine) WinningCombos() [][]int {
	combos := make([][]int, 0)
	max := m.px / m.ax
	for num_a := 0; num_a <= max; num_a++ {
		rx := m.px - (m.ax * num_a)
		if rx%m.bx == 0 {
			num_b := rx / m.bx
			if m.py == m.ay*num_a+m.by*num_b {
				combos = append(combos, []int{num_a, num_b})
			}
		}
	}
	return combos
}
