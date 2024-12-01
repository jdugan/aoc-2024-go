package utility

import (
	"github.com/elliotchance/pie/v2"
)

// ========== MATH HELPERS ================================

func Distance(a int, b int) int {
	return pie.Abs(a - b)
}
