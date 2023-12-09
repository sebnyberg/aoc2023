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

func part1(fname string) int {
	ls := lines(fname)

	nodes := make(map[string][2]string)
	for i := 2; i < len(ls); i++ {
		parts := strings.Fields(ls[i])
		nodes[parts[0]] = [2]string{
			parts[2][1 : len(parts[2])-1],
			parts[3][:len(parts[3])-1],
		}
	}

	cur := "AAA"
	var i int
	m := len(ls[0])
	mvtIdx := [...]int{'L': 0, 'R': 1}
	for i = 0; cur != "ZZZ"; i++ {
		mvt := ls[0][i%m]
		cur = nodes[cur][mvtIdx[mvt]]
	}
	return i
}

func part2(fname string) int {
	// Finally a problem that is a little bit more interesting.
	//
	// I tried the naive way and it did not work. This tells me that while the
	// paths do end up at end-Z locations at some point, the way to get there
	// will not be by regular exploration.
	//
	// Rather, for each starting position, due to the pigeon-hole principle,
	// there is a finite number of positions that the node can end up at during
	// its traversal. In other words, the first time that a path ends up in the
	// same location of the instructions, and at the same location in terms of
	// the list of movements, then we know that there are no more states to be
	// discovered. In fact, there may even be infinite cycles that are even
	// shorter.
	//
	ls := lines(fname)

	var curr []string
	nodes := make(map[string][2]string)
	for i := 2; i < len(ls); i++ {
		parts := strings.Fields(ls[i])
		if parts[0][len(parts[0])-1] == 'A' {
			curr = append(curr, parts[0])
		}
		nodes[parts[0]] = [2]string{
			parts[2][1 : len(parts[2])-1],
			parts[3][:len(parts[3])-1],
		}
	}

	mvtIdx := [...]int{'L': 0, 'R': 1}
	type key struct {
		location string
		mvtIdx   int
	}
	mvt := ls[0]
	m := len(mvt)
	findCycleLength := func(curr string) int {
		seen := make(map[key]int)
		var times []int
		var t int
		for {
			k := key{curr, t % m}
			if _, exists := seen[k]; exists {
				break
			}
			seen[k] = t
			if curr[len(curr)-1] == 'Z' {
				times = append(times, t)
			}
			curr = nodes[curr][mvtIdx[mvt[t%m]]]
			t++
		}
		head := seen[key{curr, t % m}]
		tail := t - times[0]
		if head != tail {
			panic("cycle with offset")
		}
		return times[0]
	}

	gcd := func(a, b int) int {
		for b != 0 {
			a, b = b, a%b
		}
		return a
	}

	lcm := 1
	for i := range curr {
		delta := findCycleLength(curr[i])
		lcm = lcm * delta / gcd(lcm, delta)
	}

	return lcm
}

func main() {
	fmt.Println(part1("input"))
	fmt.Println(part2("input"))
}
