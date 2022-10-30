package main

import (
	"fmt"
	"log"
	"nonograms/print"
	"nonograms/repository"
	"nonograms/solve"
	"time"
)

func logError(msg string, err error) {
	if err != nil {
		if err.Error() != repository.SumsDoNotMatch {
			log.Fatalf(msg, err)
		} else {
			log.Printf("Crossword has a mistake: %s, continuing", err)
		}
	}
}

func main() {
	top, left, err := repository.ReadInput()
	logError("Failed to get input, error: %s", err)

	start := time.Now()
	grid := solve.Solve(top, left)

	print.PrintGrid(grid)
	fmt.Printf("Duration: %d ms", time.Since(start)/1000/1000)
}
