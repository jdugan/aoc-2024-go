package day21

import (
	"aoc/2024/pkg/utility"
	"strings"

	"github.com/RyanCarrier/dijkstra/v2"
	"github.com/elliotchance/pie/v2"
)

// ========== DEFINITION ==================================

type ControlBot struct {
	points   map[string]Point
	segments map[string][][]string
}

// ========== RECEIVERS ===================================

func (b ControlBot) Initialize() ControlBot {
	// build structs
	points := make(map[string]Point)
	graph := dijkstra.NewGraph()
	vmap := make(map[string]int)
	ivmap := make(map[int]string)
	segments := make(map[string][][]string)

	// add points
	points["1,0"] = Point{x: 1, y: 0, value: "^"}
	points["2,0"] = Point{x: 2, y: 0, value: "A"}
	points["0,1"] = Point{x: 0, y: 1, value: "<"}
	points["1,1"] = Point{x: 1, y: 1, value: "v"}
	points["2,1"] = Point{x: 2, y: 1, value: ">"}

	// add vertices and map nodes
	for _, p := range points {
		idx := graph.AddNewEmptyVertex()
		vmap[p.value] = idx
		ivmap[idx] = p.value
	}

	// add arcs
	for _, p := range points {
		fnode := vmap[p.value]
		for _, aid := range p.AdjacentIds() {
			ap, ok := points[aid]
			if ok {
				tnode := vmap[ap.value]
				graph.AddArc(fnode, tnode, 1)
			}
		}
	}

	// map all possible segments to save time
	pvmap := b.PointValueMap(points)
	for _, p1 := range points {
		for _, p2 := range points {
			sid := p1.value + ";" + p2.value
			tpaths := make([][]string, 0)
			if p1.value != p2.value {
				fnode := vmap[p1.value]
				tnode := vmap[p2.value]
				best, _ := graph.ShortestAll(fnode, tnode)
				for _, vids := range best.Paths {
					path := make([]string, 0)
					for _, vid := range vids {
						path = append(path, ivmap[vid])
					}
					path = b.TranslateInstructions(pvmap, path)
					tpaths = append(tpaths, path)
				}
			}
			segments[sid] = tpaths
		}
	}

	// return robot
	robot := ControlBot{points: points, segments: segments}
	return robot
}

func (b ControlBot) InstructionsForCodes(codes []string) []string {
	new_codes := make([]string, 0)
	for _, code := range codes {
		ncs := b.InstructionsForCode(code)
		new_codes = append(new_codes, ncs...)
	}
	// min := pie.Min(pie.Map(new_codes, func(c string) int { return len(c) }))
	// new_codes = pie.Filter(new_codes, func(c string) bool { return len(c) == min })
	return pie.Unique(new_codes)
}

func (b ControlBot) InstructionsForCode(code string) []string {
	// set up
	possibles := make([][]string, 0)
	possibles = append(possibles, []string{"A"})
	code = "A" + code
	digits := strings.Split(code, "")
	head, tail := pie.Shift(digits)
	// loop digits in code
	for len(tail) > 0 {
		sid := head + ";" + tail[0]
		segments := b.segments[sid]
		new_possibles := make([][]string, 0)
		for _, possible := range possibles {
			if len(segments) > 0 {
				for _, segment := range segments {
					np := append(possible, segment...)
					np = append(np, "A")
					new_possibles = append(new_possibles, np)
				}
			} else {
				np := append(possible, "A")
				new_possibles = append(new_possibles, np)
			}
		}
		possibles = new_possibles
		head, tail = pie.Shift(tail)
	}
	codes := make([]string, 0)
	for _, possible := range possibles {
		code := strings.Join(possible, "")
		codes = append(codes, code[1:])
	}
	// min := pie.Min(pie.Map(codes, func(c string) int { return len(c) }))
	// codes = pie.Filter(codes, func(c string) bool { return len(c) == min })
	return pie.Unique(codes)
}

func (b ControlBot) ShortestCodes(codes []string) ([]string, int) {
	lengths := pie.Map(codes, func(c string) int { return len(c) })
	minimum := pie.Min(lengths)
	shorties := pie.Filter(codes, func(c string) bool { return len(c) == minimum })
	return shorties, minimum
}

// ---------- UTILITIES -----------------------------------

func (b ControlBot) PointValueMap(points map[string]Point) map[string]Point {
	pmap := make(map[string]Point)
	for _, p := range points {
		pmap[p.value] = p
	}
	return pmap
}
func (b ControlBot) TranslateInstructions(pvmap map[string]Point, path []string) []string {
	codes := make([]string, 0)
	head, tail := pie.Shift(path)
	for len(tail) > 0 {
		fp := pvmap[head]
		tp := pvmap[tail[0]]
		dir := b.TranslatePointsToDirection(fp, tp)
		codes = append(codes, dir)
		head, tail = pie.Shift(tail)
	}
	return codes
}

func (b ControlBot) TranslatePointsToDirection(fp Point, tp Point) string {
	fx, fy := utility.CoordFromId(fp.Id())
	tx, ty := utility.CoordFromId(tp.Id())
	dx := tx - fx
	dy := ty - fy
	dir := "X"
	switch {
	case dx == 1 && dy == 0:
		dir = ">"
	case dx == 0 && dy == -1:
		dir = "^"
	case dx == 0 && dy == 1:
		dir = "v"
	case dx == -1 && dy == 0:
		dir = "<"
	}
	return dir
}
