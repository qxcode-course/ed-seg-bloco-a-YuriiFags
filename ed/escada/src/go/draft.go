package main

import "fmt"


func calcularDegraus(num int) int {
    if num == 1 || num == 2 {
        return 1
    }

    if num == 3 {
        return 2
    }

    resultado := (calcularDegraus(num - 1)) + (calcularDegraus(num - 3))
    return(resultado)

}

func main() {
	var qtdDegraus int
	fmt.Scan(&qtdDegraus)
    fmt.Println(calcularDegraus(qtdDegraus))

}
