package solve

import (
	"fmt"
)

type Values []int
type Line []int
type Grid []Line

const Filled = 1
const Empty = 0
const Unknown = 2

func createGrid(x, y int) Grid {
	grid := make([]Line, y)

	for i := 0; i < y; i++ {
		grid[i] = make(Line, x)
		for j := 0; j < x; j++ {
			grid[i][j] = Unknown
		}
	}

	return grid
}

func constructMerge(result *Line, before Line, changed *bool) func(new Line) {
	return func(new Line) {
		if len(*result) == 0 {
			*result = new
		} else {
			for i := 0; i < len(*result); i++ {
				if (*result)[i] != new[i] {
					(*result)[i] = before[i]
				} else {
					*changed = true
				}
			}
		}
	}
}

func possibleLocations(value, start int, line Line) []int {
	var result []int

	for i := start; i < len(line)-value+1; i++ {
		if (i != 0 && line[i-1] == Filled) || (i != len(line)-value && line[i+value] == Filled) {
			continue
		}

		possible := true

		for j := 0; j < value; j++ {
			if line[i+j] == Empty {
				possible = false
				break
			}
		}

		if possible {
			result = append(result, i)
		}
	}

	return result
}

func possibleLines(values Values, line Line, index, start int, merge func(new Line)) {
	value := values[index]
	locations := possibleLocations(value, start, line)

	for _, v := range locations {
		l := make(Line, len(line))
		copy(l, line)

		if v != 0 {
			l[v-1] = Empty
		}

		if v != len(line)-value {
			l[v+value] = Empty
		}

		for i := 0; i < value; i++ {
			l[v+i] = Filled
		}

		if len(values) == index+1 {
			var total, expected int

			for i := 0; i < len(l); i++ {
				if l[i] == Unknown {
					l[i] = Empty
				}

				if l[i] == Filled {
					total++
				}
			}

			for _, v := range values {
				expected += v
			}

			if total == expected {
				merge(l)
			}
		} else {
			possibleLines(values, l, index+1, v+value+1, merge)
		}
	}
}

func solveLine(values Values, line Line) (Line, bool, error) {
	var result Line
	changed := false

	possibleLines(values, line, 0, 0, constructMerge(&result, line, &changed))

	if len(result) == 0 {
		return line, false, nil // fmt.Errorf("cannot fit all values")
	}

	for i := 0; i < len(line); i++ {
		if line[i] != result[i] {
			changed = true
		}
	}

	return result, changed, nil
}

func Solve(top, left []Values) (Grid, error) {
	x, y := len(top), len(left)
	grid := createGrid(x, y)

	var changed, finished bool

	for {
		changed, finished = false, true

		for i := 0; i < y; i++ {
			done := true

			for j := 0; j < x; j++ {
				if grid[i][j] == Unknown {
					done = false
					break
				}
			}

			if !done {
				var err error
				grid[i], changed, err = solveLine(left[i], grid[i])

				if err != nil {
					return Grid{}, err
				}

				finished = false
			}
		}

		for i := 0; i < x; i++ {
			done := true

			column := make(Line, y)
			var newColumn Line

			for j := 0; j < y; j++ {
				column[j] = grid[j][i]

				if column[j] == Unknown {
					done = false
				}
			}

			if !done {
				var err error
				newColumn, changed, err = solveLine(top[i], column)

				if err != nil {
					return Grid{}, err
				}

				for j := 0; j < y; j++ {
					grid[j][i] = newColumn[j]
				}

				finished = false
			}
		}

		if !changed {
			break
		}
	}

	if !finished {
		return Grid{}, fmt.Errorf("stuck")
	}

	return grid, nil
}
