package arrays

import (
	"fmt"
)

func sudoku2(grid [][]string) bool {
	for i := range grid{
		for j :=  range grid[i] {
			if grid[i][j] != "." {
				for k := 0; k < len(grid); k++ {
					if grid[i][k] == grid[i][j] && j != k {
						return false
					}
					if grid[k][j] == grid[i][j] && i != k {
						return false
					}
				}
			}
		}
	}
	for r := 0; r < len(grid); r += 3 {
		for c := 0; c < len(grid); c += 3 {

			for i := r; i < r + 3; i++ {
				for j := c; j < c + 3; j++ {
					if grid[i][j] != "."  {
						val := grid[i][j]
						fmt.Println(val)
						for a := r; a < r+3; a++ {
							for b := c; b < c+3; b++ {
								if a==i && b==j {
									continue
								}
								if val == grid[a][b]{
									return false
								}
								if a==r+2 && b==c+2 {
									val = grid[a][b]
								}
								if val != grid[a][b] {
									continue
								}
							}
						}
					}
				}
			}
		}
	}
	return true
}
