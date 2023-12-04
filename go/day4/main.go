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

func mustint(s string) int {
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
	ls := lines(fn)
	var res int
	for _, l := range ls {
		a := strings.Split(l, ":")[1]
		b := strings.Split(a, "|")
		winning := make(map[int]bool)
		for _, s := range strings.Split(strings.Trim(b[0], " "), " ") {
			if s == "" {
				continue
			}
			x := mustint(s)
			winning[x] = true
		}
		var count int
		for _, s := range strings.Split(strings.Trim(b[1], " "), " ") {
			if s == "" {
				continue
			}
			x := mustint(s)
			if winning[x] {
				count++
			}
		}
		if count > 0 {
			res += (1 << (count - 1))
		}
	}

	return res
}

func part2(fn string) int {
	ls := lines(fn)
	var res int
	n := len(ls)
	mult := make([]int, n)
	for i := range mult {
		mult[i] = 1
	}
	for i, l := range ls {
		a := strings.Split(l, ":")[1]
		b := strings.Split(a, "|")
		winning := make(map[int]bool)
		for _, s := range strings.Split(strings.Trim(b[0], " "), " ") {
			if s == "" {
				continue
			}
			x := mustint(s)
			winning[x] = true
		}
		var count int
		for _, s := range strings.Split(strings.Trim(b[1], " "), " ") {
			if s == "" {
				continue
			}
			x := mustint(s)
			if winning[x] {
				count++
			}
		}
		res += mult[i]
		for j := 0; j < count; j++ {
			mult[i+j+1] += mult[i]
		}
	}

	return res
}

func main() {
	fmt.Println(part1("input"))
	fmt.Println(part2("input"))
}
