package repository

import (
	"nonograms/solve"
	"os"
)

func ReadInput() (t, l []solve.Values, e error) {
	bytes, err := os.ReadFile("input.txt")

	if err != nil {
		return []solve.Values{}, []solve.Values{}, err
	}

	return getInput(string(bytes))
}

func getInput(input string) (t, l []solve.Values, e error) {
	top, left, err := parseInput(input)

	if err != nil {
		return []solve.Values{}, []solve.Values{}, err
	}

	err = validateInput(top, left)

	if err != nil {
		return top, left, err
	}

	return top, left, nil
}
