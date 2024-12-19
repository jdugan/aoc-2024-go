package day19

import (
	"fmt"
	"strings"
)

// ========== DEFINITION ==================================

type Design struct {
	pattern string
}

// ========== RECEIVERS ===================================

func (d Design) FindSequences(towels []string) []Sequence {
	seqs := make([]Sequence, 0)
	possibles := []Sequence{Sequence{pattern: d.pattern}}
	for len(possibles) > 0 {
		fmt.Println("-------------------")
		// fmt.Println(possibles)
		fmt.Println(len(possibles))
		fmt.Println(len(possibles[0].pattern))
		new_possibles := make([]Sequence, 0)
		for _, seq := range possibles {
			for _, towel := range towels {
				if len(seq.pattern) >= len(towel) && towel == seq.pattern[:len(towel)] {
					nts := append(seq.towels, towel)
					npt := strings.Replace(seq.pattern, towel, "", 1)
					nseq := Sequence{towels: nts, pattern: npt}
					if npt == "" {
						seqs = append(seqs, nseq)
					} else {
						new_possibles = append(new_possibles, nseq)
					}
				}
			}
		}
		possibles = new_possibles
	}
	fmt.Println(d.pattern, len(seqs), seqs)
	return seqs
}

// ---------- UTILITIES -----------------------------------
