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

func part1(lines []string) int {
	// This is a classical DP-style problem. At any index in the input, you may
	// choose to place group j or not. If you do, you not only place the group,
	// but also invalidate the space after the end of the group.
	mem := make(map[[2]int]int)
	var res int
	for i := range lines {
		for k := range mem {
			delete(mem, k)
		}
		p := strings.Fields(lines[i])
		row := p[0]
		var groups []int
		for _, s := range strings.Split(p[1], ",") {
			groups = append(groups, atoi(s))
		}
		x := dfs(mem, row, groups, 0, 0)
		res += x
	}

	return res
}

func dfs(mem map[[2]int]int, row string, groups []int, i, j int) int {
	k := [2]int{i, j}
	if v, exists := mem[k]; exists {
		return v
	}
	if j == len(groups) {
		// Ensure last bit can be left empty
		for k := i; k < len(row); k++ {
			if row[k] == '#' {
				return 0
			}
		}
		return 1
	}
	if i >= len(row) {
		if j == len(groups) {
			return 1
		}
		// Unsettled groups, no arrangement
		return 0
	}
	if row[i] == '.' {
		// Can't place anything here, continue
		res := dfs(mem, row, groups, i+1, j)
		mem[k] = res
		return mem[k]
	}
	// Check whether the current group can be placed at the current position
	place := func() int {
		if row[i] == '.' {
			return 0
		}
		// Check if placement here is possible
		if len(row)-i < groups[j] {
			return 0
		}
		for k := i; k < i+groups[j]; k++ {
			if row[k] == '.' {
				return 0
			}
		}
		if i+groups[j] < len(row) && row[i+groups[j]] == '#' {
			return 0
		}
		return dfs(mem, row, groups, i+groups[j]+1, j+1)
	}

	res := place()
	if row[i] == '?' {
		// Try to skip the current position
		res += dfs(mem, row, groups, i+1, j)
	}

	mem[k] = res
	return mem[k]
}

func part2(lines []string) int {
	mem := make(map[[2]int]int)
	var res int
	for i := range lines {
		for k := range mem {
			delete(mem, k)
		}
		p := strings.Fields(lines[i])
		row := p[0]
		for i := 0; i < 4; i++ {
			row += "?" + p[0]
		}
		var groups []int
		for _, s := range strings.Split(p[1], ",") {
			groups = append(groups, atoi(s))
		}
		m := len(groups)
		for i := 0; i < 4; i++ {
			groups = append(groups, groups[:m]...)
		}
		x := dfs(mem, row, groups, 0, 0)
		res += x
	}

	return res
}

func main() {
	fmt.Println(part1(readlines("input")))
	fmt.Println(part2(readlines("input")))
}
