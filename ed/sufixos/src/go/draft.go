package main

import (
	"fmt"
	
) 


func verificarLetras(palavra string, n int) {
    if len(palavra) == n {
        return
    }
    fmt.Print(string(palavra[n]))
    verificarLetras(palavra, n + 1)
}

func imprimirContrario(palavra string, pos int) {
	if len(palavra) == pos {
		return
	}

    inicio := len(palavra) - 1 - pos
    verificarLetras(palavra,inicio)
    fmt.Println()
    imprimirContrario(palavra, pos + 1)

}
func main() {
	var palavra string
    fmt.Scan(&palavra)
    imprimirContrario(palavra,0)
}
