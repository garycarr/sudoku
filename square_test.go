package soduku

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewSquare(t *testing.T) {
	tt := []struct {
		description    string
		expectedSquare square
		input          [][]int
		pos            position
		printGrid      bool
	}{
		{
			description: "one",
			input: [][]int{
				[]int{0, 0, 0, 4, 5, 6, 7, 8, 9},
				[]int{0, 0, 0, 7, 8, 9, 1, 2, 3},
				[]int{0, 0, 0, 1, 2, 3, 4, 5, 6},
				[]int{0, 0, 0, 5, 6, 7, 8, 9, 1},
				[]int{0, 0, 0, 8, 9, 1, 2, 3, 4},
				[]int{0, 0, 0, 2, 3, 4, 5, 6, 7},
				[]int{0, 0, 0, 6, 7, 8, 9, 1, 2},
				[]int{0, 0, 0, 9, 1, 2, 3, 4, 5},
				[]int{0, 0, 0, 3, 4, 5, 6, 7, 8},
			},
			pos: position{
				rowNumber: 0,
				colNumber: 0,
			},
			expectedSquare: square{
				possibleNums: []int{1, 2, 3},
				reg: region{
					minColNumber: 0,
					minRowNumber: 0,
					maxColNumber: 2,
					maxRowNumber: 2,
				},
			},
		},
	}
	for _, td := range tt {
		t.Run(td.description, func(t *testing.T) {
			s, err := NewSquare(td.input, td.pos)
			require.Nil(t, err)
			sort.Ints(s.possibleNums)
			assert.Equal(t, &td.expectedSquare, s)
		})
	}
}

func TestGetRegion(t *testing.T) {
	tt := []struct {
		description    string
		pos            position
		expectedRegion region
	}{
		{
			description: "top left one",
			pos: position{
				rowNumber: 0,
				colNumber: 0,
			},
			expectedRegion: region{
				minRowNumber: 0,
				maxRowNumber: 2,
				minColNumber: 0,
				maxColNumber: 2,
			},
		},
		{
			description: "top left two",
			pos: position{
				rowNumber: 1,
				colNumber: 2,
			},
			expectedRegion: region{
				minRowNumber: 0,
				maxRowNumber: 2,
				minColNumber: 0,
				maxColNumber: 2,
			},
		},
		{
			description: "top right one",
			pos: position{
				rowNumber: 1,
				colNumber: 7,
			},
			expectedRegion: region{
				minRowNumber: 0,
				maxRowNumber: 2,
				minColNumber: 6,
				maxColNumber: 8,
			},
		},
		{
			description: "bottom right one",
			pos: position{
				rowNumber: 8,
				colNumber: 7,
			},
			expectedRegion: region{
				minRowNumber: 6,
				maxRowNumber: 8,
				minColNumber: 6,
				maxColNumber: 8,
			},
		},
		{
			description: "middle one",
			pos: position{
				rowNumber: 3,
				colNumber: 4,
			},
			expectedRegion: region{
				minRowNumber: 3,
				maxRowNumber: 5,
				minColNumber: 3,
				maxColNumber: 5,
			},
		},
	}
	for _, td := range tt {
		t.Run(td.description, func(t *testing.T) {
			s := square{
				pos: td.pos,
			}
			err := s.getRegion()
			require.Nil(t, err)
			assert.Equal(t, td.expectedRegion, s.reg)
		})
	}
}

func TestGetPossibleNumbers(t *testing.T) {
	tt := []struct {
		description    string
		expectedOutput []int
		input          [][]int
		pos            position
		printGrid      bool
	}{
		{
			description: "one",
			input: [][]int{
				[]int{0, 2, 3, 4, 5, 6, 7, 8, 9},
				[]int{4, 5, 6, 7, 8, 9, 1, 2, 3},
				[]int{7, 8, 9, 1, 2, 3, 4, 5, 6},
				[]int{2, 3, 4, 5, 6, 7, 8, 9, 1},
				[]int{5, 6, 7, 8, 9, 1, 2, 3, 4},
				[]int{8, 9, 1, 2, 3, 4, 5, 6, 7},
				[]int{3, 4, 5, 6, 7, 8, 9, 1, 2},
				[]int{6, 7, 8, 9, 1, 2, 3, 4, 5},
				[]int{9, 1, 2, 3, 4, 5, 6, 7, 8},
			},
			pos: position{
				rowNumber: 0,
				colNumber: 0,
			},
			expectedOutput: []int{1},
		},
		{
			description: "two",
			input: [][]int{
				[]int{0, 0, 3, 4, 5, 6, 7, 8, 9},
				[]int{4, 5, 6, 7, 8, 9, 1, 2, 3},
				[]int{7, 8, 9, 1, 2, 3, 4, 5, 6},
				[]int{2, 3, 4, 5, 6, 7, 8, 9, 1},
				[]int{5, 6, 7, 8, 9, 1, 2, 3, 4},
				[]int{8, 9, 1, 2, 3, 4, 5, 6, 7},
				[]int{3, 4, 5, 6, 7, 8, 9, 1, 2},
				[]int{6, 7, 8, 9, 1, 2, 3, 4, 5},
				[]int{9, 1, 2, 3, 4, 5, 6, 7, 8},
			},
			pos: position{
				rowNumber: 0,
				colNumber: 0,
			},
			expectedOutput: []int{1},
		},
		{
			description: "three",
			input: [][]int{
				[]int{0, 0, 3, 4, 5, 6, 7, 8, 9},
				[]int{0, 0, 6, 7, 8, 9, 1, 2, 3},
				[]int{7, 8, 9, 1, 2, 3, 4, 5, 6},
				[]int{2, 3, 4, 5, 6, 7, 8, 9, 1},
				[]int{5, 6, 7, 8, 9, 1, 2, 3, 4},
				[]int{8, 9, 1, 2, 3, 4, 5, 6, 7},
				[]int{3, 4, 5, 6, 7, 8, 9, 1, 2},
				[]int{6, 7, 8, 9, 1, 2, 3, 4, 5},
				[]int{9, 1, 2, 3, 4, 5, 6, 7, 8},
			},
			pos: position{
				rowNumber: 0,
				colNumber: 0,
			},
			expectedOutput: []int{1},
		},
		{
			description: "four",
			input: [][]int{
				[]int{0, 0, 0, 4, 5, 6, 7, 8, 9},
				[]int{0, 0, 0, 7, 8, 9, 1, 2, 3},
				[]int{0, 0, 0, 1, 2, 3, 4, 5, 6},
				[]int{2, 3, 4, 5, 6, 7, 8, 9, 1},
				[]int{5, 6, 7, 8, 9, 1, 2, 3, 4},
				[]int{8, 9, 1, 2, 3, 4, 5, 6, 7},
				[]int{3, 4, 5, 6, 7, 8, 9, 1, 2},
				[]int{6, 7, 8, 9, 1, 2, 3, 4, 5},
				[]int{9, 1, 2, 3, 4, 5, 6, 7, 8},
			},
			pos: position{
				rowNumber: 0,
				colNumber: 0,
			},
			expectedOutput: []int{1},
		},
		{
			description: "five",
			input: [][]int{
				[]int{0, 0, 0, 4, 5, 6, 7, 8, 9},
				[]int{0, 0, 0, 7, 8, 9, 1, 2, 3},
				[]int{0, 0, 0, 1, 2, 3, 4, 5, 6},
				[]int{0, 0, 0, 5, 6, 7, 8, 9, 1},
				[]int{0, 0, 0, 8, 9, 1, 2, 3, 4},
				[]int{0, 0, 0, 2, 3, 4, 5, 6, 7},
				[]int{0, 0, 0, 6, 7, 8, 9, 1, 2},
				[]int{0, 0, 0, 9, 1, 2, 3, 4, 5},
				[]int{0, 0, 0, 3, 4, 5, 6, 7, 8},
			},
			pos: position{
				rowNumber: 0,
				colNumber: 0,
			},
			expectedOutput: []int{1, 2, 3},
		},
	}

	for _, td := range tt {
		t.Run(td.description, func(t *testing.T) {
			s := &square{
				pos: td.pos,
			}
			err := s.getRegion()
			assert.Nil(t, err)
			assert.Nil(t, s.getPossibleNumbers(td.input))
			sort.Ints(s.possibleNums)
			assert.Equal(t, td.expectedOutput, s.possibleNums)

			if td.printGrid {
				PrintGrid(td.input)
			}
		})
	}
}

func TestGetEmptySquares(t *testing.T) {
	tt := []struct {
		description string
		input       [][]int
		expPos      []position
		printGrid   bool
	}{
		{
			description: "one",
			input: [][]int{
				[]int{0, 2, 3, 4, 5, 6, 7, 8, 9},
				[]int{4, 5, 6, 7, 8, 9, 1, 2, 3},
				[]int{7, 8, 9, 1, 2, 3, 4, 5, 6},
				[]int{2, 3, 4, 5, 6, 7, 8, 9, 1},
				[]int{5, 6, 7, 8, 9, 1, 2, 3, 4},
				[]int{8, 9, 1, 2, 3, 4, 5, 6, 7},
				[]int{3, 4, 5, 6, 7, 8, 9, 1, 2},
				[]int{6, 7, 8, 9, 1, 2, 3, 4, 5},
				[]int{9, 1, 2, 3, 4, 5, 6, 7, 8},
			},
			expPos: []position{
				{
					rowNumber: 0,
					colNumber: 0,
				},
			},
		},
		{
			description: "two",
			input: [][]int{
				[]int{0, 0, 3, 4, 5, 6, 7, 8, 9},
				[]int{0, 0, 6, 7, 8, 9, 1, 2, 3},
				[]int{7, 8, 9, 1, 2, 3, 4, 5, 6},
				[]int{2, 3, 4, 5, 6, 7, 8, 9, 1},
				[]int{5, 6, 7, 8, 9, 1, 2, 3, 4},
				[]int{8, 9, 1, 2, 3, 4, 5, 6, 7},
				[]int{3, 4, 5, 6, 7, 8, 9, 1, 2},
				[]int{6, 7, 8, 9, 1, 2, 3, 4, 5},
				[]int{9, 1, 2, 3, 4, 5, 6, 7, 8},
			},
			expPos: []position{
				{
					rowNumber: 0,
					colNumber: 0,
				},
				{
					rowNumber: 0,
					colNumber: 1,
				},
				{
					rowNumber: 1,
					colNumber: 0,
				},
				{
					rowNumber: 1,
					colNumber: 1,
				},
			},
		},
	}

	for _, td := range tt {
		t.Run(td.description, func(t *testing.T) {
			assert.Equal(t, getEmptySquares(td.input), td.expPos)
			if td.printGrid {
				PrintGrid(td.input)
			}
		})
	}
}
