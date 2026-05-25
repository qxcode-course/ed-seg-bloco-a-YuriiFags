package main

import (
	"bufio"
	"fmt"
	"os"
)


func exist(grid [][]byte, word string) bool {

	linhas := len(grid)
	colunas := len(grid[0])

	var dfs func(int, int, int) bool

	dfs = func(l, c, idx int) bool {

		
		if idx == len(word) {
			return true
		}

		
		if l < 0 || l >= linhas || c < 0 || c >= colunas {
			return false
		}


		if grid[l][c] != word[idx] {
			return false
		}

		// salva letra original
		temp := grid[l][c]

		// marca como visitado
		grid[l][c] = '#'

		
		encontrou :=
			dfs(l+1, c, idx+1) || // baixo
				dfs(l-1, c, idx+1) || // cima
				dfs(l, c+1, idx+1) || // direita
				dfs(l, c-1, idx+1) // esquerda

		// desfaz marcação
		grid[l][c] = temp

		return encontrou
	}

	// percorre toda matriz
	for i := 0; i < linhas; i++ {
		for j := 0; j < colunas; j++ {

			if dfs(i, j, 0) {
				return true
			}
		}
	}

	return false
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var word string
	fmt.Sscanf(scanner.Text(), "%s", &word)
	grid := make([][]byte, 0)
	for scanner.Scan() {
		grid = append(grid, []byte(scanner.Text()))
	}
	if exist(grid, word) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}