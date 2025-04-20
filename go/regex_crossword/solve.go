package main

import (
	"errors"
	"fmt"
	"log"
	"regexp"
	"regexp/syntax"
)

type set []int

var flags = syntax.MatchNL | syntax.PerlX | syntax.UnicodeGroups

func (s *set) Add(v int) {
	for i := range *s {
		if (*s)[i] == v {
			return
		}
	}
	*s = append(*s, v)
}

func satisfiesAtPos(expr string, r rune, pos int) bool {
	re, err := syntax.Parse(expr, flags)
	if err != nil {
		panic(err)
	}

	re = re.Simplify()
	prog, err := syntax.Compile(re)
	if err != nil {
		panic(err)
	}

	var el int
	s := 0
	step := []int{}
	queue := set([]int{prog.Start})

	for s <= pos {
		for len(queue) > 0 {
			el, queue = queue[0], queue[1:]
			inst := prog.Inst[el]
			switch inst.Op {
			case syntax.InstAlt:
				queue.Add(int(inst.Out))
				queue.Add(int(inst.Arg))
			case syntax.InstCapture:
				queue.Add(int(inst.Out))
			case syntax.InstRune, syntax.InstRune1, syntax.InstRuneAny:
				step = append(step, el)
			}
		}
		for i := range step {
			inst := prog.Inst[step[i]]
			if s == pos && inst.MatchRune(r) {
				return true
			}
			queue = append(queue, int(inst.Out))
		}
		step = []int{}
		s++
	}

	return false
}

func compileRegex(expr []string) ([]*regexp.Regexp, error) {
	res := []*regexp.Regexp{}
	for _, row := range expr {
		re, err := regexp.Compile("^" + row + "$")
		if err != nil {
			return res, err
		}
		res = append(res, re)
	}
	return res, nil
}

func Solve(rows, cols []string) (string, error) {
	rowRe, err := compileRegex(rows)
	if err != nil {
		log.Println(err)
		return "", err
	}

	colRe, err := compileRegex(cols)
	if err != nil {
		return "", err
	}

	var start, end = 32, 90

	cells := len(cols) * len(rows)
	solution := make([]rune, cells)
	for i := range solution {
		solution[i] = rune(start)
	}

	for i := 0; i < cells; {
		r := i / len(cols)
		c := i % len(cols)

		solvedCell := false
	iterate:
		for u := int(solution[i]) + 1; u < end; u++ {
			rn := rune(u)
			colOk := satisfiesAtPos(cols[c], rn, r)
			rowOk := satisfiesAtPos(rows[r], rn, c)
			if colOk && rowOk {
				solution[i] = rn
				solvedCell = true

				if c == len(cols)-1 {
					row := solution[r*len(cols) : i+1]
					if rowRe[r].MatchString(string(row)) {
						solvedCell = true
					} else {
						solvedCell = false
						continue iterate
					}
				}
				if r == len(rows)-1 {
					col := []rune{}
					for j := 0; j < len(rows); j++ {
						col = append(col, solution[j*len(cols)+c])
					}
					if colRe[c].MatchString(string(col)) {
						solvedCell = true
					} else {
						solvedCell = false
						continue iterate
					}
				}
			}
			if solvedCell {
				break iterate
			}
		}

		if !solvedCell {
			solution[i] = rune(start)
			i--
			if i < 0 {
				return "", errors.New("No solution")
			}
		} else {
			i++
		}
	}
	return string(solution), nil
}

func main() {
	vals := []string{"regex_example_0", "regex_example_1"}
	cols := []string{"regex_example_0", "regex_example_1", "regex_example_2", "regex_example_3"}

	solution, _ := Solve(vals, cols)
	fmt.Println(solution)
}
