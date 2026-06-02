package main

import (
	"fmt"
)

var (
	n      int
	count  int
	cols   []bool
	diag1  []bool
	diag2  []bool
)

func solve(row int) {
	if row == n {
		count++
		return
	}

	for col := 0; col < n; col++ {
		d1 := row - col + (n - 1)
		d2 := row + col

		if !cols[col] && !diag1[d1] && !diag2[d2] {
			cols[col] = true
			diag1[d1] = true
			diag2[d2] = true

			solve(row + 1)

			cols[col] = false
			diag1[d1] = false
			diag2[d2] = false
		}
	}
}

func main() {
	if _, err := fmt.Scan(&n); err != nil {
		return
	}

	cols = make([]bool, n)
	diag1 = make([]bool, 2*n-1)
	diag2 = make([]bool, 2*n-1)

	solve(0)

	fmt.Println(count)
}

