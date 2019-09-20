package soduku

import (
	"fmt"
)

// square is used to return data on a given square in the grid
type square struct {
	pos          position
	possibleNums []int
	reg          region
}

type position struct {
	rowNumber int
	colNumber int
}

type region struct {
	minColNumber int
	maxColNumber int
	minRowNumber int
	maxRowNumber int
}

// allRegions are the 9 sections on the grid to make between 1 and 9
var allRegions = []region{
	region{minRowNumber: 0, maxRowNumber: 2, minColNumber: 0, maxColNumber: 2},
	region{minRowNumber: 0, maxRowNumber: 2, minColNumber: 3, maxColNumber: 5},
	region{minRowNumber: 0, maxRowNumber: 2, minColNumber: 6, maxColNumber: 8},
	region{minRowNumber: 3, maxRowNumber: 5, minColNumber: 0, maxColNumber: 2},
	region{minRowNumber: 3, maxRowNumber: 5, minColNumber: 3, maxColNumber: 5},
	region{minRowNumber: 3, maxRowNumber: 5, minColNumber: 6, maxColNumber: 8},
	region{minRowNumber: 6, maxRowNumber: 8, minColNumber: 0, maxColNumber: 2},
	region{minRowNumber: 6, maxRowNumber: 8, minColNumber: 3, maxColNumber: 5},
	region{minRowNumber: 6, maxRowNumber: 8, minColNumber: 6, maxColNumber: 8},
}

func NewSquares(grid [][]int) ([]*square, error) {
	ss := []*square{}
	poss := getEmptySquares(grid)
	for _, pos := range poss {
		s, err := NewSquare(grid, pos)
		if err != nil {
			return ss, err
		}
		ss = append(ss, s)
	}
	return ss, nil
}

func NewSquare(grid [][]int, pos position) (*square, error) {
	s := &square{
		pos: pos,
	}
	err := s.getRegion()
	if err != nil {
		return s, err
	}

	if err := s.getPossibleNumbers(grid); err != nil {
		return s, err
	}
	return s, nil
}

// getEmptySquares returns the positions of empty squares
func getEmptySquares(grid [][]int) []position {
	pos := []position{}

	// identify missing positions
	for rowNumber, row := range grid {
		for colNumber, num := range row {
			if num > 0 {
				continue
			}
			pos = append(pos, position{
				rowNumber: rowNumber,
				colNumber: colNumber,
			})
		}
	}
	return pos
}

// possibleNumbers returns the numbers that can possibly placed into a given position
func (s *square) getPossibleNumbers(grid [][]int) error {
	s.possibleNums = []int{}
	possibleNumbers := map[int]bool{}
	for i := 1; i <= 9; i++ {
		possibleNumbers[i] = true
	}

	// check the row it is on
	for col := 0; col <= 8; col++ {
		possibleNumbers[grid[s.pos.rowNumber][col]] = false
	}

	// check the column it is in
	for row := 0; row <= 8; row++ {
		possibleNumbers[grid[row][s.pos.colNumber]] = false
	}

	// Check the grid it is in
	for row := s.reg.minRowNumber; row <= s.reg.maxRowNumber; row++ {
		for col := s.reg.minColNumber; col <= s.reg.maxColNumber; col++ {
			possibleNumbers[grid[row][col]] = false
		}
	}

	for num, stillPossible := range possibleNumbers {
		if stillPossible {
			s.possibleNums = append(s.possibleNums, num)
		}
	}
	return nil
}

// getRegion returns he grid position that the position is in
func (s *square) getRegion() error {
	reg := region{}
	switch {
	case s.pos.rowNumber >= 0 && s.pos.rowNumber <= 2:
		reg.minRowNumber = 0
		reg.maxRowNumber = 2
	case s.pos.rowNumber >= 3 && s.pos.rowNumber <= 5:
		reg.minRowNumber = 3
		reg.maxRowNumber = 5
	case s.pos.rowNumber >= 6 && s.pos.rowNumber <= 8:
		reg.minRowNumber = 6
		reg.maxRowNumber = 8
	default:
		return fmt.Errorf("rowNumber %d is invalid", s.pos.rowNumber)
	}

	switch {
	case s.pos.colNumber >= 0 && s.pos.colNumber <= 2:
		reg.minColNumber = 0
		reg.maxColNumber = 2
	case s.pos.colNumber >= 3 && s.pos.colNumber <= 5:
		reg.minColNumber = 3
		reg.maxColNumber = 5
	case s.pos.colNumber >= 6 && s.pos.colNumber <= 8:
		reg.minColNumber = 6
		reg.maxColNumber = 8
	default:
		return fmt.Errorf("colNumber %d is invalid", s.pos.colNumber)
	}
	s.reg = reg
	return nil
}
