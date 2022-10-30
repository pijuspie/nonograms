package solve

import (
	"fmt"
	"testing"
)

func Test_SolveRow(t *testing.T) {
	grid := Grid{Line{2, 2, 2, 2}}
	finished := make(chan bool)

	go solveRow(&grid, Values{4}, 0, finished)
	<-finished

	fmt.Println(grid)
}

func Test_SolveColumn(t *testing.T) {
	grid := Grid{Line{2}, Line{2}, Line{2}, Line{2}}
	finished := make(chan bool)

	go solveColumn(&grid, Values{4}, 0, finished)
	<-finished

	fmt.Println(grid)
}
