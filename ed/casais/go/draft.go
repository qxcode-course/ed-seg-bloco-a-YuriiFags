package main

import "fmt"

func encontrarCasal(vet []int) int {
	solteiros := make(map[int]int)
	casais := 0

	for _, animal := range vet {
		par := -animal

		if solteiros[par] > 0 {
			solteiros[par]--
			casais++
		} else {
			solteiros[animal]++
		}
	}
	return casais
}
func imprimir(casais int){
    fmt.Println(casais)
}

func main(){
    var N int
    
    fmt.Scan(&N)
    vet := make([]int, N)

    for i := 0; i < N; i++ {
        fmt.Scan(&vet[i])
    }
     resultado:= encontrarCasal(vet)
     imprimir(resultado)
}