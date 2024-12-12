package day12

// ========== DEFINITION ==================================

type Plot struct {
	name      string
	ids       []string
	area      int
	perimeter int
}

// ========== RECEIVERS ===================================

func (p Plot) FencingCost() int {
	return p.area * p.perimeter
}

// ---------- UTILITY -------------------------------------
