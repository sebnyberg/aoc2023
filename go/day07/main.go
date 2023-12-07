package main

import (
	"fmt"
	"io"
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

type hand struct {
	cards string
	bid   int
}

const (
	highCard     = 1
	onePair      = 2
	twoPair      = 3
	threeOfAKind = 4
	fullHouse    = 5
	fourOfAKind  = 6
	fiveOfAKind  = 7
)

func part1(fn string) int {
	var cardval = [...]int{
		'1': 1,
		'2': 2,
		'3': 3,
		'4': 4,
		'5': 5,
		'6': 6,
		'7': 7,
		'8': 8,
		'9': 9,
		'T': 10,
		'J': 11,
		'Q': 12,
		'K': 13,
		'A': 14,
	}
	value := func(h hand) int {
		c := []byte(h.cards)
		sort.Slice(c, func(i, j int) bool {
			return c[i] < c[j]
		})
		var count [256]int
		var maxCount int
		var pairCount int
		for i := range c {
			count[c[i]]++
			if count[c[i]] == 2 {
				pairCount++
			}
			maxCount = max(maxCount, count[c[i]])
		}
		switch {
		case maxCount == 5:
			return fiveOfAKind
		case maxCount == 4:
			return fourOfAKind
		case maxCount == 3 && pairCount == 2:
			return fullHouse
		case maxCount == 3:
			return threeOfAKind
		case pairCount == 2:
			return twoPair
		case pairCount == 1:
			return onePair
		}
		return highCard
	}
	// We should be able to simply sort all hands then calculate the result.
	//
	// To make comparisons of the same hand easier, sort each hand by group size
	// first, then card value.
	ls := lines(fn)
	n := len(ls)
	hands := make([]hand, n)
	for i, l := range ls {
		p := strings.Fields(l)
		bs := []byte(p[0])
		hands[i] = hand{string(bs), atoi(p[1])}
	}
	less := func(a, b hand) bool {
		av := value(a)
		bv := value(b)
		if av != bv {
			return av < bv
		}
		aa := a.cards
		bb := b.cards
		for i := range aa {
			if aa[i] != bb[i] {
				return cardval[aa[i]] < cardval[bb[i]]
			}
		}
		return true // equal
	}
	sort.Slice(hands, func(i, j int) bool {
		return less(hands[i], hands[j])
	})
	var res int
	for i := range hands {
		res += (i + 1) * hands[i].bid
	}

	return res
}

func part2(fn string) int {
	var cardval = [...]int{
		'J': 0,
		'1': 1,
		'2': 2,
		'3': 3,
		'4': 4,
		'5': 5,
		'6': 6,
		'7': 7,
		'8': 8,
		'9': 9,
		'T': 10,
		'Q': 12,
		'K': 13,
		'A': 14,
	}
	value := func(h hand) int {
		// With jokers, things are a little bit more difficult
		//
		// Luckily, the best way to increase the total value will always be to
		// increase the count of the biggest group.
		//
		// So we count all non-joker cards first, then we assign its face to be
		// that of the most prevalent card.
		//
		c := []byte(h.cards)
		sort.Slice(c, func(i, j int) bool {
			return c[i] < c[j]
		})
		var count [256]int
		var maxCount int
		var pairCount int
		var jokers int
		for i := range c {
			if c[i] == 'J' {
				jokers++
				continue
			}
			count[c[i]]++
			if count[c[i]] == 2 {
				pairCount++
			}
			maxCount = max(maxCount, count[c[i]])
		}
		if jokers > 0 {
			if maxCount == 1 {
				pairCount++
			}
			maxCount += jokers
		}
		switch {
		case maxCount == 5:
			return fiveOfAKind
		case maxCount == 4:
			return fourOfAKind
		case maxCount == 3 && pairCount == 2:
			return fullHouse
		case maxCount == 3:
			return threeOfAKind
		case pairCount == 2:
			return twoPair
		case pairCount == 1:
			return onePair
		}
		return highCard
	}
	// We should be able to simply sort all hands then calculate the result.
	//
	// To make comparisons of the same hand easier, sort each hand by group size
	// first, then card value.
	ls := lines(fn)
	n := len(ls)
	hands := make([]hand, n)
	for i, l := range ls {
		p := strings.Fields(l)
		bs := []byte(p[0])
		hands[i] = hand{string(bs), atoi(p[1])}
	}
	less := func(a, b hand) bool {
		av := value(a)
		bv := value(b)
		if av != bv {
			return av < bv
		}
		aa := a.cards
		bb := b.cards
		for i := range aa {
			if aa[i] != bb[i] {
				return cardval[aa[i]] < cardval[bb[i]]
			}
		}
		return true // equal
	}
	sort.Slice(hands, func(i, j int) bool {
		return less(hands[i], hands[j])
	})
	var res int
	for i := range hands {
		res += (i + 1) * hands[i].bid
	}

	return res
}

func main() {
	fmt.Println(part1("input"))
	fmt.Println(part2("input"))
}
