package main

import (
	"fmt"
	"os"
	"strings"
)

const BEAM_SYM string = "S"
const SPLITTER_SYM string = "^"

type SplitterRow []int
type BeamRow []int

func main() {
	input, _ := os.ReadFile("input.txt")
	beams := parseInitBeam(string(input))
	splitters := parseSplitters(string(input))

	total_splits := 0
	for _, splitter := range splitters {
		// fmt.Println("beams:", beams, "\nsplitter:", splitter)
		step_splits, next_beams := splitBeams(beams, splitter)
		beams = next_beams
		total_splits += step_splits
	}
	fmt.Println(total_splits)
}

func parseInitBeam(input string) BeamRow {
	beams := make([]int, 0, 0)
	for i, sym := range strings.Split(input, "\n")[0] {
		if string(sym) == BEAM_SYM {
			beams = append(beams, i)
			break
		}
	}
	return beams
}

func parseSplitters(input string) []SplitterRow {
	splitters := make([]SplitterRow, 0, 0)
	for _, row := range strings.Split(input, "\n")[1:] {
		splitter_row := make(SplitterRow, 0, 0)
		for i, sym := range row {
			if string(sym) == SPLITTER_SYM {
				splitter_row = append(splitter_row, i)
			}
		}
		if len(splitter_row) > 0 {
			splitters = append(splitters, splitter_row)
		}
	}
	return splitters
}

// returns number of splits and updated beam row
func splitBeams(beams BeamRow, splitters SplitterRow) (int, BeamRow) {
	next_beams := make(BeamRow, 0, 0)
	beam_id := 0
	splitter_id := 0
	n_splits := 0

	for beam_id < len(beams) && splitter_id < len(splitters) {
		if beams[beam_id] < splitters[splitter_id] {
			next_beams = append(next_beams, beams[beam_id])
			beam_id++
		} else if beams[beam_id] > splitters[splitter_id] {
			splitter_id++
		} else {
			n_splits += 1
			do_add_left := true
			if len(next_beams) >= 2 {
				last := len(next_beams) - 1
				if next_beams[last] == beams[beam_id]-1 {
					do_add_left = false
				}
				if next_beams[last-1] == beams[beam_id]-1 {
					do_add_left = false
				}
			} else if len(next_beams) >= 1 {
				last := len(next_beams) - 1
				if next_beams[last] == beams[beam_id]-1 {
					do_add_left = false
				}
			}
			if do_add_left {
				next_beams = append(next_beams, beams[beam_id]-1)
			}
			// always add right
			next_beams = append(next_beams, beams[beam_id]+1)

			beam_id += 1
			splitter_id += 1
		}
	}
	return n_splits, next_beams
}
