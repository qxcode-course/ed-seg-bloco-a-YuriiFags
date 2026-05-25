package main

import (
	"bufio"
	"fmt"
	"os"
)

func dfs(grid [][]byte, i, j int) {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) {
		return
	}
	if grid[i][j] != 'O' {
		return
	}

	grid[i][j] = 'V'

	dfs(grid, i-1, j)
	dfs(grid, i+1, j)
	dfs(grid, i, j+1)
	dfs(grid, i, j-1)
}

// NÃO ALTERE A ASSINATURA DA FUNÇÃO solve
func solve(board [][]byte) {
	if len(board) == 0 {
		return
	}

	// percorre bordas laterais
	for i := 0; i < len(board); i++ {
		if board[i][0] == 'O' {
			dfs(board, i, 0)
		}

		if board[i][len(board[0])-1] == 'O' {
			dfs(board, i, len(board[0])-1)
		}
	}

	// percorre bordas superior e inferior
	for j := 0; j < len(board[0]); j++ {
		if board[0][j] == 'O' {
			dfs(board, 0, j)
		}

		if board[len(board)-1][j] == 'O' {
			dfs(board, len(board)-1, j)
		}
	}

	// transforma os cercados
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {

			if board[i][j] == 'O' {
				board[i][j] = 'X'
			} else if board[i][j] == 'V' {
				board[i][j] = 'O'
			}
		}
	}
}

// NÃO ALTERE A MAIN
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var nrows, ncols int
	fmt.Sscanf(scanner.Text(), "%d %d", &nrows, &ncols)
	board := make([][]byte, nrows)
	for i := 0; i < nrows; i++ {
		scanner.Scan()
		board[i] = []byte(scanner.Text())
	}
	solve(board)
	for _, row := range board {
		fmt.Println(string(row))
	}
}
