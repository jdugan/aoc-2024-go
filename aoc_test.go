package aoc

import (
	"testing"

	. "github.com/franela/goblin"

	"aoc/2024/pkg/day01"
	"aoc/2024/pkg/day02"
	"aoc/2024/pkg/day03"
	"aoc/2024/pkg/day04"
	"aoc/2024/pkg/day05"
	"aoc/2024/pkg/day06"
	"aoc/2024/pkg/day07"
	"aoc/2024/pkg/day08"
	"aoc/2024/pkg/day09"
	"aoc/2024/pkg/day10"
	"aoc/2024/pkg/day11"
	"aoc/2024/pkg/day12"
	"aoc/2024/pkg/day13"
	"aoc/2024/pkg/day14"
	"aoc/2024/pkg/day15"
	"aoc/2024/pkg/day16"
	"aoc/2024/pkg/day17"
	"aoc/2024/pkg/day18"
	"aoc/2024/pkg/day19"
	"aoc/2024/pkg/day20"
	"aoc/2024/pkg/day21"
	"aoc/2024/pkg/day22"
	"aoc/2024/pkg/day23"
	"aoc/2024/pkg/day24"
	"aoc/2024/pkg/day25"
)

func Test(t *testing.T) {
	g := Goblin(t)

	g.Describe("AOC", func() {
		g.It("Should get correct answers for Day 01", func(done Done) {
			g.Assert(day01.Puzzle1()).Equal(2375403)
			g.Assert(day01.Puzzle2()).Equal(23082277)
			done()
		})
		g.It("Should get correct answers for Day 02", func(done Done) {
			g.Assert(day02.Puzzle1()).Equal(572)
			g.Assert(day02.Puzzle2()).Equal(612)
			done()
		})
		g.It("Should get correct answers for Day 03", func(done Done) {
			g.Assert(day03.Puzzle1()).Equal(180233229)
			g.Assert(day03.Puzzle2()).Equal(95411583)
			done()
		})
		g.It("Should get correct answers for Day 04", func(done Done) {
			g.Assert(day04.Puzzle1()).Equal(2639)
			g.Assert(day04.Puzzle2()).Equal(2005)
			done()
		})
		g.It("Should get correct answers for Day 05", func(done Done) {
			g.Assert(day05.Puzzle1()).Equal(6260)
			g.Assert(day05.Puzzle2()).Equal(5346)
			done()
		})
		g.It("Should get correct answers for Day 06", func(done Done) {
			g.Assert(day06.Puzzle1()).Equal(5177)
			// g.Assert(day06.Puzzle2()).Equal(1686) // SLOW, ~120s
			done()
		})
		g.It("Should get correct answers for Day 07", func(done Done) {
			g.Assert(day07.Puzzle1()).Equal(1399219271639)
			g.Assert(day07.Puzzle2()).Equal(275791737999003)
			done()
		})
		g.It("Should get correct answers for Day 08", func(done Done) {
			g.Assert(day08.Puzzle1()).Equal(413)
			g.Assert(day08.Puzzle2()).Equal(1417)
			done()
		})
		g.It("Should get correct answers for Day 09", func(done Done) {
			g.Assert(day09.Puzzle1()).Equal(6415184586041)
			g.Assert(day09.Puzzle2()).Equal(6436819084274)
			done()
		})
		g.It("Should get correct answers for Day 10", func(done Done) {
			g.Assert(day10.Puzzle1()).Equal(566)
			g.Assert(day10.Puzzle2()).Equal(1324)
			done()
		})
		g.It("Should get correct answers for Day 11", func(done Done) {
			g.Assert(day11.Puzzle1()).Equal(194482)
			g.Assert(day11.Puzzle2()).Equal(232454623677743)
			done()
		})
		g.It("Should get correct answers for Day 12", func(done Done) {
			g.Assert(day12.Puzzle1()).Equal(1473276)
			g.Assert(day12.Puzzle2()).Equal(901100)
			done()
		})
		g.It("Should get correct answers for Day 13", func(done Done) {
			g.Assert(day13.Puzzle1()).Equal(31761)
			g.Assert(day13.Puzzle2()).Equal(90798500745591)
			done()
		})
		g.It("Should get correct answers for Day 14", func(done Done) {
			g.Assert(day14.Puzzle1()).Equal(218965032)
			g.Assert(day14.Puzzle2()).Equal(7037)
			done()
		})
		g.It("Should get correct answers for Day 15", func(done Done) {
			g.Assert(day15.Puzzle1()).Equal(1398947)
			g.Assert(day15.Puzzle2()).Equal(1397393)
			done()
		})
		g.It("Should get correct answers for Day 16", func(done Done) {
			g.Assert(day16.Puzzle1()).Equal(123540)
			g.Assert(day16.Puzzle2()).Equal(665)
			done()
		})
		g.It("Should get correct answers for Day 17", func(done Done) {
			g.Assert(day17.Puzzle1()).Equal("4,1,7,6,4,1,0,2,7")
			g.Assert(day17.Puzzle2()).Equal(164279024971453)
			done()
		})
		g.It("Should get correct answers for Day 18", func(done Done) {
			g.Assert(day18.Puzzle1()).Equal(336)
			g.Assert(day18.Puzzle2()).Equal("24,30")
			done()
		})
		g.It("Should get correct answers for Day 19", func(done Done) {
			g.Assert(day19.Puzzle1()).Equal(371)
			g.Assert(day19.Puzzle2()).Equal(650354687260341)
			done()
		})
		g.It("Should get correct answers for Day 20", func(done Done) {
			g.Assert(day20.Puzzle1()).Equal(1399)
			g.Assert(day20.Puzzle2()).Equal(994807)
			done()
		})
		g.It("Should get correct answers for Day 21", func(done Done) {
			g.Assert(day21.Puzzle1()).Equal(125742)
			g.Assert(day21.Puzzle2()).Equal(157055032722640)
			done()
		})
		g.It("Should get correct answers for Day 22", func(done Done) {
			g.Assert(day22.Puzzle1()).Equal(20068964552)
			// g.Assert(day22.Puzzle2()).Equal(2246) // slow, ~25s
			done()
		})
		g.It("Should get correct answers for Day 23", func(done Done) {
			g.Assert(day23.Puzzle1()).Equal(1110)
			g.Assert(day23.Puzzle2()).Equal("ej,hm,ks,ms,ns,rb,rq,sc,so,un,vb,vd,wd") // pretty slow, ~8s
			done()
		})
		g.It("Should get correct answers for Day 24", func(done Done) {
			g.Assert(day24.Puzzle1()).Equal(65635066541798)
			g.Assert(day24.Puzzle2()).Equal("dgr,dtv,fgc,mtj,vvm,z12,z29,z37")
			done()
		})
		g.It("Should get correct answers for Day 25", func(done Done) {
			g.Assert(day25.Puzzle1()).Equal(3525)
			g.Assert(day25.Puzzle2()).Equal(50)
			done()
		})
	})
}
