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

var idx = map[string]int{"red": 0, "green": 1, "blue": 2}

type game struct {
	idx  int
	sets [][3]int
}

func parse(lines []string) []game {
	n := len(lines)
	res := make([]game, n)
	for i, s := range lines {
		res[i].idx = mustint(strings.Split(strings.Split(s, ":")[0], " ")[1])
		s = strings.Split(s, ":")[1]
		sets := strings.Split(s, ";")
		res[i].sets = make([][3]int, len(sets))
		for j, set := range sets {
			balls := strings.Split(set, ",")
			for _, ball := range balls {
				count := strings.Split(strings.Trim(ball, " "), " ")[0]
				color := strings.Split(strings.Trim(ball, " "), " ")[1]
				res[i].sets[j][idx[color]] = mustint(count)
			}
		}
	}
	return res
}

var counts = [3]int{12, 13, 14}

func part1(fname string) int {
	games := parse(lines(fname))
	var res int
outer:
	for _, g := range games {
		for _, set := range g.sets {
			if set[0] > counts[0] || set[1] > counts[1] || set[2] > counts[2] {
				continue outer
			}
		}
		res += g.idx
	}
	return res
}

func part2(fname string) int {
	games := parse(lines(fname))
	var res int
	for _, g := range games {
		var mins [3]int
		for _, set := range g.sets {
			for i, cnt := range set {
				mins[i] = max(mins[i], cnt)
			}
		}
		res += mins[0] * mins[1] * mins[2]
	}
	return res
}

func main() {
	fmt.Println(part1("input"))
	fmt.Println(part2("input"))
}
