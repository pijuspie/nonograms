package solve

import (
	"log"
)

func constructMerge(result *Line, before Line) func(new Line) {
	return func(new Line) {
		if len(*result) == 0 {
			*result = new
		} else {
			for i := 0; i < len(*result); i++ {
				if (*result)[i] != new[i] {
					(*result)[i] = before[i]
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

func solveLine(values Values, line Line) Line {
	var result Line

	possibleLines(values, line, 0, 0, constructMerge(&result, line))

	if len(result) == 0 {
		log.Print("Crossword has a mistake, error: cannot fit all values, continuing")
	}

	return result
}
