package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func longestIncreasingPath(matrix [][]int) int {
	linha := len(matrix)
	coluna := len(matrix[0])

	memo := make([][]int, linha)

	for i := range memo {
		memo[i] = make([]int, coluna)
	}
	direcoes := [][]int {
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
	}
	var dfs func(int, int) int

	dfs = func(l, c int) int {
		if memo[l][c] != 0 {
			return memo[l][c]
		}
		melhor := 1
		for _, d := range direcoes {
			nl := l + d[0]
			nc := c + d[1]

			if nl >= 0 && nl < linha && nc >= 0 && nc < coluna && matrix[nl][nc] > matrix[l][c] {
				melhor = max(melhor, 1 + dfs(nl,nc))
			}
		}
		memo[l][c] = melhor
		return melhor
	}
	ans := 0
	for i := 0; i < linha; i++ {
		for j := 0; j < coluna; j++ {
			ans = max(ans, dfs(i,j))
		}
	}
	return ans
}


func max(a,b int) int {
	if a > b {
		return a
	}
	return b
}

// Não modifique a função main
func main() {
	scanner := bufio.NewScanner(os.Stdin)

	if !scanner.Scan() {
		return
	}
	parts := strings.Fields(scanner.Text())
	if len(parts) < 2 {
		return
	}
	nl, _ := strconv.Atoi(parts[0])
	nc, _ := strconv.Atoi(parts[1])

	matrix := make([][]int, nl)
	for i := 0; i < nl; i++ {
		if !scanner.Scan() {
			return
		}
		tokens := strings.Fields(scanner.Text())
		row := make([]int, nc)
		for j := 0; j < nc && j < len(tokens); j++ {
			v, _ := strconv.Atoi(tokens[j])
			row[j] = v
		}
		matrix[i] = row
	}

	fmt.Println(longestIncreasingPath(matrix))
}
