package soduku

import (
	"errors"
	"fmt"
)

// CheckedGrid stores whether a grid is valid, and complete
type CheckedGrid struct {
	Complete bool
	Message  string
	Valid    bool
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

// // SolveGrid attempts to solve a given suduko board. It returns the grid as complete as it
// // could achieve, and a struct indicating the status of the grid
// func SolveGrid(grid [][]int) ([][]int, CheckedGrid, error) {
// 	// previousNumEPS holds the previous loops count of how many emptyPositions exist.
// 	// if this number does not decrease then there is no more need to iterate
// 	previousNumEPS := 0
//
// 	cg := CheckedGrid{}
//
// 	for {
// 		eps := getEmptyPositions(grid)
// 		if len(eps) == 0 || len(eps) == previousNumEPS {
// 			break
// 		}
// 		previousNumEPS = len(eps)
//
// 		for _, ep := range eps {
// 			foundNum, err := traverseImmediateLines(grid, ep)
// 			if err != nil {
// 				return nil, cg, err
// 			}
// 			if foundNum > 0 {
// 				grid[ep.rowNumber][ep.colNumber] = foundNum
// 			}
//
// 			if err := traverseAdjacent(grid, ep); err != nil {
// 				return nil, cg, err
// 			}
// 		}
// 	}
// 	cg = CheckGrid(grid)
//
// 	if !cg.Valid {
// 		return grid, cg, errors.New("the grid is invalid")
// 	}
// 	if cg.Complete {
// 		return grid, cg, nil
// 	}
//
// 	var err error
// 	grid, err = bruteForceGuess(grid)
// 	if err != nil {
// 		return grid, cg, err
// 	}
// 	cg = CheckGrid(grid)
// 	if !cg.Valid {
// 		return grid, cg, errors.New("the grid is invalid after brute forcing")
// 	}
// 	return grid, cg, err
// }

// SolveGrid attempts to solve a given suduko board. It returns the grid as complete as it
// could achieve, and a struct indicating the status of the grid
func SolveGrid(grid [][]int) ([][]int, CheckedGrid, error) {
	// previousNumEPS holds the previous loops count of how many emptyPositions exist.
	// if this number does not decrease then there is no more need to iterate
	previousNumEPS := 0

	cg := CheckedGrid{}

	for {
		epsPNs, err := getEmptyPositionsAndPossibleNumbers(grid)
		if err != nil {
			return nil, cg, err
		}
		// count epsPN.possibleNums?
		if len(epsPNs) == 0 || len(epsPNs) == previousNumEPS {
			break
		}
		previousNumEPS = len(epsPNs)

		for _, epPN := range epsPNs {
			foundNum, err := traverseImmediateLines(grid, epPN)
			if err != nil {
				return nil, cg, err
			}
			if foundNum > 0 {
				grid[epPN.ep.rowNumber][epPN.ep.colNumber] = foundNum
			}

			if err := traverseAdjacent(grid, epPN); err != nil {
				return nil, cg, err
			}
		}
	}
	cg = CheckGrid(grid)

	PrintGrid(grid)
	if !cg.Valid {
		return grid, cg, errors.New("the grid is invalid")
	}
	if cg.Complete {
		return grid, cg, nil
	}

	var err error
	grid, err = bruteForceGuess(grid)
	if err != nil {
		return grid, cg, err
	}
	cg = CheckGrid(grid)
	if !cg.Valid {
		return grid, cg, errors.New("the grid is invalid after brute forcing")
	}
	return grid, cg, err
}

// CheckGrid returns where a given grid is complete, and if it is valid
func CheckGrid(grid [][]int) CheckedGrid {
	cg := CheckedGrid{Valid: true, Complete: true, Message: ""}

	// Check all rows
	totalRows := 0
	for rowNum, row := range grid {
		totalRows++
		foundNumbersRows := make(map[int]int, 9)

		for i := 0; i <= 8; i++ {
			if row[i] > 0 {
				foundNumbersRows[row[i]]++
			} else {
				cg.Complete = false
			}
		}
		for num, count := range foundNumbersRows {
			if count != 1 {
				cg.Complete = false
				if count > 1 {
					cg.Message = fmt.Sprintf("%s A duplicate of %d was found in row %d\n", cg.Message, num, rowNum)
					cg.Valid = false
				}
			}
		}
	}
	if totalRows != 9 {
		cg.Message = fmt.Sprintf("%s Expected 9 rows, found %d", cg.Message, totalRows)
		cg.Valid = false
		cg.Complete = false
	}

	// Check all columns
	for colNum := 0; colNum <= 8; colNum++ {
		foundNumbersCol := make(map[int]int, 9)
		for rowNum := 0; rowNum <= 8; rowNum++ {
			num := grid[rowNum][colNum]
			if num > 0 {
				foundNumbersCol[num]++
			} else {
				cg.Complete = false
			}
		}

		for num, count := range foundNumbersCol {
			if count != 1 {
				cg.Complete = false
				if count > 1 {
					cg.Message = fmt.Sprintf("%s A duplicate of %d was found in column %d\n", cg.Message, num, colNum)
					cg.Valid = false
				}
			}
		}
	}

	// Check all the regions
	for _, reg := range allRegions {
		foundNumbersGrid := make(map[int]int, 9)
		for row := reg.minRowNumber; row <= reg.maxRowNumber; row++ {
			for col := reg.minColNumber; col <= reg.maxColNumber; col++ {
				num := grid[row][col]
				if num > 0 {
					foundNumbersGrid[num]++
				}
			}
		}
		for num, count := range foundNumbersGrid {
			if count != 1 {
				cg.Complete = false
				if count > 1 {
					gridPosition := fmt.Sprintf("rowNumber {%d, %d}, colNumber {%d, %d}",
						reg.minRowNumber, reg.maxRowNumber, reg.minColNumber, reg.maxColNumber)
					cg.Message = fmt.Sprintf("%s A duplicate of %d was found in grid %q\n", cg.Message, num, gridPosition)
					cg.Valid = false
				}
			}
		}
	}
	return cg
}

// traverseImmediateLines checks the column, row, and grid the emptyPosition is in
// to look for entries it can make
func traverseImmediateLines(grid [][]int, ep emptyPositionAndPossibleNumbers) (int, error) {
	// // check the box it is in
	// reg, err := getRegion(ep)
	// if err != nil {
	// 	return 0, err
	// }
	//
	// pn, err := possibleNumbers(grid, ep, reg)
	// if err != nil {
	// 	return 0, err
	// }

	if len(ep.possibleNums) == 1 {
		return ep.possibleNums[0], nil
	}
	// There is more than one number available, so we cannot determine which one to use
	return 0, nil
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
func traverseAdjacent(grid [][]int, epPN emptyPositionAndPossibleNumbers) error {
	reg, err := getRegion(epPN.ep)
	if err != nil {
		return err
	}

	r := adjacentRowsAndCols(reg, epPN.ep)

	for _, num := range epPN.possibleNums {
		foundNumColAndRow := 0
		foundNum := 0
		// check the adjacent columns
		for i := 0; i <= 8; i++ {

			for _, c := range r.adjacentCols {
				if grid[i][c] == num {
					foundNum++
					if foundNum == 2 {
						// We found the number in both adjacent columns, so it has to be in this column
						// Now check to see if the boxes next to the position are populated, if they
						// are we know this is the correct position for this number
						foundNumColAndRow++

						alreadyPopulated := 0
						for _, r := range r.adjacentRows {
							if grid[r][epPN.ep.colNumber] != 0 {
								alreadyPopulated++
							}
						}
						if alreadyPopulated == 2 {
							grid[epPN.ep.rowNumber][epPN.ep.colNumber] = num
						}
						break
					}
				}
			}
		}

		foundNum = 0
		for _, row := range r.adjacentRows {

			for col := 0; col <= 8; col++ {
				if grid[row][col] == num {
					foundNum++
				}
				if foundNum == 2 {
					// We found the number in both adjacent rows, so it has to be in this row
					// Now check to see if the boxes next to the position are populated, if they
					// are we know this is the correct position for this number
					foundNumColAndRow++
					alreadyPopulated := 0
					for _, c := range r.adjacentCols {
						if grid[epPN.ep.rowNumber][c] != 0 {
							alreadyPopulated++
						}
					}
					if alreadyPopulated == 2 {
						grid[epPN.ep.rowNumber][epPN.ep.colNumber] = num
					}
					break
				}
			}
		}

		// Because the entry was identified in both row and column, we know this is the correct location
		// even though there is empty boxes next to the position
		if foundNumColAndRow == 2 {
			grid[epPN.ep.rowNumber][epPN.ep.colNumber] = num
		}
	}
	return nil
}

// bruteForceGuess adds in numbers to empty positions and sees if it can solve the rest of the grid
// This is a weak brute force, it should try combinations of numbers
func bruteForceGuess(grid [][]int) ([][]int, error) {
	epPNs, err := getEmptyPositionsAndPossibleNumbers(grid)
	if err != nil {
		return nil, err
	}

	// eps := getEmptyPositions(grid)
	copyGrid := func(grid [][]int) [][]int {
		tempGrid := make([][]int, len(grid))
		for i := range grid {
			tempGrid[i] = make([]int, len(grid[i]))
			copy(tempGrid[i], grid[i])
		}
		return tempGrid
	}

	tempGrid := copyGrid(grid)

	for _, epPN := range epPNs {
		for _, pn := range epPN.possibleNums {
			foundNum, err := traverseImmediateLines(tempGrid, epPN)
			if err != nil {
				return nil, err
			}
			if foundNum > 0 {
				tempGrid[epPN.ep.rowNumber][epPN.ep.colNumber] = pn
			}

			if err := traverseAdjacent(tempGrid, epPN); err != nil {
				return nil, err
			}
			cg := CheckGrid(tempGrid)
			if !cg.Valid {
				tempGrid = copyGrid(grid)
			}
			if cg.Complete {
				return tempGrid, nil
			}
		}
	}
	return grid, nil
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

// PrintGrid prints out the grid to the terminal
func PrintGrid(grid [][]int) {
	println("")
	for _, row := range grid {
		line := ""
		for _, num := range row {
			line = fmt.Sprintf("%s | %d", line, num)
		}
		println(fmt.Sprintf("%s |", line))
	}
	println("")
}
