package main

import "fmt"

func matchingStrings(strings []string, consultas []string) []int {

	frequencia := make(map[string]int)

	for _, palavra := range strings {
		frequencia[palavra]++
	}

	resultado := make([]int,0)

	for _, consulta := range consultas {
		resultado = append(resultado, frequencia[consulta])
	}
	return resultado
}

func main() {
	qtd := 0

    fmt.Scan(&qtd)
    strings := make([]string, qtd)

    for i:=0; i < qtd; i++ {
        fmt.Scan(&strings[i])
    }


    qtd1 := 0
    fmt.Scan(&qtd1)

    consultas := make([]string, qtd1)
    for i := 0; i < qtd1; i++ {
        fmt.Scan(&consultas[i])
    }
    resultado := matchingStrings(strings, consultas)
    
    for i, valor := range resultado {
        if i > 0{
            fmt.Print(" ")
        }
        fmt.Print(valor)
    }
    fmt.Println()
}
