package main

import "fmt"

func Rotacionar(vetor []int, rotacoes int)[]int {
	n := len(vetor)

	rotacoes %= n

	resultado := make([]int, n)

	for i := 0; i < n; i++ {
		novaPosicao := (i + rotacoes) % n
		resultado[novaPosicao] = vetor[i]
	}
    return resultado 
}

func main() {
    tamanho := 0
    fmt.Scan(&tamanho)

    numerosDeRotacoes := 0
    fmt.Scan(&numerosDeRotacoes)

    vetor := make([]int,tamanho)

    for i:=0;i<tamanho;i++{
        fmt.Scan(&vetor[i])
    }
    resultado := Rotacionar(vetor,numerosDeRotacoes)

    fmt.Print("[ ")

    for _, valor := range resultado{
        fmt.Print(valor," ")
    }
    
    fmt.Println("]")
}
