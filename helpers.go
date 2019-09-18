package soduku

import "fmt"

func DrawGrid(box [][]int) {
	println("")
	for _, row := range box {
		line := ""
		for _, num := range row {
			line = fmt.Sprintf("%s | %d", line, num)
		}
		println(fmt.Sprintf("%s |", line))
	}
	println("")
}
