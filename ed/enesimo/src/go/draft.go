package main

import "fmt"

// x: número que está sendo testado
// div: divisor que está sendo testado
func eh_primo(x int, div int) bool {
	if x <= 1 {
		return false
	}

	if x == div {
		return true
	}
	if x%div == 0 {
		return false
	}
	return eh_primo(x, div+1)

}

func nEnesimo(n int, numeroAtual int) int {
	if eh_primo(numeroAtual,2){
		if n == 1 {
			return numeroAtual
		}
		return nEnesimo(n-1, numeroAtual+1)
	}
	return nEnesimo(n, numeroAtual + 1)

}
func main() {
	var x int
	fmt.Scan(&x)
	
	fmt.Println(nEnesimo(x,2))
}
