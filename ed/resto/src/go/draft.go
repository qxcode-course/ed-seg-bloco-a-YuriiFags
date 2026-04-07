package main

import "fmt"

func main() {
	var value int

	fmt.Scan(&value)

    empilhar(value)

}

func empilhar(value int) {
	if value == 0 {
		return
	}

	div := value / 2
	resto := value % 2

	empilhar(div)
	fmt.Printf("%v %v\n", div, resto)
}
