package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

type direction struct {
	v, h int
}

type position struct {
	row, col int
}

type field struct {
	field []string
	boxes map[position]bool
}

type connected_box struct {
	row, col, layer int
}

func main() {
	file_name := "2024/day15_pt2/input.txt"
	input, err := os.ReadFile(file_name)
	if err != nil {
		fmt.Printf("No such file %s\n", file_name)
		return
	}
	start, fld, directions := parseInput(string(input))
	fmt.Println(start)
	printField(start, &fld)
	end_pos := runRobot(start, &fld, directions)
	printField(end_pos, &fld)
	fmt.Println(getBoxScore(&fld))
}

func getBoxScore(fld *field) int {
	total_score := 0
	for box := range (*fld).boxes {
		score := 100*box.row + box.col
		total_score += score
	}
	return total_score
}

func printField(pos position, fld *field) {
	print := make([]string, len((*fld).field))
	for row := range len((*fld).field) {
		row_str := ""
		for col := range len((*fld).field[0]) {
			_, is_box := (*fld).boxes[position{row, col}]
			if string((*fld).field[row][col]) == "#" {
				row_str += "#"
			} else if (col != 0) && (string(row_str[col-1]) == "[") {
				row_str += "]"
			} else if (pos.row == row) && (pos.col == col) {
				row_str += "@"
			} else if is_box {
				row_str += "["
			} else {
				row_str += "."
			}
		}
		print[row] = row_str
	}
	fmt.Println(strings.Join(print, "\n"))
}

func runRobot(start position, fld *field, directions []direction) position {
	pos := start
	for i, dir := range directions {
		connected_boxes := make(map[connected_box]bool)
		can_do := checkStep(pos, fld, dir, false, 0, &connected_boxes)
		if can_do {
			pos = makeStep(pos, fld, dir, &connected_boxes)
		}
		fmt.Printf("\nStep #%d (%d, %d) - %t - pos after (%d, %d)\n", i, dir.v, dir.h, can_do, pos.row, pos.col)
		// fmt.Println("Connected boxes:", connected_boxes)
		printField(pos, fld)
	}
	return pos
}

func canMoveOneLayer(pos position, dir direction, fld *field) (bool, []position) {
	if (pos.row < 0) || (pos.row >= len((*fld).field)) {
		return false, nil
	}
	if (pos.col < 0) || (pos.col >= len((*fld).field[0])) {
		return false, nil
	}
	if string((*fld).field[pos.row][pos.col]) == "#" {
		return false, nil
	}

	direct := position{row: -1, col: -1}
	shifted := position{row: -1, col: -1}
	if dir.v == 1 {
		direct = pos
		shifted = position{row: pos.row, col: pos.col - 1}
	} else if dir.v == -1 {
		direct = pos
		shifted = position{row: pos.row, col: pos.col - 1}
	} else if dir.h == 1 {
		direct = pos
	} else { // dir.h == -1
		direct = position{row: pos.row, col: pos.col - 1}
	}

	// fmt.Println("direct:", direct, "shifted:", shifted)
	_, has_direct := (*fld).boxes[direct]
	if has_direct {
		return true, []position{direct}
	}
	_, has_shifted := (*fld).boxes[shifted]
	if has_shifted {
		return true, []position{shifted}
	}

	return true, nil
}

func getConnectionPoints(pos position, dir direction, is_box bool) []position {
	if !is_box {
		next_pos := position{row: pos.row + dir.v, col: pos.col + dir.h}
		return []position{next_pos}
	}

	if dir.v == 1 {
		return []position{
			{row: pos.row + 1, col: pos.col},
			{row: pos.row + 1, col: pos.col + 1},
		}
	}

	if dir.v == -1 {
		return []position{
			{row: pos.row - 1, col: pos.col},
			{row: pos.row - 1, col: pos.col + 1},
		}
	}

	if dir.h == 1 {
		return []position{
			{row: pos.row, col: pos.col + 2},
		}
	}

	// dir.h == -1
	return []position{
		{row: pos.row, col: pos.col - 1},
	}
}

func checkStep(pos position, fld *field, dir direction, is_box bool, layer int, connected_boxes *map[connected_box]bool) bool {
	connection_points := getConnectionPoints(pos, dir, is_box)
	// if !is_box {
	// 	fmt.Println("connection points:", connection_points)
	// }

	for _, conn_point := range connection_points {
		can_move_point, next_obj := canMoveOneLayer(conn_point, dir, fld)
		fmt.Println(can_move_point, next_obj)
		if !can_move_point {
			return false
		}

		if next_obj != nil {
			(*connected_boxes)[connected_box{next_obj[0].row, next_obj[0].col, layer}] = true
			if !checkStep(next_obj[0], fld, dir, true, layer+1, connected_boxes) {
				return false
			}
		}
	}

	return true
}

func makeStep(pos position, fld *field, dir direction, connected_boxes *map[connected_box]bool) position {
	// fmt.Println("connected boxes:", *connected_boxes)
	next_pos := position{row: pos.row + dir.v, col: pos.col + dir.h}

	connected_boxes_lst := make([]connected_box, len(*connected_boxes))
	i := 0
	for box := range *connected_boxes {
		connected_boxes_lst[i] = box
		i += 1
	}

	slices.SortFunc(connected_boxes_lst, func(a, b connected_box) int {
		if a.layer > b.layer {
			return 1
		}
		if a.layer < b.layer {
			return -1
		}
		return 0
	})
	slices.Reverse(connected_boxes_lst)
	for _, box := range connected_boxes_lst {
		box_pos := position{row: box.row, col: box.col}
		next_box_pos := position{row: box.row + dir.v, col: box.col + dir.h}
		(*fld).boxes[next_box_pos] = true
		delete((*fld).boxes, box_pos)
		// fmt.Printf("Deleting box %d, %d, layer %d\n", box_pos.row, box_pos.col, box.layer)
		// fmt.Printf("Adding box %d, %d\n", next_box_pos.row, next_box_pos.col)
	}

	return next_pos
}

func parseInput(input string) (position, field, []direction) {
	parts := strings.Split(input, "\n\n")
	pos, fld := parseField(parts[0])
	return pos, fld, parseDirections(parts[1])
}

func parseDirections(dir_raw string) []direction {
	dir_raw = strings.Replace(dir_raw, "\n", "", -1)
	directions := make([]direction, len(dir_raw))
	for i, sym := range dir_raw {
		if string(sym) == "<" {
			directions[i] = direction{v: 0, h: -1}
		} else if string(sym) == ">" {
			directions[i] = direction{v: 0, h: 1}
		} else if string(sym) == "^" {
			directions[i] = direction{v: -1, h: 0}
		} else if string(sym) == "v" {
			directions[i] = direction{v: 1, h: 0}
		}
	}
	return directions
}

func parseField(field_raw string) (position, field) {
	rows_raw := strings.Split(field_raw, "\n")
	fld := make([]string, len(rows_raw))
	start := position{}
	boxes := make(map[position]bool)

	for r, row_raw := range rows_raw {
		row := ""
		for c, sym := range row_raw {
			if string(sym) == "@" {
				start.row = r
				start.col = 2 * c
				row += ".."
			} else if string(sym) == "O" {
				boxes[position{r, 2 * c}] = true
				row += ".."
			} else if string(sym) == "#" {
				row += "##"
			} else {
				row += ".."
			}
		}
		fld[r] = row
	}
	return start, field{fld, boxes}
}
