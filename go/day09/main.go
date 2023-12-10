package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func maybe(err error) {
	if err != nil {
		panic(err)
	}
}

func atoi(s string) int {
	x, err := strconv.Atoi(s)
	maybe(err)
	return x
}

func lines(fname string) []string {
	f, err := os.Open(fname)
	maybe(err)
	contents, err := io.ReadAll(f)
	maybe(err)
	res := strings.Split(string(contents), "\n")
	if res[len(res)-1] == "" {
		// Skip any empty line
		res = res[:len(res)-1]
	}
	return res
}

func part1(fn string) int {
	curr := []int{}
	next := []int{}
	nextReading := func(line string) int {
		// While deltas are non-zero, calculate next list of deltas and tush the
		// last element to a stack. The stack is used to calculate the final number.
		stack := []int{}
		curr = curr[:0]
		next = next[:0]
		for _, s := range strings.Fields(line) {
			curr = append(curr, atoi(s))
		}
		stack = append(stack, curr[len(curr)-1])
		var done bool
		for !done {
			done = true
			next = next[:0]
			for i := 1; i < len(curr); i++ {
				d := curr[i] - curr[i-1]
				if d != 0 {
					done = false
				}
				next = append(next, d)
			}
			stack = append(stack, next[len(next)-1])
			curr, next = next, curr
		}
		// Calculate next value
		var sum int
		for i := range stack {
			sum += stack[i]
		}
		return sum
	}
	ls := lines(fn)
	var sum int
	for i := range ls {
		if ls[i] == "" {
			continue
		}
		sum += nextReading(ls[i])

	}
	return sum
}

func part2(fn string) int {
	curr := []int{}
	next := []int{}
	nextReading := func(line string) int {
		// While deltas are non-zero, calculate next list of deltas and tush the
		// last element to a stack. The stack is used to calculate the final number.
		stack := []int{}
		curr = curr[:0]
		next = next[:0]
		for _, s := range strings.Fields(line) {
			curr = append(curr, atoi(s))
		}
		stack = append(stack, curr[0])
		var done bool
		for !done {
			done = true
			next = next[:0]
			for i := 1; i < len(curr); i++ {
				d := curr[i] - curr[i-1]
				if d != 0 {
					done = false
				}
				next = append(next, d)
			}
			stack = append(stack, next[0])
			curr, next = next, curr
		}
		// Calculate next value
		var res int
		for i := len(stack) - 1; i >= 0; i-- {
			res = stack[i] - res
		}
		return res
	}
	ls := lines(fn)
	nextReading(ls[3])
	var sum int
	for i := range ls {
		if ls[i] == "" {
			continue
		}
		sum += nextReading(ls[i])
	}
	return sum
}

func main() {
	fmt.Println(part1("input"))
	fmt.Println(part2("input"))
}
