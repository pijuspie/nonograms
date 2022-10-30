package solve

import (
	"fmt"
	"testing"
)

func Test_possibleLocations(t *testing.T) {
	result := possibleLocations(3, 1, Line{2, 2, 2, 2, 2})
	fmt.Println(result)
}

func merge(new Line) {
	fmt.Println(new)
}

func Test_possibleLines(t *testing.T) {
	possibleLines(Values{2, 2}, Line{2, 2, 2, 2, 2, 2}, 0, 0, merge)
}

func Test_constructMerge(t *testing.T) {
	var result Line

	line := Line{2, 2, 2, 2, 2, 2}
	merge := constructMerge(&result, line)
	possibleLines(Values{2, 2}, line, 0, 0, merge)

	fmt.Println(result)
}

func Test_solveLine(t *testing.T) {
	result := solveLine(Values{2, 2}, Line{2, 2, 2, 2, 2, 2})
	fmt.Println(result)
}
