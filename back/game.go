package back

import (
	"math/rand"
	"time"
)

type State struct {
	Width  int
	Height int
	Cells  [][]bool
}

func NewState(width, height int) *State {
	cells := make([][]bool, width)
	for i := range cells {
		cells[i] = make([]bool, height)
	}
	return &State{Width: width, Height: height, Cells: cells}
}

func (s *State) Set(x, y int, alive bool) {
	s.Cells[x][y] = alive
}

func (s *State) Get(x, y int) bool {
	return s.Cells[x][y]
}

func (s *State) Update() *State {
	next := NewState(s.Width, s.Height)

	for x := 0; x < s.Width; x++ {
		for y := 0; y < s.Height; y++ {
			neighbors := s.countAliveNeighbors(x, y)

			if s.Get(x, y) {
				// Any live cell with fewer than two live neighbors dies, as if by underpopulation.
				// Any live cell with two or three live neighbors lives on to the next generation.
				// Any live cell with more than three live neighbors dies, as if by overpopulation.
				next.Set(x, y, neighbors == 2 || neighbors == 3)
			} else {
				// Any dead cell with exactly three live neighbors becomes a live cell, as if by reproduction.
				next.Set(x, y, neighbors == 3)
			}
		}
	}

	return next
}

func (s *State) countAliveNeighbors(x, y int) int {
	count := 0

	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}
			x2 := x + i
			y2 := y + j
			if x2 < 0 || x2 >= s.Width || y2 < 0 || y2 >= s.Height {
				continue
			}
			if s.Get(x2, y2) {
				count++
			}
		}
	}

	return count
}

func (s *State) Randomize() {
	rand.Seed(time.Now().UnixNano())
	for x := 0; x < s.Width; x++ {
		for y := 0; y < s.Height; y++ {
			if rand.Float64() < 0.5 {
				s.Set(x, y, true)
			} else {
				s.Set(x, y, false)
			}
		}
	}
}
