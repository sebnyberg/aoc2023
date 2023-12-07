package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func maybe(err error) {
	if err != nil {
		panic(err)
	}
}

func lines(fname string) []string {
	f, err := os.Open(fname)
	maybe(err)
	contents, err := io.ReadAll(f)
	maybe(err)
	return strings.Split(string(contents), "\n")
}

func part1(fname string) int {
	var res int
	for _, l := range lines(fname) {
		first := -1
		last := -1
		for _, ch := range l {
			if ch >= '0' && ch <= '9' {
				if first == -1 {
					first = int(ch - '0')
				}
				last = int(ch - '0')
			}
		}
		if first != -1 {
			res += first*10 + last
		}
	}
	return res
}

func part2(fname string) int {
	var res int
	for _, l := range lines(fname) {
		first := -1
		last := -1
		update := func(x int) {
			if first == -1 {
				first = x
			}
			last = x
		}
		for i := range l {
			if ch := l[i]; ch >= '0' && ch <= '9' {
				update(int(ch - '0'))
				continue
			}

			for j, s := range []string{
				"one", "two", "three", "four", "five",
				"six", "seven", "eight", "nine",
			} {
				if len(s) > len(l)-i {
					continue
				}
				if l[i:i+len(s)] == s {
					update(j + 1)
					i += len(s) - 1
				}
			}
		}
		if first != -1 {
			res += first*10 + last
		}
	}
	return res
}

func main() {
	fmt.Println(part1("input"))
	fmt.Println(part2("input"))
}
