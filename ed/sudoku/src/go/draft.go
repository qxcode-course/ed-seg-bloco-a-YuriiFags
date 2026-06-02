package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	
	grid := make([][]rune, n)
	for i := 0; i < n; i++ {
		var line string
		fmt.Scan(&line)
		grid[i] = []rune(line)
	}

	
	solve(grid, 0, n)

	
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Printf("%c", grid[i][j])
		}
		fmt.Println()
	}
}

func solve(grid [][]rune, index int, n int) bool {
	if index == n*n {
		return true
	}

	row := index / n
	col := index % n

	
	if grid[row][col] != '.' {
		return solve(grid, index+1, n)
	}

	
	for num := 1; num <= n; num++ {
		numChar := rune('0' + num)

		if isValid(grid, row, col, numChar, n) {
			grid[row][col] = numChar

			if solve(grid, index+1, n) {
				return true
			}

			grid[row][col] = '.'
		}
	}

	return false
}

func isValid(grid [][]rune, row, col int, num rune, n int) bool {
	
	for j := 0; j < n; j++ {
		if grid[row][j] == num {
			return false
		}
	}

	
	for i := 0; i < n; i++ {
		if grid[i][col] == num {
			return false
		}
	}


	if n == 4 {
		quadrantSize := 2
		startRow := (row / quadrantSize) * quadrantSize
		startCol := (col / quadrantSize) * quadrantSize

		for i := 0; i < quadrantSize; i++ {
			for j := 0; j < quadrantSize; j++ {
				if grid[startRow+i][startCol+j] == num {
					return false
				}
			}
		}
	} else if n == 9 {
		quadrantSize := 3
		startRow := (row / quadrantSize) * quadrantSize
		startCol := (col / quadrantSize) * quadrantSize

		for i := 0; i < quadrantSize; i++ {
			for j := 0; j < quadrantSize; j++ {
				if grid[startRow+i][startCol+j] == num {
					return false
				}
			}
		}
	}

	return true
}