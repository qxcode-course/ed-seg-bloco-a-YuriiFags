package main

import (
	"bufio"
	"fmt"
	"os"
)

type Pos struct {
	l, c int
}

func labirinto(mat [][]rune, inicio, fim Pos) {
	caminho := NewStack[Pos]()
	becos := NewStack[Pos]()

	nl := len(mat)
	nc := len(mat[0])

	visitado := make([][]bool, nl)

	for i := range visitado {
		visitado[i] = make([]bool, nc)
	}

	dl := []int{1, -1, 0, 0}
	dc := []int{0, 0, -1, 1}

	caminho.Push(inicio)

	for !caminho.IsEmpty() {
		atual := caminho.Top()

		if !visitado[atual.l][atual.c] {
			visitado[atual.l][atual.c] = true
		}
		if atual == fim {
			break
		}
		achou := false

		for i := 0; i < 4; i++ {
			nLinha := atual.l + dl[i]
			nColuna := atual.c + dc[i]

			if nLinha < 0 || nLinha >= nl || nColuna >= nc || nColuna < 0 {
				continue
			}
			if visitado[nLinha][nColuna] {
				continue
			}
			if mat[nLinha][nColuna] == '#' {
				continue
			}
			caminho.Push(Pos{nLinha, nColuna})
			achou = true
			break
		}
		if !achou {
			becos.Push(atual)

			if mat[atual.l][atual.c] != 'I' &&
				mat[atual.l][atual.c] != 'F' {
				mat[atual.l][atual.c] = '.'
			}
			caminho.Pop()

		}
	}
	temp := NewStack[Pos]()
	for !caminho.IsEmpty() {
		p := caminho.Pop()

		if mat[p.l][p.c] != '#' {
			mat[p.l][p.c] = '.'
		}
	}
	for !temp.IsEmpty() {
		p := caminho.Pop()
		temp.Push(p)
	}
	for !temp.IsEmpty() {
		p := temp.Pop()
		if mat[p.l][p.c] != '#' && mat[p.l][p.c] != 'I' && mat[p.l][p.c] != 'F' {
			mat[p.l][p.c] = '.'
		}
		caminho.Push(p)
	}
	for !becos.IsEmpty() {
		p := becos.Pop()
		if mat[p.l][p.c] == '.' {
			mat[p.l][p.c] = ' '
		}
	}
	// mat[inicio.l][inicio.c] = 'I'
	// mat[fim.l][fim.c] = 'F'
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var nl, nc int

	scanner.Scan()
	fmt.Sscanf(scanner.Text(), "%d %d", &nl, &nc)

	mat := make([][]rune, nl)

	var inicio, fim Pos

	for i := 0; i < nl; i++ {
		scanner.Scan()
		linha := []rune(scanner.Text())

		mat[i] = linha

		for j := 0; j < nc; j++ {
			if linha[j] == 'I' {
				inicio = Pos{i, j}
			}
			if linha[j] == 'F' {
				fim = Pos{i, j}
			}
		}
	}
	labirinto(mat, inicio, fim)
	for _, linha := range mat {
		fmt.Println(string(linha))
	}
}
