package repository

import (
	"fmt"
	"nonograms/solve"
)

const SumsDoNotMatch = "top values sum does not match left values sum"

func validateInput(top, left []solve.Values) (e error) {
	topSum, err := validateValues(top, len(left))

	if err != nil {
		return err
	}

	leftSum, err := validateValues(left, len(top))

	if err != nil {
		return err
	}

	if topSum != leftSum {
		return fmt.Errorf(SumsDoNotMatch)
	}

	return nil
}

func validateValues(values []solve.Values, length int) (t int, e error) {
	var total int

	for _, vals := range values {
		var sum int

		for _, value := range vals {
			sum += value
		}

		if sum+len(vals)-1 > length {
			return 0, fmt.Errorf("too many values in line")
		}

		total += sum
	}

	return total, nil
}
