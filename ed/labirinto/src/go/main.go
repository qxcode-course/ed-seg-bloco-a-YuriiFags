package main

import (
	"bufio"
	"fmt"
	"os"
)

type Pos struct {
	l int
	c int
}

func caminho(grid [][]rune, startPos , endPos Pos)bool {
	l, c := startPos.l, startPos.c 
	if l < 0 || l >= len(grid) || c < 0 || c >= len(grid[l]){
		return false
	}

	if grid[l][c] != ' ' {
		return false
	}
	grid[l][c] = '.'

	if l == endPos.l && c == endPos.c {
		return true
	}

	if caminho(grid, Pos{l + 1,c}, endPos) ||
	 caminho(grid, Pos{l - 1,c}, endPos) ||
	  caminho(grid, Pos{l, c+ 1}, endPos) ||
	   caminho(grid, Pos{l, c - 1 }, endPos) {
		return true
	}
	grid[l][c] = ' '
	return false
	
}
	//_, _, _ = grid, startPos, endPos
	
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	nl_nc := scanner.Text()
	var nl, nc int
	fmt.Sscanf(nl_nc, "%d %d", &nl, &nc)
	grid := make([][]rune, nl)

	// Lê a gridriz
	for i := range nl {
		scanner.Scan()
		grid[i] = []rune(scanner.Text())
	}

	// Procura posições de início e endPos e conserta para _
	var startPos, endPos Pos
	for l := range nl {
		for c := range nc {
			if grid[l][c] == 'I' {
				grid[l][c] = ' '
				startPos = Pos{l, c}
			}
			if grid[l][c] == 'F' {
				grid[l][c] = ' '
				endPos = Pos{l, c}
			}
		}
	}

	caminho(grid, startPos, endPos)

	// Imprime o labirinto final
	for _, line := range grid {
		fmt.Println(string(line))
	}
}
