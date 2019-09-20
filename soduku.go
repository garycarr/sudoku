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

type adjacentToCheck struct {
	adjacentRows []int
	adjacentCols []int
}

// SolveGrid attempts to solve a given suduko board. It returns the grid as complete as it
// could achieve, and a struct indicating the status of the grid
func SolveGrid(grid [][]int) ([][]int, CheckedGrid, error) {
	// previousNumSquares holds the previous loops count of how many empty squares exist
	previousNumSquares := 0

	cg := CheckedGrid{}

	for {
		ss, err := NewSquares(grid)
		if err != nil {
			return nil, cg, err
		}

		if len(ss) == 0 || len(ss) == previousNumSquares {
			break
		}
		previousNumSquares = len(ss)

		for _, s := range ss {
			if len(s.possibleNums) == 1 {
				grid[s.pos.rowNumber][s.pos.colNumber] = s.possibleNums[0]
			}
		}
		ss, err = NewSquares(grid)
		if err != nil {
			return nil, cg, err
		}
		for _, s := range ss {
			if err := traverseAdjacent(grid, s); err != nil {
				return nil, cg, err
			}
		}
	}
	cg = CheckGrid(grid)
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
func traverseAdjacent(grid [][]int, s *square) error {
	r := adjacentRowsAndCols(s.reg, s.pos)

	for _, num := range s.possibleNums {
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
							if grid[r][s.pos.colNumber] != 0 {
								alreadyPopulated++
							}
						}
						if alreadyPopulated == 2 {
							grid[s.pos.rowNumber][s.pos.colNumber] = num
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
						if grid[s.pos.rowNumber][c] != 0 {
							alreadyPopulated++
						}
					}
					if alreadyPopulated == 2 {
						grid[s.pos.rowNumber][s.pos.colNumber] = num
					}
					break
				}
			}
		}

		// Because the entry was identified in both row and column, we know this is the correct location
		// even though there is empty boxes next to the position
		if foundNumColAndRow == 2 {
			grid[s.pos.rowNumber][s.pos.colNumber] = num
		}
	}
	return nil
}

// bruteForceGuess adds in numbers to empty positions and sees if it can solve the rest of the grid
// This is a weak brute force, it should try combinations of numbers
func bruteForceGuess(grid [][]int) ([][]int, error) {
	ss, err := NewSquares(grid)
	if err != nil {
		return nil, err
	}

	// poss := getEmptyPositions(grid)
	copyGrid := func(grid [][]int) [][]int {
		tempGrid := make([][]int, len(grid))
		for i := range grid {
			tempGrid[i] = make([]int, len(grid[i]))
			copy(tempGrid[i], grid[i])
		}
		return tempGrid
	}

	tempGrid := copyGrid(grid)

	for _, s := range ss {
		if len(s.possibleNums) == 1 {
			tempGrid[s.pos.rowNumber][s.pos.colNumber] = s.possibleNums[0]
		}

		if err := traverseAdjacent(tempGrid, s); err != nil {
			return nil, err
		}
		cg := CheckGrid(tempGrid)
		if !cg.Valid {
			tempGrid = copyGrid(grid)
			continue
		}
		if cg.Complete {
			return tempGrid, nil
		}
	}
	return grid, nil
}

// adjacentRowsAndCols the rows and columns next to the position, but within the same grid
func adjacentRowsAndCols(reg region, pos position) adjacentToCheck {
	r := adjacentToCheck{}

	switch reg.maxRowNumber - pos.rowNumber {
	case 0:
		r.adjacentRows = append(r.adjacentRows, pos.rowNumber-1)
		r.adjacentRows = append(r.adjacentRows, pos.rowNumber-2)
	case 1:
		r.adjacentRows = append(r.adjacentRows, pos.rowNumber+1)
		r.adjacentRows = append(r.adjacentRows, pos.rowNumber-1)
	case 2:
		r.adjacentRows = append(r.adjacentRows, pos.rowNumber+1)
		r.adjacentRows = append(r.adjacentRows, pos.rowNumber+2)
	}

	switch reg.maxColNumber - pos.colNumber {
	case 0:
		r.adjacentCols = append(r.adjacentCols, pos.colNumber-1)
		r.adjacentCols = append(r.adjacentCols, pos.colNumber-2)
	case 1:
		r.adjacentCols = append(r.adjacentCols, pos.colNumber+1)
		r.adjacentCols = append(r.adjacentCols, pos.colNumber-1)
	case 2:
		r.adjacentCols = append(r.adjacentCols, pos.colNumber+1)
		r.adjacentCols = append(r.adjacentCols, pos.colNumber+2)
	}

	return r
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
