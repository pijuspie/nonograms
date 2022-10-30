package repository

import (
	"testing"
)

var input string = `7 6
2
4
5
5
3 1
1 1
2
2 2
5 1
5 1
3 1
3
1`

func Test_parseInput(t *testing.T) {
	top, left, err := parseInput(input)

	if err != nil {
		t.Error(err)
	}

	if len(top) != 7 {
		t.Error("top length should be 7")
	}

	if top[0][0] != 2 || top[4][1] != 1 {
		t.Error("wrong top values")
	}

	if len(left) != 6 {
		t.Error("left length should be 6")
	}

	if left[0][0] != 2 || top[4][1] != 1 {
		t.Error("wrong left values")
	}
}
