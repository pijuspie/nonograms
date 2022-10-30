package repository

import (
	"fmt"
	"nonograms/solve"
	"strconv"
	"strings"
)

func parseInput(input string) (t, l []solve.Values, e error) {
	input = strings.ReplaceAll(input, "\r\n", "\n")
	lines := strings.Split(input, "\n")

	if len(lines) < 3 {
		return []solve.Values{}, []solve.Values{}, fmt.Errorf("not enough lines")
	}

	values := make([]solve.Values, len(lines))

	for i, line := range lines {
		strings := strings.Split(line, " ")
		values[i] = make(solve.Values, len(strings))

		if len(strings) == 0 {
			return []solve.Values{}, []solve.Values{}, fmt.Errorf("must be at least one value in line")
		}

		for j, string := range strings {
			value, err := strconv.Atoi(string)

			if err != nil {
				return []solve.Values{}, []solve.Values{}, err
			}

			if value <= 0 {
				return []solve.Values{}, []solve.Values{}, fmt.Errorf("value must be greater than 0")
			}

			values[i][j] = value
		}
	}

	x, y := values[0][0], values[0][1]

	if len(values) != x+y+1 {
		return []solve.Values{}, []solve.Values{}, fmt.Errorf("not enough lines according to grid size")
	}

	top := values[1 : x+1]
	left := values[x+1 : x+y+1]

	return top, left, nil
}
