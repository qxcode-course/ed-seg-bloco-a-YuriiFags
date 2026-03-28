package main

import "fmt"

func validarLancamento(a, b int) bool {
	if a >= 10 && b >= 10 {
		return true
	}
	return false
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func pontuacao(a, b int) int {
	return abs(a - b)
}

func main() {
	var qtdCompetidores, A, B int
	fmt.Scan(&qtdCompetidores)

	melhorIndice := -1
	melhorPontuacao := 101

	for i := 0; i < qtdCompetidores; i++ {
		fmt.Scan(&A, &B)

		if validarLancamento(A, B) {
			Pontuacao := pontuacao(A, B)

			if Pontuacao < melhorPontuacao {
				melhorPontuacao = Pontuacao
				melhorIndice = i
			}
		}

	}
    if melhorIndice == -1 {
        fmt.Println("sem ganhador")
    }else {
        fmt.Println(melhorIndice)
    }

}
