package main

import (
	"bufio"
	"fmt"
	"os"
)

// Não mude a assinatura desta função, ela é a função chamada pelo LeetCode
func exist(grid [][]byte, word string) bool {
	lin := len(grid)
	col := len(grid[0])


	var dfs func(int, int, int) bool
	 
	dfs = func(l, c, idx int) bool {
		if idx == len(word) {
			return true
		}
		if l < 0 || l >= lin || c < 0 || c >= col{
			return false
		}
		if grid[l][c] != word[idx] {
			return false
		}
		t := grid[l][c]
		grid[l][c] = '#'

		achou := dfs(l+1,c,idx+1)||
		dfs(l-1,c,idx+1)||
		dfs(l,c+1,idx+1)||
		dfs(l,c-1, idx+1)

		grid[l][c] = t
		return achou
	}
	for i := 0; i < lin;i++ {
		for j := 0; j < col; j++ {
			if dfs(i,j,0){
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
