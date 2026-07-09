package main

import (
	"bufio"
	"fmt"
	"os"
)

// Função que será chamada no LeetCode
func countBattleships(board [][]byte) int {
	l := len(board)
	c := len(board[0])

	qtd := 0

	for i := 0; i < l; i++{
		for j := 0; j < c; j++{
			if board[i][j] == '.'{
				continue
			}
			if i > 0 && board[i-1][j] == 'X'{
				continue
			}
			if j > 0 && board[i][j-1] == 'X'{
				continue
			}
			qtd++
		}
	}
	return qtd
}

// Não modifique a função main
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	var nl, nc int
	fmt.Sscanf(line, "%d %d", &nl, &nc)

	board := make([][]byte, nl)
	for i := 0; i < nl; i++ {
		scanner.Scan()
		board[i] = []byte(scanner.Text())
	}

	result := countBattleships(board)
	fmt.Println(result)
}
