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

func solve(times, dists []int) int {
	calc := func(tHeld, tTotal int) int {
		return (tTotal - tHeld) * tHeld
	}
	_ = calc

	findPeak := func(tTotal int) int {
		lo := 0
		hi := tTotal
		for hi-lo > 1 {
			mid := lo + (hi-lo)/2
			cmid := calc(mid, tTotal)
			if calc(mid-1, tTotal) < cmid {
				lo = mid
			} else {
				hi = mid
			}
		}
		if calc(lo, tTotal) < calc(hi, tTotal) {
			return hi
		}
		return lo
	}
	_ = findPeak

	findLowest := func(hi, tTotal, minDist int) int {
		lo := 0
		for lo < hi {
			mid := lo + (hi-lo)/2
			if calc(mid, tTotal) <= minDist {
				lo = mid + 1
			} else {
				hi = mid
			}
		}
		return lo
	}
	_ = findLowest

	findHighest := func(lo, tTotal, minDist int) int {
		hi := tTotal
		for lo < hi {
			mid := lo + (hi-lo)/2
			if calc(mid, tTotal) <= minDist {
				hi = mid
			} else {
				lo = mid + 1
			}
		}
		return lo - 1
	}
	_ = findHighest

	res := 1
	for i := 0; i < len(times); i++ {
		t := times[i]
		d := dists[i]
		peak := findPeak(t)
		lo := findLowest(peak, t, d)
		hi := findHighest(peak, t, d)
		delta := hi - lo + 1
		fmt.Printf("t: %v, d: %v, delta: %v\n", t, d, delta)
		res *= delta
	}
	// findPeak := func() {}

	return res
}

func part1(fn string) int {
	// This problem has two phases
	// In the first phase, we want to find the peak of the mountain range
	// I.e. the distance traveled will be monotonically increasing until its
	// peak, then descending until the result is once again, zero.
	// This gives us three binary searches - one where we are moving lo and hi
	// until there is a trio (lo, mid, hi) for which lo < mid > hi, then two
	// where we find the first position in the range [0,mid] and [mid,n] such
	// that it is either just inside or just outside the range of valid values.
	ls := lines(fn)
	timeStrs := strings.Fields(ls[0])[1:]
	distStrs := strings.Fields(ls[1])[1:]
	n := len(timeStrs)
	times := make([]int, n)
	dists := make([]int, n)
	for i := range timeStrs {
		times[i] = mustint(timeStrs[i])
		dists[i] = mustint(distStrs[i])
	}
	return solve(times, dists)
}

func part2(fn string) int {
	ls := lines(fn)
	times := []int{mustint(strings.Join(strings.Fields(ls[0])[1:], ""))}
	dists := []int{mustint(strings.Join(strings.Fields(ls[1])[1:], ""))}
	return solve(times, dists)
}

func main() {
	fmt.Println(part1("input"))
	fmt.Println(part2("input"))
}
