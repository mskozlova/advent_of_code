package main

import (
	"fmt"
	"os"
	"strings"
)

const BEAM_SYM string = "S"
const SPLITTER_SYM string = "^"

type BeamStatus struct {
	pos         int
	is_split    bool
	n_timelines int
}
type SplitterRow map[int]int
type BeamRow map[int]BeamStatus

func main() {
	input, _ := os.ReadFile("input.txt")
	beams := parseInitBeam(string(input))
	splitters := parseSplitters(string(input))

	beam_rows := make([]BeamRow, 0, 0)
	beam_rows = append(beam_rows, beams)
	for _, splitter := range splitters {
		// fmt.Println("beams:", beams, "\nsplitter:", splitter)
		next_beams := splitBeams(&beams, splitter)
		beam_rows = append(beam_rows, next_beams)
		beams = next_beams
	}
	calculateTimelines(beam_rows)
	fmt.Println(beam_rows[0])
}

func parseInitBeam(input string) BeamRow {
	beams := make(BeamRow)
	for i, sym := range strings.Split(input, "\n")[0] {
		if string(sym) == BEAM_SYM {
			beams[i] = BeamStatus{i, false, 0}
			break
		}
	}
	return beams
}

func parseSplitters(input string) []SplitterRow {
	splitters := make([]SplitterRow, 0, 0)
	for _, row := range strings.Split(input, "\n")[1:] {
		splitter_row := make(SplitterRow)
		for i, sym := range row {
			if string(sym) == SPLITTER_SYM {
				splitter_row[i] = i
			}
		}
		if len(splitter_row) > 0 {
			splitters = append(splitters, splitter_row)
		}
	}
	return splitters
}

func splitBeams(beams *BeamRow, splitters SplitterRow) BeamRow {
	next_beams := make(BeamRow)

	for beam_id, beam := range *beams {
		_, is_splitter := splitters[beam.pos]
		if !is_splitter {
			next_beams[beam.pos] = BeamStatus{beam.pos, false, 1}
			beam_id++
		} else {
			status_copy := (*beams)[beam_id]
			status_copy.is_split = true
			(*beams)[beam_id] = status_copy

			next_beams[beam.pos-1] = BeamStatus{beam.pos - 1, false, 1}
			next_beams[beam.pos+1] = BeamStatus{beam.pos + 1, false, 1}
		}
	}
	return next_beams
}

func calculateTimelines(beam_rows []BeamRow) {
	for row := len(beam_rows) - 2; row >= 0; row-- {
		for _, beam := range beam_rows[row] {
			copy := beam
			if !beam.is_split {
				copy.n_timelines = beam_rows[row+1][beam.pos].n_timelines
			} else {
				copy.n_timelines = beam_rows[row+1][beam.pos+1].n_timelines + beam_rows[row+1][beam.pos-1].n_timelines
			}
			beam_rows[row][beam.pos] = copy
		}
	}
}
