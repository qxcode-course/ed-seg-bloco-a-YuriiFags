package main

import (
	"bufio"
	"fmt"
	"os"


)

type Pos struct {
	l, c int
}

func (p Pos) getNeig() []Pos {
	return []Pos {
		{p.l - 1, p.c},
		{p.l + 1, p.c},
		{p.l, p.c -1},
		{p.l, p.c + 1},
	}
}

func inside(grid [][]rune, pos Pos) bool {
	nrows := len(grid)
	ncols := len(grid[0])
	return pos.l >= 0 && pos.l < nrows && pos.c >= 0 && pos.c < ncols
}

func match(grid [][]rune, pos Pos, char rune) bool {
	return inside(grid, pos) && grid[pos.l][pos.c] == char
}

func search(grid [][]rune, startPos Pos, endPos Pos) {
	q := NewQueue[Pos]()
	visited := make(map[Pos]bool)
	parent := make(map[Pos]Pos)

	q.Enqueue(startPos)
	visited[startPos] = true

	for !q.IsEmpty(){
		atual, _ := q.Dequeue()
		if atual == endPos {
			break
		}
		for _, vis := range atual.getNeig(){
			if !match(grid, vis, ' '){
				continue
			}
			if visited[vis] {
				continue
			}
			visited[vis] = true
			parent[vis] = atual
			q.Enqueue(vis)
		}
	}
	for p := endPos; p != startPos; p = parent[p] {
		grid[p.l][p.c] = '.'
	}
	grid[startPos.l][startPos.c] = '.'

}



func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var nl, nc int
	scanner.Scan()
	line := scanner.Text()
	fmt.Sscanf(line, "%d %d", &nl, &nc)
	mat := make([][]rune, nl) // Inicializa a matriz de runes

	// Carregando matriz
	for i := range nl {
		scanner.Scan()
		line := scanner.Text()
		mat[i] = []rune(line)
	}

	var inicio, fim Pos

	// Procurando inicio e fim e colocando ' ' nas posições iniciais
	for l := range nl {
		for c := range nc {
			if mat[l][c] == 'I' {
				mat[l][c] = ' '
				inicio = Pos{l, c}
			}
			if mat[l][c] == 'F' {
				mat[l][c] = ' '
				fim = Pos{l, c}
			}
		}
	}

	search(mat, inicio, fim)

	for _, line := range mat {
		fmt.Println(string(line)) // Converte o slice de runes de volta para string para imprimir
	}
}
