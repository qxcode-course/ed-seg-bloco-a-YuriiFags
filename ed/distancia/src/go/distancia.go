package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func solve(sequence string, L int) string {
	s := []rune(sequence)
	n := len(s)

	// Função para verificar se é válido colocar digit na posição pos
	isValid := func(pos int, digit rune) bool {
		// Verificar posições anteriores
		for j := 0; j < pos; j++ {
			if s[j] == digit {
				// Distância deve ser > L
				if pos-j <= L {
					return false
				}
			}
		}
		// Verificar posições posteriores (já preenchidas)
		for j := pos + 1; j < n; j++ {
			if s[j] != '.' && s[j] == digit {
				// Distância deve ser > L
				if j-pos <= L {
					return false
				}
			}
		}
		return true
	}

	var backtrack func(int) bool
	backtrack = func(pos int) bool {
		if pos == n {
			return true
		}

		if s[pos] != '.' {
			// Já preenchido, continua
			return backtrack(pos + 1)
		}

		// Tenta cada dígito de 0 a L
		for digit := 0; digit <= L; digit++ {
			d := rune('0' + digit)
			if isValid(pos, d) {
				s[pos] = d
				if backtrack(pos + 1) {
					return true
				}
				s[pos] = '.'
			}
		}

		return false
	}

	backtrack(0)
	return string(s)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// Lê a sequência
	scanner.Scan()
	sequence := strings.TrimSpace(scanner.Text())

	// Lê o limite L
	scanner.Scan()
	L, _ := strconv.Atoi(strings.TrimSpace(scanner.Text()))

	// Resolve e imprime
	result := solve(sequence, L)
	fmt.Println(result)
}
