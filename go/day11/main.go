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

func readlines(fname string) []string {
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

func expand(universe []string) [][]byte {
	n := len(universe[0])
	m := len(universe)
	colCount := make([]int, n)
	rowCount := make([]int, m)
	for i := range universe {
		for j, v := range universe[i] {
			if v == '#' {
				colCount[j]++
				rowCount[i]++
			}
		}
	}
	res := make([][]byte, 0, m)
	var k int
	for i := range universe {
		res = append(res, make([]byte, 0, n))
		for j, v := range universe[i] {
			res[k] = append(res[k], byte(v))
			if colCount[j] == 0 {
				res[k] = append(res[k], '.')
			}
		}
		if rowCount[i] == 0 {
			res = append(res, make([]byte, len(res[k])))
			k++
			for j := range res[k] {
				res[k][j] = '.'
			}
		}
		k++
	}
	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func part1(lines []string) int {
	expanded := expand(lines)
	var galaxies [][2]int
	var res int
	for i := range expanded {
		for j, v := range expanded[i] {
			if v == '.' {
				continue
			}
			for _, other := range galaxies {
				res += abs(other[0]-i) + abs(other[1]-j)
			}
			galaxies = append(galaxies, [2]int{i, j})
		}
	}
	return res
}

func part2(lines []string) int {
	universe := lines
	n := len(universe[0])
	m := len(universe)
	colCount := make([]int, n)
	rowCount := make([]int, m)
	for i := range universe {
		for j, v := range universe[i] {
			if v == '#' {
				colCount[j]++
				rowCount[i]++
			}
		}
	}

	galaxies := make(map[[2]int]struct{})
	var res int
	for i := range universe {
		for j, v := range universe[i] {
			if v == '.' {
				continue
			}
			for g := range galaxies {
				for k := min(i, g[0]); k < max(i, g[0]); k++ {
					if rowCount[k] == 0 {
						res += 1_000_000
					} else {
						res++
					}
				}
				for k := min(j, g[1]); k < max(j, g[1]); k++ {
					if colCount[k] == 0 {
						res += 1_000_000
					} else {
						res++
					}
				}
			}
			galaxies[[2]int{i, j}] = struct{}{}
		}
	}

	return res
}

func main() {
	fmt.Println(part1(readlines("input")))
	fmt.Println(part2(readlines("input")))
}
