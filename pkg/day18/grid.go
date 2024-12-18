package day18

import (
	"aoc/2024/pkg/utility"
	"fmt"

	"github.com/RyanCarrier/dijkstra/v2"
)

// ========== DEFINITION ==================================

type Grid struct {
	max    int
	bytes  []string
	points map[string]Point
}

// ========== RECEIVERS ===================================

func (g *Grid) DropFirst(size int) {
	for _, id := range g.bytes[:size] {
		p := g.points[id]
		p.value = "#"
		g.points[id] = p
	}
}

func (g Grid) FindShortestPath() int {

	// add nodes and map ids to indices
	graph := dijkstra.NewGraph()
	vmap := make(map[string]int)
	ivmap := make(map[int]string)
	for pid, _ := range g.points {
		idx := graph.AddNewEmptyVertex()
		vmap[pid] = idx
		ivmap[idx] = pid
	}
	// add edges using map to translate ids to indices
	for pid, p := range g.points {
		for _, aid := range p.AdjacentIds() {
			ap, ok := g.points[aid]
			if ok && !ap.IsCorrupted() {
				fnode := vmap[pid]
				tnode := vmap[aid]
				graph.AddArc(fnode, tnode, uint64(1))
			}
		}
	}

	oid := g.OriginId()
	tid := g.TerminusId()
	onode := vmap[oid]
	tnode := vmap[tid]
	best, _ := graph.Shortest(onode, tnode)
	dist := int(best.Distance)

	return dist
}

func (g Grid) FindTerminalByte() string {
	return "0,0"
}

// ---------- UTILITIES -----------------------------------

func (g Grid) OriginId() string {
	return utility.CoordToId(0, 0)
}

func (g Grid) TerminusId() string {
	return utility.CoordToId(g.max, g.max)
}

// ---------- VISUALS -------------------------------------

func (g Grid) Print() {
	for y := 0; y <= g.max; y++ {
		row := ""
		for x := 0; x <= g.max; x++ {
			id := utility.CoordToId(x, y)
			row += g.points[id].value
		}
		fmt.Println(row)
	}
	fmt.Println("")
	fmt.Println("")
}
