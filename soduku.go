package soduku

import (
	"errors"
	"fmt"

	"github.com/davecgh/go-spew/spew"
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
	cg := CheckedGrid{}
	grid, err := solver(grid)
	if err != nil {
		return nil, cg, err
	}
	cg = CheckGrid(grid)
	if !cg.Valid {
		return grid, cg, errors.New("the grid is invalid")
	}
	if cg.Complete {
		return grid, cg, nil
	}

	grid, err = bruteForceGuess(grid)
	if err != nil {
		return grid, cg, err
	}

	return grid, CheckGrid(grid), nil
}

func solver(grid [][]int) ([][]int, error) {
	// previousNumSquares holds the previous loops count of how many empty squares exist
	previousNumSquares := 0

	for {
		ss, err := NewSquares(grid)
		if err != nil {
			return nil, err
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
			return nil, err
		}
		for _, s := range ss {
			if err := traverseAdjacent(grid, s); err != nil {
				return nil, err
			}
		}
	}
	return grid, nil
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
	originalGrid := copyOriginalGrid(grid)

	bruteForcedGrid, err := bruteSolver(grid)
	if err != nil {
		return nil, err
	}

	cg := CheckGrid(bruteForcedGrid)
	if cg.Complete {
		return bruteForcedGrid, nil
	}

	return originalGrid, nil
}

func bruteSolver(originalGrid [][]int) ([][]int, error) {
	bruteForcedGrid := copyOriginalGrid(originalGrid)

	// grid, completed, err
	blah := func(grid [][]int, pn int, s *square) ([][]int, bool, error) {
		grid[s.pos.rowNumber][s.pos.colNumber] = pn
		grid, err := solver(grid)
		cg := CheckGrid(grid)

		if cg.Complete {
			return bruteForcedGrid, true, nil
		}

		if err != nil {
			return bruteForcedGrid, false, err
		}

		return bruteForcedGrid, false, nil
	}

	ss, err := NewSquares(originalGrid)
	if err != nil {
		return nil, err
	}
	for _, s := range ss {
		if len(s.possibleNums) == 1 {
			bruteForcedGrid[s.pos.rowNumber][s.pos.colNumber] = s.possibleNums[0]
			cg := CheckGrid(bruteForcedGrid)
			// Check cg first, err might be for an invalid grid which is okay here
			if cg.Complete {
				return bruteForcedGrid, nil
			}
		}

		for _, pn := range s.possibleNums {
			bruteForcedGrid, completed, err := blah(bruteForcedGrid, pn, s)
			if err != nil {
				return nil, err
			}
			if completed {
				return bruteForcedGrid, nil
			}

			ss, err := NewSquares(bruteForcedGrid)
			if err != nil {
				return nil, err
			}

			for _, secondS := range ss {
				if len(s.possibleNums) == 1 {
					bruteForcedGrid[s.pos.rowNumber][secondS.pos.colNumber] = secondS.possibleNums[0]
					cg := CheckGrid(bruteForcedGrid)
					// Check cg first, err might be for an invalid grid which is okay here
					if cg.Complete && cg.Valid {
						return bruteForcedGrid, nil
					}
				}

				for _, secondPn := range secondS.possibleNums {
					spew.Dump(fmt.Sprintf("Start second looking at pn %d first square row:%d col:%d  second square row:%d col:%d",
						secondPn, secondS.pos.rowNumber, secondS.pos.colNumber, secondS.pos.rowNumber, secondS.pos.colNumber))
					spew.Dump(secondS)
					// if s.pos.colNumber == 0 && s.pos.rowNumber == 0 {
					// 	spew.Dump(fmt.Sprintf("Start second looking at pn %d", pn))
					// 	spew.Dump(secondS)
					//
					spew.Dump("before")
					PrintGrid(bruteForcedGrid)
					// 	spew.Dump("end second")
					// }
					bruteForcedGrid, completed, err := blah(bruteForcedGrid, secondPn, secondS)
					// if s.pos.colNumber == 0 && s.pos.rowNumber == 0 {
					// 	spew.Dump(fmt.Sprintf("AFTER Start second looking at pn %d", pn))
					spew.Dump("after")
					PrintGrid(bruteForcedGrid)
					// 	spew.Dump("AFTER end second")
					// }
					if err != nil {
						return nil, err
					}
					if completed {
						return bruteForcedGrid, nil
					}
					// try every other combination
					bruteForcedGrid = copyOriginalGrid(originalGrid)
				}
			}
		}
	}
	return originalGrid, nil
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

func copyOriginalGrid(grid [][]int) [][]int {
	tempGrid := make([][]int, len(grid))
	for i := range grid {
		tempGrid[i] = make([]int, len(grid[i]))
		copy(tempGrid[i], grid[i])
	}
	return tempGrid
}
