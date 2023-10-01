package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

const UNKNOWN = 0
const FILLED = 1
const EMPTY = 2

type lineProps struct {
	line   []int
	values []int
	start  int
	index  int
}

func main() {
	file := flag.String("file", "input.txt", "input data file name")
	pins := flag.String("pins", "", "pin initial values 'x:y=1;x:y=1'")
	flag.Parse()

	cV, lV := readData(*file)
	t := createTable(len(cV), len(lV), *pins)

	start := time.Now().UnixNano()
	r, iterations := solve(t, cV, lV)
	end := time.Now().UnixNano()
	duration := math.Round(float64(end-start) / 1000000)

	printTable(r, len(cV))
	fmt.Printf("Iterations: %v, duration: %v ms\r\n", iterations, duration)
}

func readData(filename string) (cV [][]int, lV [][]int) {
	d, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("Failed to read data file, err: %s", err)
	}

	r := strings.Split(string(d), "\r\n")
	v := make([][]int, len(r))

	for i := 0; i < len(r); i++ {
		s := strings.Split(r[i], " ")
		v[i] = make([]int, len(s))

		for j := 0; j < len(s); j++ {
			v[i][j], err = strconv.Atoi(s[j])
			if err != nil {
				log.Fatal("Failed to parse data")
			}
		}
	}

	cV, lV = v[1:v[0][0]+1], v[v[0][0]+1:]
	cS, lS := 0, 0

	for i := 0; i < len(cV); i++ {
		for j := 0; j < len(cV[i]); j++ {
			cS += cV[i][j]
		}
	}

	for i := 0; i < len(lV); i++ {
		for j := 0; j < len(lV[i]); j++ {
			lS += lV[i][j]
		}
	}

	if cS != lS {
		log.Fatal("Columns and rows values do not match")
	}

	return
}

func createTable(w, h int, pins string) []int {
	t := make([]int, w*h)
	p := strings.Split(pins, ";")

	for i := 0; i < len(p); i++ {
		p1 := p[i]
		if p1 == "" {
			continue
		}

		p2 := strings.Split(p1, "=")
		p3 := strings.Split(p2[0], ":")

		if len(p2) != 2 || len(p3) != 2 {
			log.Fatal("Invalid pins")
		}

		v, err := strconv.Atoi(p2[1])
		if err != nil || (v != 0 && v != 1 && v != 2) {
			log.Fatal("Invalid pin value")
		}

		x, err := strconv.Atoi(p3[0])
		if err != nil || x < 0 || x >= w {
			log.Fatal("Invalid pin x value")
		}

		y, err := strconv.Atoi(p3[1])
		if err != nil || y < 0 || y >= h {
			log.Fatal("Invalid pin y value")
		}

		t[w*y+x] = v
	}

	return t
}

func printTable(t []int, w int) {
	var s string

	for i := 0; i < len(t)/w; i++ {
		r := t[i*w : (i+1)*w]

		for i := 0; i < len(r); i++ {
			if r[i] == FILLED {
				s += "X "
			} else if r[i] == EMPTY {
				s += ". "
			} else {
				s += "  "
			}
		}

		s += "\r\n"
	}

	fmt.Println()
	fmt.Println(s)
}

func solve(t []int, cV, lV [][]int) (r []int, iterations int) {
	h, w := len(lV), len(cV)
	r = append(r, t...)

	tasks := make([]int, h+w)
	for i := 0; i < w+h; i++ {
		tasks[i] = i
	}

	fmt.Println("Starting. Tasks: " + fmt.Sprint(len(tasks)))

	for {
		iterations++
		fmt.Printf("%v (%v left)", iterations, len(tasks))
		fmt.Println()

		task := tasks[0]
		var l, a []int

		if task < h {
			l = r[task*w : (task+1)*w]
			a = analyze(l, lV[task])

			for i := 0; i < w; i++ {
				if l[i] != a[i] {
					is := false
					for j := 0; j < len(tasks); j++ {
						if tasks[j] == h+i {
							is = true
						}
					}

					if !is {
						tasks = append(tasks, h+i)
					}
				}
			}

			r = append(r[:task*w], append(a, r[(task+1)*w:]...)...)
		} else {
			task = task - h

			l = make([]int, h)
			for i := 0; i < h; i++ {
				l[i] = r[i*w+task]
			}

			a = analyze(l, cV[task])
			for i := 0; i < h; i++ {
				if l[i] != a[i] {
					is := false
					for j := 0; j < len(tasks); j++ {
						if tasks[j] == i {
							is = true
						}
					}

					if !is {
						tasks = append(tasks, i)
					}
				}
			}

			for i := 0; i < h; i++ {
				r[i*w+task] = a[i]
			}
		}

		if len(tasks) == 1 {
			return
		} else {
			tasks = tasks[1:]
		}
	}
}

func analyze(l, v []int) (r []int) {
	o := optionsAll(l, v)

	fmt.Println(len(o))

	if len(o) < 1 {
		log.Fatal("Line cannot be solved")
	}

	for i := 0; i < len(l); i++ {
		r = append(r, o[0][i])

		if o[0][i] == UNKNOWN {
			continue
		}

		for j := 1; j < len(o); j++ {
			if o[j][i] != r[i] {
				r[i] = UNKNOWN
				break
			}
		}
	}

	return
}

func optionsAll(l, v []int) (r [][]int) {
	start := &lineProps{
		values: v,
		start:  0,
		index:  0,
	}

	start.line = append(start.line, l...)
	t := []*lineProps{start}

	var s int
	for i := 0; i < len(v); i++ {
		s += v[i]
	}

	for {
		t1 := t[0]
		o := options(t1)

		if t1.index+1 == len(t1.values) {
			for i := 0; i < len(o); i++ {
				var t int
				for j := 0; j < len(l); j++ {
					if o[i].line[j] == FILLED {
						t++
					}

					if o[i].line[j] == UNKNOWN {
						o[i].line[j] = EMPTY
					}
				}

				if t == s {
					r = append(r, o[i].line)
				}
			}
		} else {
			for i := 0; i < len(o); i++ {
				t = append(t, o[i])
			}
		}

		if len(t) == 1 {
			return
		} else {
			t = t[1:]
		}
	}
}

func options(l *lineProps) (r []*lineProps) {
	var s int
	for i := l.index; i < len(l.values); i++ {
		s += l.values[i] + 1
	}

	for i := l.start; i <= len(l.line)-s+1; i++ {
		n := &lineProps{
			values: l.values,
			index:  l.index,
			start:  i,
		}

		n.line = append(n.line, l.line...)
		if succ := draw(n); succ {
			r = append(r, n)
		}
	}

	return
}

func draw(l *lineProps) bool {
	n := l.values[l.index]

	for i := 0; i < n; i++ {
		if l.line[l.start+i] == EMPTY {
			return false
		}
		l.line[l.start+i] = FILLED
	}

	if l.start > 0 {
		if l.line[l.start-1] == FILLED {
			return false
		}
		l.line[l.start-1] = EMPTY
	}

	if l.start+n < len(l.line) {
		if l.line[l.start+n] == FILLED {
			return false
		}
		l.line[l.start+n] = EMPTY
	}

	l.start += n
	l.index++

	return true
}
