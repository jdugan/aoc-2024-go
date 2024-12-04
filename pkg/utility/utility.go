package utility

import (
	"strconv"

	"github.com/elliotchance/pie/v2"
)

// ========== MATH HELPERS ================================

func CoordToId(x int, y int) string {
	sx := strconv.Itoa(x)
	sy := strconv.Itoa(y)
	return sx + "," + sy
}

func Distance(a int, b int) int {
	return pie.Abs(a - b)
}
