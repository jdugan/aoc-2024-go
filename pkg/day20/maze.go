package day20

import (
	"aoc/2024/pkg/utility"
	"fmt"

	"github.com/RyanCarrier/dijkstra/v2"
	"github.com/elliotchance/pie/v2"
)

// ========== DEFINITION ==================================

type Maze struct {
	points map[string]Point
}

// ========== RECEIVERS ===================================

func (m Maze) ShortcutCount(minimum int) int {
	shortcuts := m.FindShortcuts()
	count := 0
	for _, sc := range shortcuts {
		if sc.distance >= minimum {
			count += 1
		}
	}
	return count
}

func (m Maze) FindShortcuts() []Shortcut {
	shortcuts := make([]Shortcut, 0)
	path, _ := m.ShortestPath()
	head, tail := pie.Shift(path)
	for len(head) > 0 {
		hp := m.points[head]
		for idx, tid := range tail {
			tp := m.points[tid]
			if idx > 1 && hp.DistanceTo(tp) == 2 {
				sc := Shortcut{p1: hp, p2: tp, distance: idx - 1}
				shortcuts = append(shortcuts, sc)
			}
		}
		head, tail = pie.Shift(tail)
	}
	return shortcuts
}

// ---------- UTILITIES -----------------------------------

func (m Maze) ShortestPath() ([]string, int) {
	// add nodes and map ids to indices
	g := dijkstra.NewGraph()
	vmap := make(map[string]int)
	ivmap := make(map[int]string)
	for pid, _ := range m.points {
		idx := g.AddNewEmptyVertex()
		vmap[pid] = idx
		ivmap[idx] = pid
	}
	// add edges
	for pid, p := range m.points {
		fnode := vmap[pid]
		for _, aid := range p.AdjacentIds() {
			_, ok := m.points[aid]
			if ok {
				tnode := vmap[aid]
				g.AddArc(fnode, tnode, uint64(1))
			}
		}
	}
	// locate path
	onode := vmap[m.OriginId()]
	tnode := vmap[m.TerminusId()]
	best, _ := g.Shortest(onode, tnode)
	path := make([]string, 0)
	for _, idx := range best.Path {
		path = append(path, ivmap[idx])
	}
	distance := int(best.Distance)

	return path, distance
}

func (m Maze) OriginId() string {
	oid := ""
	for id, p := range m.points {
		if p.IsOrigin() {
			oid = id
		}
	}
	return oid
}

func (m Maze) TerminusId() string {
	tid := ""
	for id, p := range m.points {
		if p.IsTerminus() {
			tid = id
		}
	}
	return tid
}

// ---------- VISUALS -------------------------------------

func (m Maze) Dimensions() (int, int, int, int) {
	min_x := 999999999
	max_x := 0
	min_y := 999999999
	max_y := 0
	for _, p := range m.points {
		switch {
		case p.x < min_x:
			min_x = p.x
		case p.x > max_x:
			max_x = p.x
		case p.y < min_y:
			min_y = p.y
		case p.y > max_y:
			max_y = p.y
		}
	}
	return min_x, max_x, min_y, max_y
}

func (m Maze) Print() {
	min_x, max_x, min_y, max_y := m.Dimensions()

	for y := min_y - 1; y <= max_y+1; y++ {
		row := ""
		for x := min_x - 1; x <= max_x+1; x++ {
			p, ok := m.points[utility.CoordToId(x, y)]
			if ok {
				row += p.value
			} else {
				row += "#"
			}
		}
		fmt.Println(row)
	}
	fmt.Println("")
}
