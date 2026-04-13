package main

import (
	"fmt"
)

func espaços(qtd int) {
	if qtd == 0 {
		return
	}
	fmt.Print(" ")
	espaços(qtd - 1)
}

func diagonal(palavra string, pos int) {

	if len(palavra) == pos {
		return
	}

	espaços(pos)
    fmt.Printf("%v\n",string(palavra[pos]))
    diagonal(palavra,pos + 1)
}

func main() {
	var palavra string
	fmt.Scan(&palavra)
	diagonal(palavra, 0)
}
