package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func tostr(vet []int) string {
	resultado := "["
	for i, valor := range vet {
		resultado += fmt.Sprintf("%v", valor)

		if i != len(vet) - 1 {
			resultado += ", "
		}
	}
	resultado += "]"
	return resultado
}

func tostrrev(vet []int) string {
	resultado := "["
	
	for i := len(vet) - 1; i >= 0; i-- {
		resultado += fmt.Sprintf("%v",vet[i])

		if i != 0 {
			resultado += ", "
		}
	}
	resultado += "]"
	return resultado
}

// reverse: inverte os elementos do slice
func reverse(vet []int) {
	i := 0
	j := len(vet) -1

	for i < j {
		vet[i], vet[j] = vet[j], vet[i]
		i++
		j--
	}
}

// sum: soma dos elementos do slice
func sum(vet []int) int {
	total := 0

	for _, valor := range vet {
		total += valor
	}
	return  total
}

// mult: produto dos elementos do slice
func mult(vet []int) int {
	total := 1

	for _, valor := range vet {
		total *= valor
	}
	return total
}

// min: retorna o índice e valor do menor valor
// crie uma função recursiva interna do modelo
// var rec func(v []int) (int, int)
// para fazer uma recursão que retorna valor e índice
func min(vet []int) int {
    if len(vet) == 0 {
        return -1
    }

    var rec func(v []int) (int, int)

    rec = func(v []int) (int, int) {
        if len(v) == 1 {
            return v[0], 0
        }

        valMenor, idMenor := rec(v[1:])

        if v[0] <= valMenor {
            return v[0], 0
        }

        return valMenor, idMenor + 1
    }

    _, indice := rec(vet)
    return indice
}

func main() {
	var vet []int
	scanner := bufio.NewScanner(os.Stdin)
	for {
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		args := strings.Fields(line)
		fmt.Println("$" + line)

		switch args[0] {
		case "end":
			return
		case "read":
			vet = nil
			for _, arg := range args[1:] {
				if val, err := strconv.Atoi(arg); err == nil {
					vet = append(vet, val)
				}
			}
		case "tostr":
			fmt.Println(tostr(vet))
		case "torev":
			fmt.Println(tostrrev(vet))
		case "reverse":
			reverse(vet)
		case "sum":
			fmt.Println(sum(vet))
		case "mult":
			fmt.Println(mult(vet))
		case "min":
			fmt.Println(min(vet))
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
