package main

import (
	"fmt"
	"japoniski/solve"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func readInput() ([]solve.Values, []solve.Values, error) {
	bytes, err := os.ReadFile("input.txt")

	if err != nil {
		return []solve.Values{}, []solve.Values{}, err
	}

	lines := strings.Split(strings.ReplaceAll(string(bytes), "\r\n", "\n"), "\n")
	var result []solve.Values

	for i, v := range lines {
		values := strings.Split(v, " ")
		result = append(result, make(solve.Values, len(values)))

		for j, v2 := range values {
			number, err := strconv.Atoi(v2)

			if err != nil {
				return []solve.Values{}, []solve.Values{}, err
			}

			result[i][j] = number
		}
	}

	if len(result) != result[0][0]+result[0][1]+1 {
		return []solve.Values{}, []solve.Values{}, fmt.Errorf("not enough lines according to grid size")
	}

	top := result[1 : result[0][0]+1]
	left := result[result[0][0]+1 : result[0][0]+result[0][1]+1]

	return top, left, nil
}

func validateInputValues(values, other []solve.Values) error {
	for _, v := range values {
		var sum int

		for _, v2 := range v {
			sum += v2
		}

		if sum+len(v)-1 > len(other) {
			return fmt.Errorf("more filled squares than line length")
		}
	}

	return nil
}

func validateInput(top, left []solve.Values) error {
	err := validateInputValues(top, left)

	if err != nil {
		return err
	}

	err = validateInputValues(left, top)

	if err != nil {
		return err
	}

	return nil
}

func printResult(grid solve.Grid, start, end int64) {
	fmt.Println()

	y := len(grid[0])

	for i := 0; i < len(grid); i++ {
		var line string

		for j := 0; j < y; j++ {
			square := grid[i][j]

			if square == solve.Filled {
				line += "# "
			} else if square == solve.Empty {
				line += "  "
			}
		}

		fmt.Println(line)
	}

	fmt.Println()

	formated := fmt.Sprintf("Duration: %ds", (end-start)/1000)
	fmt.Println(formated)
}

func fatalError(msg string, err error) {
	if err != nil {
		log.Fatalf(msg, err)
	}
}

func main() {
	top, left, err := readInput()
	fatalError("Failed to read input, error: %s", err)

	err = validateInput(top, left)
	fatalError("Wrong input, error: %s", err)

	start := time.Now().UnixMilli()
	result, err := solve.Solve(top, left)
	end := time.Now().UnixMilli()

	fatalError("Failed to solve, error: %s", err)
	printResult(result, start, end)
}
