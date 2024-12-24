package day24

import (
	"slices"
	"sort"
	"strconv"
	"strings"
)

// ========== DEFINITION ==================================

type Device struct {
	wires map[string]int
	gates []Gate
}

// ========== RECEIVERS ===================================

func (d Device) Checksum() int {
	wires := d.CompleteCircuit()
	checksum := d.WireGroupValue(wires, "z")
	return checksum
}

// ---------- PRIVATE -------------------------------------

func (d Device) CompleteCircuit() map[string]int {
	wires := d.wires
	gates := d.gates
	for len(gates) > 0 {
		new_gates := make([]Gate, 0)
		for _, g := range gates {
			id, val := g.Evaluate(wires)
			if val != -1 {
				wires[id] = val
			} else {
				new_gates = append(new_gates, g)
			}
		}
		gates = new_gates
	}
	return wires
}

func (d Device) WireGroupValue(wires map[string]int, group string) int {
	ids := make([]string, 0)
	vals := make([]string, 0)
	for id := range wires {
		if string(id[0]) == group {
			ids = append(ids, id)
		}
	}
	sort.Strings(ids)
	slices.Reverse(ids)
	for _, id := range ids {
		vals = append(vals, strconv.Itoa(wires[id]))
	}
	binary := strings.Join(vals, "")
	checksum, _ := strconv.ParseInt(binary, 2, 64)
	return int(checksum)
}
