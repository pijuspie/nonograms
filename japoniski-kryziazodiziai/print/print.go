package print

import (
	"fmt"
	"nonograms/solve"
)

func PrintGrid(grid solve.Grid) {
	fmt.Println()

	for _, row := range grid {
		for _, square := range row {
			if square == solve.Filled {
				fmt.Print("# ")
			} else {
				fmt.Print("  ")
			}
		}

		fmt.Println()
	}

	fmt.Println()
}
