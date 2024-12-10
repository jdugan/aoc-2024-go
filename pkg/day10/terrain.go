package day10

import (
	"strings"

	"github.com/elliotchance/pie/v2"
)

// ========== DEFINITION ==================================

type Terrain struct {
	points map[string]Point
}

// ========== RECEIVERS ===================================

func (t Terrain) AdjacentPoints(p Point) []Point {
	aps := make([]Point, 0)
	for _, aid := range p.AdjacentIds() {
		p, ok := t.points[aid]
		if ok {
			aps = append(aps, p)
		}
	}
	return aps
}

func (t Terrain) GetTrailHeads() []Point {
	theads := make([]Point, 0)
	for _, p := range t.points {
		if p.value == 0 {
			theads = append(theads, p)
		}
	}
	return theads
}

func (t Terrain) GetTrailSystem() map[string][][]Point {
	tmap := make(map[string][][]Point)
	for _, thead := range t.GetTrailHeads() {
		trails := make([][]Point, 0)
		possibles := [][]Point{[]Point{thead}}
		for len(possibles) > 0 {
			new_possibles := make([][]Point, 0)
			for _, possible := range possibles {
				bp := pie.Last(possible)
				aps := t.AdjacentPoints(bp)
				for _, ap := range aps {
					if ap.value == bp.value+1 {
						cp := make([]Point, len(possible))
						copy(cp, possible)
						cp = append(cp, ap)
						if ap.value == 9 {
							trails = append(trails, cp)
						} else {
							new_possibles = append(new_possibles, cp)
						}
					}
				}

			}
			possibles = new_possibles
		}
		tmap[thead.Id()] = trails
	}
	return tmap
}

func (t Terrain) GetTrailHeadRatings() map[string]int {
	smap := make(map[string]int)
	tmap := t.GetTrailSystem()
	for id, trails := range tmap {
		tids := make([]string, 0)
		for _, trail := range trails {
			ids := make([]string, 0)
			for _, p := range trail {
				ids = append(ids, p.Id())
			}
			tids = append(tids, strings.Join(ids, "|"))
		}
		smap[id] = len(pie.Unique(tids))
	}
	return smap
}

func (t Terrain) GetTrailSystemRating() int {
	smap := t.GetTrailHeadRatings()
	return pie.Sum(pie.Values(smap))
}

func (t Terrain) GetTrailHeadScores() map[string]int {
	smap := make(map[string]int)
	tmap := t.GetTrailSystem()
	for id, trails := range tmap {
		dests := make([]string, 0)
		for _, trail := range trails {
			dests = append(dests, pie.Last(trail).Id())
		}
		smap[id] = len(pie.Unique(dests))
	}
	return smap
}

func (t Terrain) GetTrailSystemScore() int {
	smap := t.GetTrailHeadScores()
	return pie.Sum(pie.Values(smap))
}

func (t *Terrain) UpdatePoint(p Point) {
	t.points[p.Id()] = p
}

// ---------- UTILITIES -----------------------------------
