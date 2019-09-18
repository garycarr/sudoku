package soduku

import (
	"fmt"
)

type emptyPosition struct {
	rowNumber int
	colNumber int
}

type region struct {
	minColNumber int
	maxColNumber int
	minRowNumber int
	maxRowNumber int
}

type adjacentToCheck struct {
	adjacentRows []int
	adjacentCols []int
}

// SolveGrid attempts to solve a given suduko board
func SolveGrid(wholeGrid [][]int) ([][]int, error) {
	previousNumEPS := 0

	for {
		eps := getEmptyPositions(wholeGrid)
		if len(eps) == 0 || len(eps) == previousNumEPS {
			break
		}
		previousNumEPS = len(eps)

		for _, ep := range eps {
			foundNum, err := traverseImmediateLines(wholeGrid, ep)
			if err != nil {
				return nil, err
			}
			if foundNum > 0 {
				wholeGrid[ep.rowNumber][ep.colNumber] = foundNum
			}

			if err := traverseAdjacent(wholeGrid, ep); err != nil {
				return nil, err
			}
		}
	}

	return wholeGrid, nil
}

// getEmptyPositions returns the positions of empty boxes
func getEmptyPositions(wholeGrid [][]int) []emptyPosition {
	eps := []emptyPosition{}

	// identify missing positions
	for rowNumber, row := range wholeGrid {
		for colNumber, num := range row {
			if num > 0 {
				continue
			}
			eps = append(eps, emptyPosition{
				rowNumber: rowNumber,
				colNumber: colNumber,
			})
		}
	}
	return eps
}

// traverseImmediateLines checks the column, row, and grid the emptyPosition is in
// to look for entries it can make
func traverseImmediateLines(wholeGrid [][]int, ep emptyPosition) (int, error) {
	// check the box it is in
	reg, err := getRegion(ep)
	if err != nil {
		return 0, err
	}

	pn, err := possibleNumbers(wholeGrid, ep, reg)
	if err != nil {
		return 0, err
	}

	if len(pn) == 1 {
		return pn[0], nil
	}
	// There is more than one number available, so we cannot determine which one to use
	return 0, nil
}

// possibleNumbers returns the numbers that can possibly placed into a given position
func possibleNumbers(wholeGrid [][]int, ep emptyPosition, reg region) ([]int, error) {
	possibleNumbers := map[int]bool{}
	for i := 1; i <= 9; i++ {
		possibleNumbers[i] = false
	}

	// check the row it is on
	for col := 0; col <= 8; col++ {
		if possibleNumbers[wholeGrid[ep.rowNumber][col]] {
			continue
		}
		possibleNumbers[wholeGrid[ep.rowNumber][col]] = true
	}

	// check the column it is in
	for row := 0; row <= 8; row++ {
		if possibleNumbers[wholeGrid[row][ep.colNumber]] {
			continue
		}
		possibleNumbers[wholeGrid[row][ep.colNumber]] = true
	}

	// Check the grid it is in
	for row := reg.minRowNumber; row <= reg.maxRowNumber; row++ {
		for col := reg.minColNumber; col <= reg.maxColNumber; col++ {
			if possibleNumbers[wholeGrid[row][col]] {
				continue
			}
			possibleNumbers[wholeGrid[row][col]] = true
		}
	}

	nums := []int{}
	for num, found := range possibleNumbers {
		if !found {
			nums = append(nums, num)
		}
	}

	return nums, nil
}

// traverseAdjacent looks at the rows and columns next to the position to find entries
// For example if there if the grid looks like this
//
// 1, 0, 0, 0, 0, 0, 0, 0, 0
// 0, 0, 0, 0, 0, 0, 0, 0, 0
// 0, 0, 0, 0, 1, 0, 0, 0, 0
// 0, 0, 0, 0, 0, 0, 1, 0, 0
// 0, 0, 0, 0, 0, 0, 0, 0, 0
// 0, 0, 0, 0, 0, 0, 0, 0, 0
// 0, 0, 0, 0, 0, 0, 0, 1, 0
// 0, 0, 0, 0, 0, 0, 0, 0, 0
// 0, 0, 0, 0, 0, 0, 0, 0, 0
//
// Then at position {1,8} there has to be a 1, as it cannot go anywhere else in the top right grid
func traverseAdjacent(wholeGrid [][]int, ep emptyPosition) error {
	reg, err := getRegion(ep)
	if err != nil {
		return err
	}

	pn, err := possibleNumbers(wholeGrid, ep, reg)
	if err != nil {
		return err
	}

	r := adjacentRowsAndCols(reg, ep)

	for _, num := range pn {
		foundNumColAndRow := 0
		foundNum := 0
		// check the adjacent columns
		for i := 0; i <= 8; i++ {

			for _, c := range r.adjacentCols {
				if wholeGrid[i][c] == num {
					foundNum++
					if foundNum == 2 {
						// We found the number in both adjacent columns, so it has to be in this column
						// Now check to see if the boxes next to the position are populated, if they
						// are we know this is the correct position for this number
						foundNumColAndRow++

						alreadyPopulated := 0
						for _, r := range r.adjacentRows {
							if wholeGrid[r][ep.colNumber] != 0 {
								alreadyPopulated++
							}
						}
						if alreadyPopulated == 2 {
							wholeGrid[ep.rowNumber][ep.colNumber] = num
						}
						break
					}
				}
			}
		}

		foundNum = 0
		for _, row := range r.adjacentRows {

			for col := 0; col <= 8; col++ {
				if wholeGrid[row][col] == num {
					foundNum++
				}
				if foundNum == 2 {
					// We found the number in both adjacent rows, so it has to be in this row
					// Now check to see if the boxes next to the position are populated, if they
					// are we know this is the correct position for this number
					foundNumColAndRow++
					alreadyPopulated := 0
					for _, c := range r.adjacentCols {
						if wholeGrid[ep.rowNumber][c] != 0 {
							alreadyPopulated++
						}
					}
					if alreadyPopulated == 2 {
						wholeGrid[ep.rowNumber][ep.colNumber] = num
					}
					break
				}
			}
		}

		// Because the entry was identified in both row and column, we know this is the correct location
		// even though there is empty boxes next to the position
		if foundNumColAndRow == 2 {
			wholeGrid[ep.rowNumber][ep.colNumber] = num
		}
	}
	return nil
}

// adjacentRowsAndCols the rows and columns next to the position, but within the same grid
func adjacentRowsAndCols(reg region, ep emptyPosition) adjacentToCheck {
	r := adjacentToCheck{}

	switch reg.maxRowNumber - ep.rowNumber {
	case 0:
		r.adjacentRows = append(r.adjacentRows, ep.rowNumber-1)
		r.adjacentRows = append(r.adjacentRows, ep.rowNumber-2)
	case 1:
		r.adjacentRows = append(r.adjacentRows, ep.rowNumber+1)
		r.adjacentRows = append(r.adjacentRows, ep.rowNumber-1)
	case 2:
		r.adjacentRows = append(r.adjacentRows, ep.rowNumber+1)
		r.adjacentRows = append(r.adjacentRows, ep.rowNumber+2)
	}

	switch reg.maxColNumber - ep.colNumber {
	case 0:
		r.adjacentCols = append(r.adjacentCols, ep.colNumber-1)
		r.adjacentCols = append(r.adjacentCols, ep.colNumber-2)
	case 1:
		r.adjacentCols = append(r.adjacentCols, ep.colNumber+1)
		r.adjacentCols = append(r.adjacentCols, ep.colNumber-1)
	case 2:
		r.adjacentCols = append(r.adjacentCols, ep.colNumber+1)
		r.adjacentCols = append(r.adjacentCols, ep.colNumber+2)
	}

	return r
}

// getRegion returns he grid position that the emptyPosition is in
func getRegion(ep emptyPosition) (region, error) {
	reg := region{}
	switch {
	case ep.rowNumber >= 0 && ep.rowNumber <= 2:
		reg.minRowNumber = 0
		reg.maxRowNumber = 2
	case ep.rowNumber >= 3 && ep.rowNumber <= 5:
		reg.minRowNumber = 3
		reg.maxRowNumber = 5
	case ep.rowNumber >= 6 && ep.rowNumber <= 8:
		reg.minRowNumber = 6
		reg.maxRowNumber = 8
	default:
		return reg, fmt.Errorf("rowNumber %d is invalid", ep.rowNumber)
	}

	switch {
	case ep.colNumber >= 0 && ep.colNumber <= 2:
		reg.minColNumber = 0
		reg.maxColNumber = 2
	case ep.colNumber >= 3 && ep.colNumber <= 5:
		reg.minColNumber = 3
		reg.maxColNumber = 5
	case ep.colNumber >= 6 && ep.colNumber <= 8:
		reg.minColNumber = 6
		reg.maxColNumber = 8
	default:
		return reg, fmt.Errorf("colNumber %d is invalid", ep.colNumber)
	}
	return reg, nil
}
