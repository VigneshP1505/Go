package binarytrees

import "fmt"

func ZigzagTraversalSkip(grid [][]int) {
	result := []int{}
	for i := 0; i < len(grid); i++ {
		if i%2 == 0 {
			for j := 0; j < len(grid[i]); j += 2 {
				result = append(result, grid[i][j])
			}
		} else {
			j := len(grid[i]) - 1
			if j%2 == 0 {
				j--
			}
			for ; j >= 0; j -= 2 {
				result = append(result, grid[i][j])
			}
		}
	}
	fmt.Println(result)
}
