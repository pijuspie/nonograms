package solve

import (
	"fmt"
	"log"
)

type Values []int
type Line []int
type Grid []Line

const Empty = 0
const Filled = 1
const Unknown = 2

func Solve(top, left []Values) Grid {
	x, y := len(top), len(left)
	grid := createGrid(x, y)
	lastGrid := make(Grid, y)

	finished := make(chan bool)
	var done bool
	var changed bool

	for {
		for i := 0; i < y; i++ {
			go solveRow(&grid, left[i], i, finished)
		}

		for i := 0; i < x; i++ {
			go solveColumn(&grid, top[i], i, finished)
		}

		for i := 0; i < x+y; i++ {
			<-finished
		}

		changed = false
		done = true
		PrintGrid(grid)
		for i, row := range grid {
			for j, column := range row {
				if column == Unknown {
					done = false
				}

				if len(lastGrid[0]) != 0 {
					if column != lastGrid[i][j] {
						changed = true
					}
				} else {
					changed = true
				}
			}
		}

		if done {
			break
		} else if !changed {
			log.Print("Crossword has a mistake, error: stuck")
			break
		} else {
			for i, row := range grid {
				lastGrid[i] = make(Line, x)
				copy(lastGrid[i], row)
			}

			lastGrid = grid
		}
	}

	return grid
}

func createGrid(x, y int) Grid {
	grid := make(Grid, y)

	for i := 0; i < y; i++ {
		grid[i] = make(Line, x)
		for j := 0; j < x; j++ {
			grid[i][j] = Unknown
		}
	}

	return grid
}

func solveRow(grid *Grid, values Values, i int, finished chan bool) {
	x := len((*grid)[0])
	done := true

	for j := 0; j < x; j++ {
		if (*grid)[i][j] == Unknown {
			done = false
			break
		}
	}

	if done {
		finished <- true
		return
	}

	newRow := solveLine(values, (*grid)[i])

	if len(newRow) == 0 {
		finished <- false
		return
	}

	(*grid)[i] = newRow
	finished <- true
}

func solveColumn(grid *Grid, values Values, i int, finished chan bool) {
	y := len(*grid)
	done := true

	for j := 0; j < y; j++ {
		if (*grid)[j][i] == Unknown {
			done = false
		}
	}

	if done {
		finished <- true
		return
	}

	column := make(Line, y)

	for j := 0; j < y; j++ {
		column[j] = (*grid)[j][i]
	}

	newColumn := solveLine(values, column)

	if len(newColumn) == 0 {
		finished <- false
		return
	}

	for j := 0; j < y; j++ {
		(*grid)[j][i] = newColumn[j]
	}

	finished <- true
}

// print

func PrintGrid(grid Grid) {
	fmt.Println()

	for _, row := range grid {
		for _, square := range row {
			if square == Filled {
				fmt.Print("# ")
			} else {
				fmt.Print("  ")
			}
		}

		fmt.Println()
	}

	fmt.Println()
}
