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

func readlines(fname string) []string {
	f, err := os.Open(fname)
	maybe(err)
	contents, err := io.ReadAll(f)
	maybe(err)
	res := strings.Split(string(contents), "\n")
	for res[len(res)-1] == "" {
		// Skip any empty line
		res = res[:len(res)-1]
	}
	return res
}

type movement [2]int

func (m movement) String() string {
	switch m {
	case goRight:
		return "right"
	case goLeft:
		return "left"
	case goUp:
		return "up"
	case goDown:
		return "down"
	case goDownLeft:
		return "down-left"
	case goDownRight:
		return "down-right"
	case goUpLeft:
		return "up-left"
	case goUpRight:
		return "up-right"
	default:
		return "unknown"
	}
}

func (m movement) less(other movement) bool {
	if m[0] == other[0] {
		return m[1] < other[1]
	}
	return m[0] < other[0]
}

var goRight = movement{0, 1}
var goLeft = movement{0, -1}
var goUp = movement{-1, 0}
var goDown = movement{1, 0}
var goDownLeft = movement{1, -1}
var goDownRight = movement{1, 1}
var goUpRight = movement{-1, 1}
var goUpLeft = movement{-1, -1}
var dirs = []movement{goUp, goRight, goDown, goLeft}
var dirIdx = map[movement]int{
	goUp:    0,
	goRight: 1,
	goDown:  2,
	goLeft:  3,
}

var movements = [...][]movement{
	'F': {goRight, goDown},
	'J': {goLeft, goUp},
	'L': {goUp, goRight},
	'7': {goDown, goLeft},
	'|': {goUp, goDown},
	'-': {goLeft, goRight},
}

func hasMvt(ch byte, dir movement) bool {
	if ch == '.' {
		return false
	}
	for _, mvt := range movements[ch] {
		if mvt == dir {
			return true
		}
	}
	return false
}

func findStartTile(grid []string) (int, int) {
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 'S' {
				return i, j
			}
		}
	}
	return -1, -1
}

func replaceStartTileInplace(grid []string, si, sj int) {
	m := len(grid)
	n := len(grid[0])
	firstMvts := []movement{}
	for j, d := range dirs {
		ii := si + d[0]
		jj := sj + d[1]
		if ii < 0 || jj < 0 || ii >= m || jj >= n {
			// out of bounds
			continue
		}
		if hasMvt(grid[ii][jj], dirs[(j+2)%4]) {
			firstMvts = append(firstMvts, d)
		}
	}
	if len(firstMvts) != 2 {
		panic("invalid first tile")
	}
	sort.Slice(firstMvts, func(i, j int) bool {
		return firstMvts[i].less(firstMvts[j])
	})
	var startTile byte
	for _, ch := range "LJF7|-" {
		mvts := movements[ch]
		sort.Slice(mvts, func(i, j int) bool {
			return mvts[i].less(mvts[j])
		})
		if mvts[0] == firstMvts[0] && mvts[1] == firstMvts[1] {
			startTile = byte(ch)
			break
		}
	}
	// Adjust puzzle accordingly
	grid[si] = fmt.Sprintf("%v%c%v", grid[si][:sj], startTile, grid[si][sj+1:])
}

func part1(lines []string) int {
	// This is a classical flood-fill / BFS problem. Starting with 'S', visit
	// neighbours in each valid direction and mark unvisited nodes as seen
	//
	m := len(lines)
	n := len(lines[0])

	si, sj := findStartTile(lines)
	replaceStartTileInplace(lines, si, sj)

	seen := make([][]bool, m)
	for i := range seen {
		seen[i] = make([]bool, n)
	}
	seen[si][sj] = true
	curr := [][]int{{si, sj}}
	next := [][]int{}
	var steps int
	for len(curr) > 0 {
		next = next[:0]
		for _, x := range curr {
			i := x[0]
			j := x[1]
			ch := lines[i][j]
			for _, d := range movements[ch] {
				ii := i + d[0]
				jj := j + d[1]
				// Out of bounds not possible unless puzzle is malformed
				if seen[ii][jj] {
					continue
				}
				seen[ii][jj] = true
				next = append(next, []int{ii, jj})
			}
		}
		steps++
		curr, next = next, curr
	}

	return steps - 1
}

func part2(lines []string) int {
	// The second part is a bit different.
	//
	// If we travel through the loop and count its turns, then the total turns
	// must either end up taking four extra right or four extra left turns.
	//
	// If there are four extra right turns, then we should fill any space on the
	// left side of any location in the loop, and vice versa.
	//
	// To do this, we pick a direction, travel through the loop one round,
	// counting turns. Then depending on left/right, we will travel through the
	// loop one more time, filling left/right neighbourd depending on the
	// outcome of the previous loop.
	//
	m := len(lines)
	n := len(lines[0])

	si, sj := findStartTile(lines)
	replaceStartTileInplace(lines, si, sj)

	// Convert to byte
	grid := make([][]byte, m)
	for i := range grid {
		grid[i] = []byte(lines[i])
	}
	lines = nil // lines should not be used after this point

	// Move through the loop until we know what side to go for
	var leftRightDelta int
	step := func(pos [2]int, dir movement) (nextPos [2]int, nextDir movement) {
		nextPos[0] = pos[0] + dir[0]
		nextPos[1] = pos[1] + dir[1]
		oppositeDir := dirs[(dirIdx[dir]+2)%4]
		nextMvts := movements[grid[nextPos[0]][nextPos[1]]]
		if nextMvts[0] == oppositeDir {
			nextDir = nextMvts[1]
		} else {
			nextDir = nextMvts[0]
		}
		return nextPos, nextDir
	}

	// Pick any of the two directions from the starting point.
	// Then move through the loop, counting turns
	initialDir := movements[grid[si][sj]][0]
	dir := initialDir
	pos := [2]int{si, sj}
	isloop := make([][]bool, m)
	for i := range isloop {
		isloop[i] = make([]bool, n)
	}
	isloop[si][sj] = true

	for {
		nextPos, nextDir := step(pos, dir)
		isloop[nextPos[0]][nextPos[1]] = true

		var delta int
		if (dirIdx[dir]+1)%4 == dirIdx[nextDir] {
			delta = 1 // right turn
		} else if (dirIdx[dir]+3)%4 == dirIdx[nextDir] {
			delta = -1 // left turn
		} else {
			delta = 0 // no change, included for clarity
		}
		leftRightDelta += delta

		// fmt.Printf(
		// 	"pos: (%v, %v), curDir: %v, nextDir: %v, delta: %v, deltaTot: %v\n",
		// 	pos[0], pos[1], dir, nextDir, delta, leftRightDelta,
		// )
		pos, dir = nextPos, nextDir
		if pos == [2]int{si, sj} {
			break
		}
	}

	var fillCount int
	curr := [][]int{}
	next := [][]int{}
	fill := func(si, sj int) {
		if si < 0 || si >= m || sj < 0 || sj >= n || isloop[si][sj] || grid[si][sj] == 'I' {
			return
		}
		// fmt.Printf("filling (%v,%v)\n", si, sj)
		curr = curr[:0]
		fillCount++
		grid[si][sj] = 'I'
		curr = append(curr, []int{si, sj})
		for len(curr) > 0 {
			next = next[:0]
			for _, x := range curr {
				i := x[0]
				j := x[1]
				for _, d := range dirs {
					ii := i + d[0]
					jj := j + d[1]
					if ii < 0 || jj < 0 || ii >= m || jj >= n || isloop[ii][jj] || grid[ii][jj] == 'I' {
						continue
					}
					// fmt.Printf("filling (%v,%v)\n", ii, jj)
					grid[ii][jj] = 'I'
					fillCount++
					next = append(next, []int{ii, jj})
				}
			}
			curr, next = next, curr
		}
	}

	// If leftRightDelta == -4, then we should keep track of the left side.
	// Otherwise, we keep track of the right side
	//
	// My intuition is telling me that we only need to consider straight pipes
	// when filling, but it doesn't hurt to always fill on the side.
	//
	type key struct {
		tile   byte
		outDir movement
	}
	dir = initialDir
	pos = [2]int{si, sj}
	if leftRightDelta == -4 {
		// Flip direction so that we always fill the right side of the track
		mvts := movements[grid[pos[0]][pos[1]]]
		if mvts[0] == initialDir {
			dir = mvts[1]
		} else {
			dir = mvts[0]
		}
	}
	// Right side of current direction should be filled with 'I's
	fillMovement := map[key][]movement{
		{'|', goUp}:    {goRight},
		{'|', goDown}:  {goLeft},
		{'J', goUp}:    {goDown, goDownRight, goRight},
		{'J', goLeft}:  {goUpLeft},
		{'-', goLeft}:  {goUp},
		{'-', goRight}: {goDown},
		{'L', goUp}:    {goUpRight},
		{'L', goRight}: {goLeft, goDownLeft, goDown},
		{'F', goRight}: {goDownRight},
		{'F', goDown}:  {goUp, goUpLeft, goLeft},
		{'7', goLeft}:  {goUp, goUpRight, goRight},
		{'7', goDown}:  {goDownLeft},
	}
	fmt.Println(dir)

	// Reset position / direction
	for {
		// Fill side
		k := key{
			grid[pos[0]][pos[1]],
			dir,
		}
		fillMvts := fillMovement[k]
		// fmt.Printf(
		// 	"pos: (%v, %v), sym: %c, dir: %v, fillMvt: %v, fillCount: %v\n",
		// 	pos[0], pos[1], grid[pos[0]][pos[1]], dir, fillMvts, fillCount,
		// )
		for _, mvt := range fillMvts {
			ii := pos[0] + mvt[0]
			jj := pos[1] + mvt[1]
			if ii < 0 || jj < 0 || ii >= m || jj >= n || grid[ii][jj] == 'I' {
				continue
			}
			fill(ii, jj)
		}
		pos, dir = step(pos, dir)
		if pos == [2]int{si, sj} {
			break
		}
	}

	return fillCount
}

func main() {
	// fmt.Println(part1(readlines("input")))
	fmt.Println(part2(readlines("input")))
}
