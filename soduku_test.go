package soduku

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSolveGridSimple(t *testing.T) {
	tt := []struct {
		description string
		drawGrid    bool
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
			output, err := SolveGrid(td.input)
			require.Nil(t, err)
			checkWholeGrid(t, output, true)

			if td.drawGrid {
				DrawGrid(output)
			}
		})
	}
}

func TestSolveGridMedium(t *testing.T) {
	tt := []struct {
		description   string
		drawGrid      bool
		input         [][]int
		expectSuccess bool
		expectOutput  [][]int
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
			expectSuccess: false,
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
			expectSuccess: false,
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
			expectSuccess: false,
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
			expectSuccess: false,
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
			expectSuccess: false,
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
			output, err := SolveGrid(td.input)
			require.Nil(t, err)

			if td.drawGrid {
				DrawGrid(output)
			}

			if !td.expectSuccess {
				assert.Equal(t, td.expectOutput, output)
			}
			checkWholeGrid(t, output, td.expectSuccess)
		})
	}
}

func TestSolveGridRealExamples(t *testing.T) {
	tt := []struct {
		description   string
		drawGrid      bool
		input         [][]int
		expectSuccess bool
		expectOutput  [][]int
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
			expectSuccess: true,
		},
		{
			description: "two",
			input: [][]int{
				[]int{0, 0, 0, 0, 9, 0, 0, 0, 1},
				[]int{6, 1, 0, 5, 2, 0, 7, 0, 0},
				[]int{5, 0, 7, 0, 0, 3, 0, 0, 0},
				[]int{2, 0, 3, 0, 1, 0, 0, 5, 0},
				[]int{0, 0, 0, 7, 0, 9, 0, 0, 0},
				[]int{0, 8, 0, 0, 3, 0, 1, 0, 7},
				[]int{0, 0, 0, 3, 0, 0, 6, 0, 4},
				[]int{0, 0, 6, 0, 4, 5, 0, 1, 9},
				[]int{4, 0, 0, 0, 6, 0, 0, 0, 0},
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
			expectSuccess: false,
		},
	}

	for _, td := range tt {
		t.Run(td.description, func(t *testing.T) {
			output, err := SolveGrid(td.input)
			require.Nil(t, err)

			if td.drawGrid {
				DrawGrid(output)
			}

			if !td.expectSuccess {
				assert.Equal(t, td.expectOutput, output)
			}
			checkWholeGrid(t, output, td.expectSuccess)
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

func checkWholeGrid(t *testing.T, wholeGrid [][]int, expectAll bool) {
	checkRows(t, wholeGrid, expectAll)
	checkColumns(t, wholeGrid, expectAll)
	checkGridBoxes(t, wholeGrid, expectAll)
}

// checkRows makes sure all rows are present, 1-9
func checkRows(t *testing.T, wholeGrid [][]int, expectAll bool) {
	for _, row := range wholeGrid {
		foundNumbers := make(map[int]int, 9)

		for i := 0; i <= 8; i++ {
			foundNumbers[row[i]]++
		}

		if expectAll {
			require.Len(t, foundNumbers, 9)
		}

		for i := 1; i <= 9; i++ {
			if expectAll {
				// each row should have 1 - 9
				assert.Equal(t, 1, foundNumbers[i])
			} else {
				// just check there were no duplicates
				if foundNumbers[i] > 1 {
					t.Errorf("Duplicate row found for %d", foundNumbers[i])
				}
			}
		}
	}
}

// checkRows makes sure all rows are present, 1-9
func checkColumns(t *testing.T, wholeGrid [][]int, expectAll bool) {
	for col := 0; col <= 8; col++ {
		foundNumbers := make(map[int]int, 9)
		for row := 0; row <= 8; row++ {
			num := wholeGrid[row][col]
			foundNumbers[num]++
		}
		if expectAll {
			require.Len(t, foundNumbers, 9)
		}
		for i := 1; i <= 9; i++ {
			if expectAll {
				// each column should have 1 - 9
				assert.Equal(t, 1, foundNumbers[i])
			} else {
				// just check there were no duplicates
				if foundNumbers[i] > 1 {
					t.Errorf("Duplicate col found for %d", foundNumbers[i])
				}
			}
		}
	}
}

// checkGridBoxes makes sure each box has 1-9
func checkGridBoxes(t *testing.T, wholeGrid [][]int, expectAll bool) {
	reg := region{
		minRowNumber: 0,
		maxRowNumber: 2,
		minColNumber: 0,
		maxColNumber: 2,
	}

	foundNumbers := make(map[int]int, 9)
	for row := reg.minRowNumber; row <= reg.maxRowNumber; row++ {
		for col := reg.minColNumber; col <= reg.maxColNumber; col++ {
			foundNumbers[wholeGrid[row][col]]++
		}
	}
	gridPosition := fmt.Sprintf("rowNumber {%d, %d}, colNumber {%d, %d}",
		reg.minRowNumber, reg.maxRowNumber, reg.minColNumber, reg.maxColNumber)

	if expectAll {
		require.Len(t, foundNumbers, 9, fmt.Sprintf("len for grid %s", gridPosition))
	}
	for i := 1; i <= 9; i++ {
		if expectAll {
			// each box should have 1 - 9
			assert.Equal(t, 1, foundNumbers[i], fmt.Sprintf("for num '%d' %s", i, gridPosition))
		} else {
			// just check there were no duplicates
			if foundNumbers[i] > 1 {
				t.Errorf("Duplicate in grid found for %d in grid %s", i, gridPosition)
			}
		}
	}
}
