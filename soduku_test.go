package soduku

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSolveGridSimple(t *testing.T) {
	tt := []struct {
		description string
		printGrid   bool
		input       [][]int
	}{
		{
			description: "zero elements missing",
			input: [][]int{
				[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
				[]int{4, 5, 6, 7, 8, 9, 1, 2, 3},
				[]int{7, 8, 9, 1, 2, 3, 4, 5, 6},
				[]int{2, 3, 4, 5, 6, 7, 8, 9, 1},
				[]int{5, 6, 7, 8, 9, 1, 2, 3, 4},
				[]int{8, 9, 1, 2, 3, 4, 5, 6, 7},
				[]int{3, 4, 5, 6, 7, 8, 9, 1, 2},
				[]int{6, 7, 8, 9, 1, 2, 3, 4, 5},
				[]int{9, 1, 2, 3, 4, 5, 6, 7, 8},
			},
		},
		{
			description: "one elements missing one",
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
		},
		{
			description: "one elements missing two",
			input: [][]int{
				[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
				[]int{4, 5, 6, 7, 8, 9, 1, 2, 3},
				[]int{7, 8, 9, 1, 2, 3, 4, 5, 6},
				[]int{2, 3, 4, 5, 6, 7, 8, 9, 1},
				[]int{5, 6, 7, 8, 9, 1, 2, 3, 4},
				[]int{8, 9, 1, 2, 3, 4, 5, 6, 7},
				[]int{3, 4, 5, 6, 7, 8, 9, 1, 2},
				[]int{6, 7, 8, 9, 1, 2, 3, 4, 5},
				[]int{9, 1, 2, 3, 4, 5, 6, 7, 0},
			},
		},
		{
			description: "two elements missing row same box one",
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
		},
		{
			description: "two elements missing row same box row two",
			input: [][]int{
				[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
				[]int{4, 5, 6, 7, 8, 9, 1, 2, 3},
				[]int{7, 8, 9, 1, 2, 3, 4, 5, 6},
				[]int{2, 3, 4, 5, 6, 7, 8, 9, 1},
				[]int{5, 6, 7, 8, 9, 1, 2, 3, 4},
				[]int{8, 9, 1, 2, 3, 4, 5, 6, 7},
				[]int{3, 4, 5, 6, 7, 8, 9, 1, 2},
				[]int{6, 7, 8, 9, 1, 2, 3, 4, 5},
				[]int{9, 1, 2, 3, 4, 5, 6, 0, 0},
			},
		},
		{
			description: "two elements missing different box row one",
			input: [][]int{
				[]int{0, 2, 3, 4, 5, 6, 7, 8, 9},
				[]int{4, 5, 6, 7, 8, 9, 1, 2, 3},
				[]int{7, 8, 9, 1, 2, 3, 4, 5, 6},
				[]int{2, 3, 4, 5, 6, 7, 8, 9, 1},
				[]int{5, 6, 7, 8, 9, 1, 2, 3, 4},
				[]int{8, 9, 1, 2, 3, 4, 5, 6, 7},
				[]int{3, 4, 5, 6, 7, 8, 9, 1, 2},
				[]int{6, 7, 8, 9, 1, 2, 3, 4, 5},
				[]int{9, 1, 2, 3, 4, 5, 6, 7, 0},
			},
		},
		{
			description: "two elements missing same box col one",
			input: [][]int{
				[]int{0, 2, 3, 4, 5, 6, 7, 8, 9},
				[]int{0, 5, 6, 7, 8, 9, 1, 2, 3},
				[]int{7, 8, 9, 1, 2, 3, 4, 5, 6},
				[]int{2, 3, 4, 5, 6, 7, 8, 9, 1},
				[]int{5, 6, 7, 8, 9, 1, 2, 3, 4},
				[]int{8, 9, 1, 2, 3, 4, 5, 6, 7},
				[]int{3, 4, 5, 6, 7, 8, 9, 1, 2},
				[]int{6, 7, 8, 9, 1, 2, 3, 4, 5},
				[]int{9, 1, 2, 3, 4, 5, 6, 7, 8},
			},
		},
		{
			description: "two elements missing same box col two",
			input: [][]int{
				[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
				[]int{4, 5, 6, 7, 8, 9, 1, 2, 3},
				[]int{7, 8, 9, 1, 2, 3, 4, 5, 6},
				[]int{2, 3, 4, 5, 6, 7, 8, 9, 1},
				[]int{5, 6, 7, 8, 9, 1, 2, 3, 4},
				[]int{8, 9, 1, 2, 3, 4, 5, 6, 7},
				[]int{3, 4, 5, 6, 7, 8, 9, 1, 2},
				[]int{6, 7, 8, 9, 1, 2, 3, 4, 0},
				[]int{9, 1, 2, 3, 4, 5, 6, 7, 0},
			},
		},
		{
			description: "two elements missing different box col one",
			input: [][]int{
				[]int{1, 2, 3, 4, 5, 6, 7, 8, 0},
				[]int{4, 5, 6, 7, 8, 9, 1, 2, 3},
				[]int{7, 8, 9, 1, 2, 3, 4, 5, 6},
				[]int{2, 3, 4, 5, 6, 7, 8, 9, 1},
				[]int{5, 6, 7, 8, 9, 1, 2, 3, 4},
				[]int{8, 9, 1, 2, 3, 4, 5, 6, 7},
				[]int{3, 4, 5, 6, 7, 8, 9, 1, 2},
				[]int{6, 7, 8, 9, 1, 2, 3, 4, 5},
				[]int{9, 1, 2, 3, 4, 5, 6, 7, 0},
			},
		},
		{
			description: "three elements missing same box row one",
			input: [][]int{
				[]int{0, 0, 0, 4, 5, 6, 7, 8, 9},
				[]int{4, 5, 6, 7, 8, 9, 1, 2, 3},
				[]int{7, 8, 9, 1, 2, 3, 4, 5, 6},
				[]int{2, 3, 4, 5, 6, 7, 8, 9, 1},
				[]int{5, 6, 7, 8, 9, 1, 2, 3, 4},
				[]int{8, 9, 1, 2, 3, 4, 5, 6, 7},
				[]int{3, 4, 5, 6, 7, 8, 9, 1, 2},
				[]int{6, 7, 8, 9, 1, 2, 3, 4, 5},
				[]int{9, 1, 2, 3, 4, 5, 6, 7, 0},
			},
		},
		{
			description: "three elements missing same box row two",
			input: [][]int{
				[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
				[]int{4, 5, 6, 7, 8, 9, 1, 2, 3},
				[]int{7, 8, 9, 1, 2, 3, 4, 5, 6},
				[]int{2, 3, 4, 5, 6, 7, 8, 9, 1},
				[]int{5, 6, 7, 8, 9, 1, 2, 3, 4},
				[]int{8, 9, 1, 2, 3, 4, 5, 6, 7},
				[]int{3, 4, 5, 6, 7, 8, 9, 1, 2},
				[]int{6, 7, 8, 9, 1, 2, 3, 4, 5},
				[]int{9, 1, 2, 3, 4, 5, 0, 0, 0},
			},
		},
		{
			description: "three elements missing different box row one",
			input: [][]int{
				[]int{0, 2, 3, 4, 5, 6, 7, 8, 9},
				[]int{4, 5, 6, 7, 8, 9, 1, 2, 3},
				[]int{7, 8, 9, 0, 2, 3, 4, 5, 6},
				[]int{2, 3, 4, 5, 6, 7, 8, 9, 1},
				[]int{5, 6, 7, 8, 9, 1, 2, 3, 4},
				[]int{8, 9, 1, 2, 3, 4, 5, 6, 7},
				[]int{3, 4, 5, 6, 7, 0, 9, 1, 2},
				[]int{6, 7, 8, 9, 1, 2, 3, 4, 5},
				[]int{9, 1, 2, 3, 4, 5, 6, 7, 0},
			},
		},
		{
			description: "three elements missing same box col one",
			input: [][]int{
				[]int{0, 2, 3, 4, 5, 6, 7, 8, 9},
				[]int{0, 5, 6, 7, 8, 9, 1, 2, 3},
				[]int{0, 8, 9, 1, 2, 3, 4, 5, 6},
				[]int{2, 3, 4, 5, 6, 7, 8, 9, 1},
				[]int{5, 6, 7, 8, 9, 1, 2, 3, 4},
				[]int{8, 9, 1, 2, 3, 4, 5, 6, 7},
				[]int{3, 4, 5, 6, 7, 8, 9, 1, 2},
				[]int{6, 7, 8, 9, 1, 2, 3, 4, 5},
				[]int{9, 1, 2, 3, 4, 5, 6, 7, 0},
			},
		},
		{
			description: "three elements missing different box col one",
			input: [][]int{
				[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
				[]int{4, 5, 6, 7, 8, 9, 1, 2, 3},
				[]int{0, 8, 9, 1, 2, 3, 4, 5, 6},
				[]int{2, 3, 4, 5, 6, 7, 8, 9, 1},
				[]int{5, 6, 7, 8, 9, 1, 2, 3, 4},
				[]int{0, 9, 1, 2, 3, 4, 5, 6, 7},
				[]int{3, 4, 5, 6, 7, 8, 9, 1, 2},
				[]int{6, 7, 8, 9, 1, 2, 3, 4, 5},
				[]int{0, 1, 2, 3, 4, 5, 6, 7, 0},
			},
		},
		{
			description: "nine elements missing same box",
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
		},
		{
			description: "nine elements missing same number",
			input: [][]int{
				[]int{0, 2, 3, 4, 5, 6, 7, 8, 9},
				[]int{4, 5, 6, 7, 8, 9, 0, 2, 3},
				[]int{7, 8, 9, 0, 2, 3, 4, 5, 6},
				[]int{2, 3, 4, 5, 6, 7, 8, 9, 0},
				[]int{5, 6, 7, 8, 9, 0, 2, 3, 4},
				[]int{8, 9, 0, 2, 3, 4, 5, 6, 7},
				[]int{3, 4, 5, 6, 7, 8, 9, 0, 2},
				[]int{6, 7, 8, 9, 0, 2, 3, 4, 5},
				[]int{9, 0, 2, 3, 4, 5, 6, 7, 8},
			},
		},
	}

	for _, td := range tt {
		t.Run(td.description, func(t *testing.T) {
			output, cg, err := SolveGrid(td.input)
			require.Nil(t, err)
			assert.Equal(t, CheckedGrid{Valid: true, Complete: true}, cg)

			if td.printGrid {
				PrintGrid(output)
			}
		})
	}
}

func TestSolveGridMedium(t *testing.T) {
	tt := []struct {
		description    string
		printGrid      bool
		input          [][]int
		expectComplete bool
		expectOutput   [][]int
	}{
		{
			description: "one",
			input: [][]int{
				[]int{0, 0, 0, 0, 5, 6, 7, 8, 9},
				[]int{0, 0, 0, 0, 8, 9, 0, 2, 3},
				[]int{0, 0, 0, 0, 2, 3, 4, 5, 6},
				[]int{2, 3, 4, 5, 6, 7, 8, 9, 0},
				[]int{5, 6, 7, 8, 9, 0, 2, 3, 4},
				[]int{8, 9, 0, 2, 3, 4, 5, 6, 7},
				[]int{3, 4, 5, 6, 7, 8, 9, 0, 2},
				[]int{6, 7, 8, 9, 0, 2, 3, 4, 5},
				[]int{0, 0, 2, 3, 4, 5, 6, 7, 8},
			},
			expectComplete: false,
			expectOutput: [][]int{
				[]int{0, 2, 3, 0, 5, 6, 7, 8, 9},
				[]int{0, 5, 6, 0, 8, 9, 1, 2, 3},
				[]int{0, 8, 9, 0, 2, 3, 4, 5, 6},
				[]int{2, 3, 4, 5, 6, 7, 8, 9, 1},
				[]int{5, 6, 7, 8, 9, 1, 2, 3, 4},
				[]int{8, 9, 1, 2, 3, 4, 5, 6, 7},
				[]int{3, 4, 5, 6, 7, 8, 9, 1, 2},
				[]int{6, 7, 8, 9, 1, 2, 3, 4, 5},
				[]int{9, 1, 2, 3, 4, 5, 6, 7, 8},
			},
		},
		{
			description: "two",
			input: [][]int{
				[]int{0, 0, 0, 0, 5, 0, 7, 8, 9},
				[]int{0, 0, 0, 0, 8, 0, 0, 2, 3},
				[]int{0, 0, 0, 0, 2, 0, 4, 5, 6},
				[]int{2, 3, 4, 5, 6, 0, 8, 9, 0},
				[]int{5, 6, 7, 8, 9, 0, 2, 0, 4},
				[]int{8, 9, 0, 2, 3, 0, 5, 6, 7},
				[]int{3, 4, 0, 6, 7, 0, 0, 0, 2},
				[]int{0, 7, 8, 9, 0, 0, 3, 4, 5},
				[]int{0, 0, 0, 0, 0, 0, 0, 0, 0},
			},
			expectComplete: false,
			expectOutput: [][]int{
				[]int{0, 0, 0, 0, 5, 0, 7, 8, 9},
				[]int{0, 5, 0, 0, 8, 0, 1, 2, 3},
				[]int{0, 8, 0, 0, 2, 0, 4, 5, 6},
				[]int{2, 3, 4, 5, 6, 7, 8, 9, 1},
				[]int{5, 6, 7, 8, 9, 1, 2, 3, 4},
				[]int{8, 9, 1, 2, 3, 4, 5, 6, 7},
				[]int{3, 4, 5, 6, 7, 8, 9, 1, 2},
				[]int{6, 7, 8, 9, 1, 2, 3, 4, 5},
				[]int{0, 0, 0, 3, 4, 5, 6, 7, 8},
			},
		},
		{
			description: "three",
			input: [][]int{
				[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
				[]int{4, 5, 6, 7, 8, 9, 0, 0, 0},
				[]int{7, 8, 9, 1, 2, 3, 4, 5, 6},
				[]int{0, 0, 0, 0, 0, 0, 1, 0, 0},
				[]int{0, 0, 0, 0, 0, 0, 0, 0, 0},
				[]int{0, 0, 0, 0, 0, 0, 0, 0, 0},
				[]int{0, 0, 0, 0, 0, 0, 0, 1, 0},
				[]int{0, 0, 0, 0, 0, 0, 0, 0, 0},
				[]int{0, 0, 0, 0, 0, 0, 0, 0, 0},
			},
			expectComplete: false,
			expectOutput: [][]int{
				[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
				[]int{4, 5, 6, 7, 8, 9, 0, 0, 1},
				[]int{7, 8, 9, 1, 2, 3, 4, 5, 6},
				[]int{0, 0, 0, 0, 0, 0, 1, 0, 0},
				[]int{0, 0, 0, 0, 0, 0, 0, 0, 0},
				[]int{0, 0, 0, 0, 0, 0, 0, 0, 0},
				[]int{0, 0, 0, 0, 0, 0, 0, 1, 0},
				[]int{0, 0, 0, 0, 0, 0, 0, 0, 0},
				[]int{0, 0, 0, 0, 0, 0, 0, 0, 0},
			},
		},
		{
			description: "four",
			input: [][]int{
				[]int{1, 2, 3, 0, 0, 0, 0, 0, 0},
				[]int{4, 5, 6, 0, 0, 0, 0, 0, 0},
				[]int{7, 8, 9, 0, 0, 0, 0, 0, 0},
				[]int{2, 3, 4, 0, 0, 0, 0, 0, 0},
				[]int{5, 6, 7, 0, 0, 0, 0, 0, 0},
				[]int{8, 9, 1, 0, 0, 0, 0, 0, 0},
				[]int{3, 0, 5, 1, 0, 0, 0, 0, 0},
				[]int{6, 0, 8, 0, 0, 0, 1, 0, 0},
				[]int{9, 0, 2, 0, 0, 0, 0, 0, 0},
			},
			expectComplete: false,
			expectOutput: [][]int{
				[]int{1, 2, 3, 0, 0, 0, 0, 0, 0},
				[]int{4, 5, 6, 0, 0, 0, 0, 0, 0},
				[]int{7, 8, 9, 0, 0, 0, 0, 0, 0},
				[]int{2, 3, 4, 0, 0, 0, 0, 0, 0},
				[]int{5, 6, 7, 0, 0, 0, 0, 0, 0},
				[]int{8, 9, 1, 0, 0, 0, 0, 0, 0},
				[]int{3, 0, 5, 1, 0, 0, 0, 0, 0},
				[]int{6, 0, 8, 0, 0, 0, 1, 0, 0},
				[]int{9, 1, 2, 0, 0, 0, 0, 0, 0},
			},
		},
		{
			description: "five",
			input: [][]int{
				[]int{1, 0, 0, 0, 0, 0, 0, 0, 0},
				[]int{0, 0, 0, 0, 0, 0, 0, 0, 0},
				[]int{0, 0, 0, 0, 0, 0, 0, 0, 0},
				[]int{0, 0, 0, 0, 0, 0, 0, 0, 0},
				[]int{0, 0, 0, 0, 0, 0, 1, 0, 0},
				[]int{0, 0, 0, 1, 0, 0, 0, 0, 0},
				[]int{0, 0, 1, 0, 0, 0, 0, 0, 0},
				[]int{0, 0, 0, 0, 0, 0, 0, 0, 0},
				[]int{0, 0, 0, 0, 0, 0, 0, 0, 0},
			},
			expectComplete: false,
			expectOutput: [][]int{
				[]int{1, 0, 0, 0, 0, 0, 0, 0, 0},
				[]int{0, 0, 0, 0, 0, 0, 0, 0, 0},
				[]int{0, 0, 0, 0, 0, 0, 0, 0, 0},
				[]int{0, 1, 0, 0, 0, 0, 0, 0, 0},
				[]int{0, 0, 0, 0, 0, 0, 1, 0, 0},
				[]int{0, 0, 0, 1, 0, 0, 0, 0, 0},
				[]int{0, 0, 1, 0, 0, 0, 0, 0, 0},
				[]int{0, 0, 0, 0, 0, 0, 0, 0, 0},
				[]int{0, 0, 0, 0, 0, 0, 0, 0, 0},
			},
		},
	}

	for _, td := range tt {
		t.Run(td.description, func(t *testing.T) {
			output, cg, err := SolveGrid(td.input)
			require.Nil(t, err)

			if td.printGrid {
				PrintGrid(output)
			}

			if !td.expectComplete {
				assert.Equal(t, td.expectOutput, output)
			}
			assert.Equal(t, CheckedGrid{Valid: true, Complete: td.expectComplete}, cg)
		})
	}
}

func TestSolveGridRealExamples(t *testing.T) {
	tt := []struct {
		description    string
		printGrid      bool
		input          [][]int
		expectComplete bool
		expectOutput   [][]int
	}{
		{
			description: "one",
			input: [][]int{
				[]int{2, 0, 7, 0, 0, 6, 0, 0, 0},
				[]int{0, 0, 0, 0, 3, 0, 2, 0, 6},
				[]int{0, 5, 6, 0, 0, 2, 0, 4, 1},
				[]int{1, 0, 0, 3, 0, 8, 7, 6, 0},
				[]int{6, 0, 9, 0, 0, 0, 1, 0, 8},
				[]int{0, 7, 4, 6, 0, 5, 0, 0, 3},
				[]int{5, 8, 0, 7, 0, 0, 4, 1, 0},
				[]int{9, 0, 1, 0, 5, 0, 0, 0, 0},
				[]int{0, 0, 0, 1, 0, 0, 3, 0, 5},
			},
			expectComplete: true,
		},
		{
			description: "two only passes with brute force",
			input: [][]int{
				[]int{3, 4, 2, 6, 9, 7, 5, 8, 1},
				[]int{6, 1, 8, 5, 2, 4, 7, 9, 3},
				[]int{5, 9, 7, 1, 8, 3, 4, 6, 2},
				[]int{2, 7, 3, 4, 1, 0, 9, 5, 6},
				[]int{1, 6, 4, 7, 5, 9, 0, 3, 0},
				[]int{9, 8, 5, 0, 3, 0, 1, 0, 7},
				[]int{8, 5, 9, 3, 7, 1, 6, 2, 4},
				[]int{7, 0, 6, 0, 4, 5, 0, 1, 9},
				[]int{4, 0, 1, 9, 6, 0, 0, 7, 5},
			},
			expectOutput: [][]int{
				[]int{3, 4, 2, 6, 9, 7, 5, 8, 1},
				[]int{6, 1, 8, 5, 2, 4, 7, 9, 3},
				[]int{5, 9, 7, 1, 8, 3, 0, 0, 0},
				[]int{2, 7, 3, 0, 1, 0, 9, 5, 0},
				[]int{1, 6, 4, 7, 5, 9, 0, 3, 0},
				[]int{9, 8, 5, 0, 3, 0, 1, 0, 7},
				[]int{8, 5, 9, 3, 7, 1, 6, 2, 4},
				[]int{7, 0, 6, 0, 4, 5, 0, 1, 9},
				[]int{4, 0, 1, 9, 6, 0, 0, 7, 5},
			},
			expectComplete: true,
		},
	}

	for _, td := range tt {
		t.Run(td.description, func(t *testing.T) {
			output, cg, err := SolveGrid(td.input)
			require.Nil(t, err)

			if td.printGrid {
				PrintGrid(output)
			}

			if !td.expectComplete {
				assert.Equal(t, td.expectOutput, output)
			}
			assert.Equal(t, CheckedGrid{Valid: true, Complete: td.expectComplete}, cg)
		})
	}
}

func TestGetRegion(t *testing.T) {
	tt := []struct {
		description    string
		ep             emptyPosition
		expectedRegion region
	}{
		{
			description: "top left one",
			ep: emptyPosition{
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
			ep: emptyPosition{
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
			ep: emptyPosition{
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
			ep: emptyPosition{
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
			ep: emptyPosition{
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
			region, err := getRegion(td.ep)
			require.Nil(t, err)
			assert.Equal(t, td.expectedRegion, region)
		})
	}
}

func TestRemainingColsToCheck(t *testing.T) {
	tt := []struct {
		description    string
		ep             emptyPosition
		reg            region
		expectedResult adjacentToCheck
	}{
		{
			description: "first",
			ep: emptyPosition{
				colNumber: 0,
				rowNumber: 0,
			},
			reg: region{
				minColNumber: 0,
				maxColNumber: 2,
				minRowNumber: 0,
				maxRowNumber: 2,
			},
			expectedResult: adjacentToCheck{
				adjacentCols: []int{1, 2},
				adjacentRows: []int{1, 2},
			},
		},
		{
			description: "second",
			ep: emptyPosition{
				colNumber: 1,
				rowNumber: 5,
			},
			reg: region{
				minColNumber: 0,
				maxColNumber: 2,
				minRowNumber: 3,
				maxRowNumber: 5,
			},
			expectedResult: adjacentToCheck{
				adjacentCols: []int{0, 2},
				adjacentRows: []int{3, 4},
			},
		},
		{
			description: "third",
			ep: emptyPosition{
				colNumber: 7,
				rowNumber: 8,
			},
			reg: region{
				minColNumber: 6,
				maxColNumber: 8,
				minRowNumber: 6,
				maxRowNumber: 8,
			},
			expectedResult: adjacentToCheck{
				adjacentCols: []int{6, 8},
				adjacentRows: []int{6, 7},
			},
		},
		{
			description: "fourth",
			ep: emptyPosition{
				colNumber: 8,
				rowNumber: 8,
			},
			reg: region{
				minColNumber: 6,
				maxColNumber: 8,
				minRowNumber: 6,
				maxRowNumber: 8,
			},
			expectedResult: adjacentToCheck{
				adjacentCols: []int{6, 7},
				adjacentRows: []int{6, 7},
			},
		},
	}

	for _, td := range tt {
		t.Run(td.description, func(t *testing.T) {
			res := adjacentRowsAndCols(td.reg, td.ep)
			sort.Ints(res.adjacentCols)
			sort.Ints(res.adjacentRows)
			assert.Equal(t, td.expectedResult, res)
		})
	}
}

func TestCheckGrid(t *testing.T) {
	tt := []struct {
		description  string
		input        [][]int
		expectReturn CheckedGrid
	}{
		{
			description: "valid and complete entry",
			input: [][]int{
				[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
				[]int{4, 5, 6, 7, 8, 9, 1, 2, 3},
				[]int{7, 8, 9, 1, 2, 3, 4, 5, 6},
				[]int{2, 3, 4, 5, 6, 7, 8, 9, 1},
				[]int{5, 6, 7, 8, 9, 1, 2, 3, 4},
				[]int{8, 9, 1, 2, 3, 4, 5, 6, 7},
				[]int{3, 4, 5, 6, 7, 8, 9, 1, 2},
				[]int{6, 7, 8, 9, 1, 2, 3, 4, 5},
				[]int{9, 1, 2, 3, 4, 5, 6, 7, 8},
			},
			expectReturn: CheckedGrid{Complete: true, Valid: true},
		},
		{
			description: "valid, not complete",
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
			expectReturn: CheckedGrid{Complete: false, Valid: true},
		},
		{
			description: "invalid, duplicate in row",
			input: [][]int{
				[]int{1, 1, 3, 4, 5, 6, 7, 8, 9},
				[]int{4, 5, 6, 7, 8, 9, 1, 2, 3},
				[]int{7, 8, 9, 1, 2, 3, 4, 5, 6},
				[]int{2, 3, 4, 5, 6, 7, 8, 9, 1},
				[]int{5, 6, 7, 8, 9, 1, 2, 3, 4},
				[]int{8, 9, 1, 2, 3, 4, 5, 6, 7},
				[]int{3, 4, 5, 6, 7, 8, 9, 1, 2},
				[]int{6, 7, 8, 9, 1, 2, 3, 4, 5},
				[]int{9, 1, 2, 3, 4, 5, 6, 7, 8},
			},
			expectReturn: CheckedGrid{Complete: false, Valid: false},
		},
		{
			description: "invalid, duplicate in column",
			input: [][]int{
				[]int{1, 0, 0, 0, 0, 0, 0, 0, 0},
				[]int{1, 0, 0, 0, 0, 0, 0, 0, 0},
				[]int{0, 0, 2, 0, 0, 0, 0, 0, 0},
				[]int{0, 0, 2, 0, 0, 0, 0, 0, 0},
				[]int{0, 0, 0, 0, 0, 0, 0, 0, 0},
				[]int{0, 0, 0, 0, 0, 0, 0, 0, 0},
				[]int{0, 0, 0, 0, 0, 0, 0, 0, 0},
				[]int{0, 0, 0, 0, 0, 0, 0, 0, 0},
				[]int{0, 0, 0, 0, 0, 0, 0, 0, 0},
			},
			expectReturn: CheckedGrid{Complete: false, Valid: false},
		},
		{
			description: "invalid, duplicate in top grid region",
			input: [][]int{
				[]int{1, 0, 0, 0, 0, 0, 0, 0, 0},
				[]int{0, 1, 0, 0, 0, 0, 0, 0, 0},
				[]int{0, 0, 1, 0, 0, 0, 0, 0, 0},
				[]int{0, 0, 0, 0, 0, 0, 0, 0, 0},
				[]int{0, 0, 0, 0, 0, 0, 0, 0, 0},
				[]int{0, 0, 0, 0, 0, 0, 0, 0, 0},
				[]int{0, 0, 0, 0, 0, 0, 0, 0, 0},
				[]int{0, 0, 0, 0, 0, 0, 0, 0, 0},
				[]int{0, 0, 0, 0, 0, 0, 0, 0, 0},
			},
			expectReturn: CheckedGrid{Complete: false, Valid: false},
		},
		{
			description: "invalid, duplicate in bottom grid region",
			input: [][]int{
				[]int{0, 0, 0, 0, 0, 0, 0, 0, 0},
				[]int{0, 0, 0, 0, 0, 0, 0, 0, 0},
				[]int{0, 0, 0, 0, 0, 0, 0, 0, 0},
				[]int{0, 0, 0, 0, 0, 0, 0, 0, 0},
				[]int{0, 0, 0, 0, 0, 0, 0, 0, 0},
				[]int{0, 0, 0, 0, 0, 0, 0, 0, 0},
				[]int{0, 0, 0, 0, 0, 0, 1, 0, 0},
				[]int{0, 0, 0, 0, 0, 0, 0, 1, 0},
				[]int{0, 0, 0, 0, 0, 0, 0, 0, 0},
			},
			expectReturn: CheckedGrid{Complete: false, Valid: false},
		},
		{
			description: "invalid, duplicate in middle grid region",
			input: [][]int{
				[]int{0, 0, 0, 0, 0, 0, 0, 0, 0},
				[]int{0, 0, 0, 0, 0, 0, 0, 0, 0},
				[]int{0, 0, 0, 0, 0, 0, 0, 0, 0},
				[]int{0, 0, 0, 0, 0, 0, 0, 0, 0},
				[]int{0, 0, 0, 0, 1, 0, 0, 0, 0},
				[]int{0, 0, 0, 0, 0, 1, 0, 0, 0},
				[]int{0, 0, 0, 0, 0, 0, 0, 0, 0},
				[]int{0, 0, 0, 0, 0, 0, 0, 0, 0},
				[]int{0, 0, 0, 0, 0, 0, 0, 0, 0},
			},
			expectReturn: CheckedGrid{Complete: false, Valid: false},
		},
	}

	for _, td := range tt {
		t.Run(td.description, func(t *testing.T) {
			cg := CheckGrid(td.input)
			assert.Equal(t, td.expectReturn.Complete, cg.Complete)
			assert.Equal(t, td.expectReturn.Valid, cg.Valid)
		})
	}
}
