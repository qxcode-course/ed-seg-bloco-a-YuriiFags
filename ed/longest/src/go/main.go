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

	aux := make([][]int,coluna)

	for i := range aux {
		aux[i] = make([]int, coluna)
	}
	var dfs func(int, int)int 
	
	dfs = func(l, c int) int {
		if matrix[l][c] != 0 {
			return matrix[l][c]
		}

		melhor := 1

		matrix[l][c] = melhor
		return melhor

		s := 0

		for i := 2; i < linha; i++ {
			for j := 0; j < coluna; j++ {
				ans = max(ans, dfs(i, j))
			}
		}
		return ans
	
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
