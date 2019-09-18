# sudoku

Sudoku solver. Pass the module a slice of ints to represent a sudoku board and it will
try and solve it.

```
SolveGrid([][]int{
  []int{2, 0, 7, 0, 0, 6, 0, 0, 0},
  []int{0, 0, 0, 0, 3, 0, 2, 0, 6},
  []int{0, 5, 6, 0, 0, 2, 0, 4, 1},
  []int{1, 0, 0, 3, 0, 8, 7, 6, 0},
  []int{6, 0, 9, 0, 0, 0, 1, 0, 8},
  []int{0, 7, 4, 6, 0, 5, 0, 0, 3},
  []int{5, 8, 0, 7, 0, 0, 4, 1, 0},
  []int{9, 0, 1, 0, 5, 0, 0, 0, 0},
  []int{0, 0, 0, 1, 0, 0, 3, 0, 5},
})

// returns

[][]int{
  []int{2, 1, 7, 8, 4, 6, 5, 3, 9},
  []int{4, 9, 8, 5, 3, 1, 2, 7, 6},
  []int{3, 5, 6, 9, 7, 2, 8, 4, 1},
  []int{1, 2, 5, 3, 9, 8, 7, 6, 4},
  []int{6, 3, 9, 4, 2, 7, 1, 5, 8},
  []int{8, 7, 4, 6, 1, 5, 9, 2, 3},
  []int{5, 8, 3, 7, 6, 9, 4, 1, 2},
  []int{9, 4, 1, 2, 5, 3, 6, 8, 7},
  []int{7, 6, 2, 1, 8, 4, 3, 9, 5},
}
 ```

 This was done as a hack on a plane in a few hours. It is not efficient and it won't solve everything (but will solve as much as it can!).
