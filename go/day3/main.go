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

var dirs = [][2]int{{1, 0}, {1, -1}, {0, -1}, {-1, -1}, {-1, 0}, {-1, 1}, {0, 1}, {1, 1}}

func part1(fn string) int {
	ls := lines(fn)

	// For each number,
	// For each digit in the number,
	// Look for an adjacent symbol,
	// If the symbol exists, add the part number
	isnum := func(ch byte) bool {
		return ch >= '0' && ch <= '9'
	}

	var res int
	for i, line := range ls {
		var j int
		for j < len(line) {
			if !isnum(line[j]) {
				j++
				continue
			}
			l := j
			r := l + 1
			for r < len(line) && isnum(line[r]) {
				r++
			}
			x := mustint(line[l:r])
		outer:
			for k := l; k < r; k++ {
				for _, d := range dirs {
					ii := i + d[0]
					kk := k + d[1]
					if ii < 0 || ii >= len(ls) || kk < 0 || kk >= len(line) {
						continue
					}
					if !isnum(ls[ii][kk]) && ls[ii][kk] != '.' {
						res += x
						break outer
					}
				}
			}

			j = r
		}
	}

	return res
}

func part2(fn string) int {
	ls := lines(fn)

	// For each number,
	// For each digit in the number,
	// Look for an adjacent symbol,
	// If the symbol exists, add the part number
	isnum := func(ch byte) bool {
		return ch >= '0' && ch <= '9'
	}

	gears := make(map[[2]int][]int)

	var res int
	for i, line := range ls {
		var j int
		for j < len(line) {
			if !isnum(line[j]) {
				j++
				continue
			}
			l := j
			r := l + 1
			for r < len(line) && isnum(line[r]) {
				r++
			}
			x := mustint(line[l:r])
		outer:
			for k := l; k < r; k++ {
				for _, d := range dirs {
					ii := i + d[0]
					kk := k + d[1]
					if ii < 0 || ii >= len(ls) || kk < 0 || kk >= len(line) {
						continue
					}
					if ls[ii][kk] == '*' {
						gears[[2]int{ii, kk}] = append(gears[[2]int{ii, kk}], x)
						break outer
					}
				}
			}

			j = r
		}
	}

	for _, nums := range gears {
		if len(nums) == 2 {
			res += nums[0] * nums[1]
		}
	}

	return res
}

func main() {
	fmt.Println(part1("input"))
	fmt.Println(part2("input"))
}
