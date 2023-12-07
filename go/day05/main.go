package main

import (
	"fmt"
	"io"
	"math"
	"os"
	"sort"
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

type interval struct {
	start, end int
}

func (i interval) overlapswith(other interval) bool {
	if i.end <= other.start {
		return false
	}
	if i.start >= other.end {
		return false
	}
	return true
}

func part1(fn string) int {
	// Given a seed number and a mapping,
	// We get a possible next value for the seed
	// There may be many such values
	// But if no value exists, the same number is used

	seedVals := map[int]int{}

	ls := lines(fn)
	for _, s := range strings.Split(ls[0], " ")[1:] {
		x := atoi(s)
		seedVals[x] = x
	}
	i := 3
	for i < len(ls) {
		next := make(map[int]int)
		for i < len(ls) && ls[i] != "" {
			parts := strings.Split(ls[i], " ")
			dest := atoi(parts[0])
			source := atoi(parts[1])
			delta := atoi(parts[2])
			diff := dest - source
			for _, seed := range seedVals {
				if seed >= source && seed < source+delta {
					next[seed] = seed + diff
				}
			}
			i++
		}
		for _, seed := range seedVals {
			if _, exists := next[seed]; !exists {
				next[seed] = seed
			}
		}
		i += 2
		seedVals = next
	}

	res := math.MaxInt32
	for _, location := range seedVals {
		res = min(res, location)
	}

	return res
}

func part2(fn string) int {
	// For each iteration, we have a list of intervals
	ivals := []interval{}

	ls := lines(fn)
	seeds := strings.Split(ls[0], " ")
	for j := 1; j < len(seeds); j += 2 {
		x := atoi(seeds[j])
		y := atoi(seeds[j+1])
		ivals = append(ivals, interval{x, x + y})
	}
	i := 3
	for i < len(ls) {
		next := []interval{}
		covered := make([][]interval, len(ivals))
		for i < len(ls) && ls[i] != "" {
			parts := strings.Split(ls[i], " ")
			dest := atoi(parts[0])
			source := atoi(parts[1])
			delta := atoi(parts[2])
			diff := dest - source
			sourceival := interval{source, source + delta}
			// Check if any of the input intervals are contained by the source
			// range.
			for j, inival := range ivals {
				if inival.end <= sourceival.start || inival.start >= sourceival.end {
					continue
				}
				// There is some overlap between the two.
				// Clamp the range
				outival := interval{
					max(inival.start, sourceival.start) + diff,
					min(inival.end, sourceival.end) + diff,
				}
				coveredival := interval{
					max(inival.start, sourceival.start),
					min(inival.end, sourceival.end),
				}
				next = append(next, outival)
				covered[j] = append(covered[j], coveredival)
			}
			i++
		}
		// Add any missing ranges to next
		for j := range covered {
			priorRange := ivals[j]
			// Sort covered ranges
			sort.Slice(covered[j], func(k, l int) bool {
				return covered[j][k].start < covered[j][l].start
			})
			if len(covered[j]) == 0 {
				next = append(next, priorRange)
				continue
			}
			// Add missing first interval
			if covered[j][0].start != priorRange.start {
				next = append(next, interval{priorRange.start, covered[j][0].start})
			}
			// Add missing last interval
			if covered[j][len(covered[j])-1].end != priorRange.end {
				next = append(next, interval{covered[j][len(covered[j])-1].end, priorRange.end})
			}
			// Add missing intermittent intervals
			for k := 1; k < len(covered[j]); k++ {
				prev := covered[j][k-1]
				post := covered[j][k]
				if prev.end != post.start {
					next = append(next, interval{prev.end, post.start})
				}
			}
		}
		// Here's the difficult bit- any uncovered range must also be added
		// to the list of next intervals
		i += 2
		ivals = next
	}

	res := math.MaxInt32
	for _, location := range ivals {
		res = min(res, location.start)
	}

	return res
}

func main() {
	fmt.Println(part1("input"))
	fmt.Println(part2("input"))
}
