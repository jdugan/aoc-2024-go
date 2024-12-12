package day12

import (
	"aoc/2024/pkg/utility"
	"fmt"

	"github.com/elliotchance/pie/v2"
)

// ========== DEFINITION ==================================

type Garden struct {
	dims   []int
	points map[string]Point
	plots  []Plot
}

// ========== RECEIVERS ===================================

func (g *Garden) Divide() {
	visits := make([]string, 0)
	plots := make([]Plot, 0)

	// for len(visits) < len(g.points) {
	rp := g.FindReferencePoint(visits)
	name := rp.value
	ids := make([]string, 0)
	area := 0
	perimeter := 0
	fmt.Println(rp)

	plots = append(plots, Plot{name: name, ids: ids, area: area, perimeter: perimeter})
	for _, p := range plots {
		fmt.Println(p)
	}
	g.plots = plots
}

func (g Garden) FencingCost() int {
	sum := 0
	for _, p := range g.plots {
		sum += p.FencingCost()
	}
	return sum
}

func (g Garden) FindReferencePoint(visits []string) Point {
	rp := Point{}
outer:
	for y := 0; y < g.dims[1]; y++ {
		for x := 0; x < g.dims[0]; x++ {
			id := utility.CoordToId(x, y)
			if !pie.Contains(visits, id) {
				rp = g.points[id]
				break outer
			}
		}
	}
	return rp
}

// ---------- UTILITY -------------------------------------
