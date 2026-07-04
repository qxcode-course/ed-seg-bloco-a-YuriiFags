package main

import (
	"bufio"
	"fmt"
	"os"
)

type Pos struct{
	l,c int
}
func burnTrees(grid [][]rune, l, c int) {
	stack := NewStack[Pos]()
	
	stack.Push(Pos{l,c}) // posição inicial

	nl := len(grid)
	nc := len(grid[0])

	for !stack.IsEmpty() {
		atual := stack.Pop()

		linha := atual.l
		coluna := atual.c

		if linha < 0 || linha >= nl || coluna < 0 || coluna >= nc {
			continue
		}

		if grid[linha][coluna] != '#' {
			continue
		}
		grid[linha][coluna]  = 'o'

		stack.Push(Pos{linha - 1,  coluna})
		stack.Push(Pos{linha + 1, coluna})
		stack.Push(Pos{linha, coluna - 1})
		stack.Push(Pos{linha, coluna + 1})
	}

	// Essa função deve usar uma list como pilha
	// e marcar as árvores na matriz como queimados
	// Uma sugestão de como fazer isso é:
	// - adicionar a primeira posição na pilha
	// - enquanto a pilha não estiver vazia:
	//   - retirar o elemento do topo
	//   - se puder ser queimado, queime e adicione seus vizinhos à pilha

}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	var nl, nc, lfire, cfire int
	fmt.Sscanf(line, "%d %d %d %d", &nl, &nc, &lfire, &cfire)

	grid := make([][]rune, 0, nl)
	for range nl {
		scanner.Scan()
		line := []rune(scanner.Text())
		grid = append(grid, line)
	}
	burnTrees(grid, lfire, cfire)
	showGrid(grid)
}

func showGrid(mat [][]rune) {
	for _, linha := range mat {
		fmt.Println(string(linha))
	}
}
