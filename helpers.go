package soduku

type emptyPosition struct {
	rowNumber int
	colNumber int
}

type emptyPositionAndPossibleNumbers struct {
	ep           emptyPosition
	possibleNums []int
}

func getEmptyPositionsAndPossibleNumbers(grid [][]int) ([]emptyPositionAndPossibleNumbers, error) {
	epNums := []emptyPositionAndPossibleNumbers{}
	for _, ep := range getEmptyPositions(grid) {
		reg, err := getRegion(ep)
		if err != nil {
			return epNums, err
		}
		pn, err := possibleNumbers(grid, ep, reg)
		if err != nil {
			return epNums, err
		}
		epNums = append(epNums, emptyPositionAndPossibleNumbers{
			ep:           ep,
			possibleNums: pn,
		})
	}
	return epNums, nil
}

// getEmptyPositions returns the positions of empty boxes
func getEmptyPositions(grid [][]int) []emptyPosition {
	eps := []emptyPosition{}

	// identify missing positions
	for rowNumber, row := range grid {
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

// possibleNumbers returns the numbers that can possibly placed into a given position
func possibleNumbers(grid [][]int, ep emptyPosition, reg region) ([]int, error) {
	possibleNumbers := map[int]bool{}
	for i := 1; i <= 9; i++ {
		possibleNumbers[i] = false
	}

	// check the row it is on
	for col := 0; col <= 8; col++ {
		if possibleNumbers[grid[ep.rowNumber][col]] {
			continue
		}
		possibleNumbers[grid[ep.rowNumber][col]] = true
	}

	// check the column it is in
	for row := 0; row <= 8; row++ {
		if possibleNumbers[grid[row][ep.colNumber]] {
			continue
		}
		possibleNumbers[grid[row][ep.colNumber]] = true
	}

	// Check the grid it is in
	for row := reg.minRowNumber; row <= reg.maxRowNumber; row++ {
		for col := reg.minColNumber; col <= reg.maxColNumber; col++ {
			if possibleNumbers[grid[row][col]] {
				continue
			}
			possibleNumbers[grid[row][col]] = true
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
