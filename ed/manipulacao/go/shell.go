package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	
)

func getMen(vet []int) []int {
	var homens [] int
	for _, valor := range vet {
		if valor > 0 {
			homens = append(homens, valor)
		}
	}
	return homens
}

func getCalmWomen(vet []int) []int {
	var mulheres[] int 
	for _, valor := range vet{
		if valor < 0 && -valor < 10 {
			mulheres = append(mulheres, valor)
		}
	}
	return mulheres
}

func sortVet(vet []int) []int {
	 var novo []int

	 for _, valor := range vet {
		novo = append(novo, valor)
	 }
	 sort.Ints(novo)
	 return novo
}

func sortStress(vet []int) []int {
	nivel := append([]int{}, vet...)

	sort.Slice(nivel, func(i,j int)bool{
		return abs(nivel[i]) < abs(nivel[j])
	})
	return nivel
}

func abs( x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func reverse(vet []int) []int {
	var invertido []int 
	for i := len(vet) - 1; i >= 0; i-- {
		invertido = append(invertido, vet[i])
	} 
	return invertido
}

func unique(vet []int) []int {
	visto := make(map[int]bool)
	var resultado [] int

	for _, valor := range vet {
		if !visto[valor] {
			resultado = append(resultado, valor)
			visto[valor] = true
		}
	}
	return resultado
}

func repeated(vet []int) []int {
	visto := make(map[int]int)
	var repetidos [] int

	for _, valor := range vet {
		visto[valor]++

		if visto[valor] > 1 {
			repetidos = append(repetidos, valor)
		}
	}
	return repetidos
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		if !scanner.Scan() {
			break
		}
		fmt.Print("$")
		line := scanner.Text()
		args := strings.Split(line, " ")
		fmt.Println(line)

		switch args[0] {
		case "get_men":
			printVec(getMen(str2vet(args[1])))
		case "get_calm_women":
			printVec(getCalmWomen(str2vet(args[1])))
		case "sort":
			printVec(sortVet(str2vet(args[1])))
		case "sort_stress":
			printVec(sortStress(str2vet(args[1])))
		case "reverse":
			array := str2vet(args[1])
			other := reverse(array)
			printVec(array)
			printVec(other)
		case "unique":
			printVec(unique(str2vet(args[1])))
		case "repeated":
			printVec(repeated(str2vet(args[1])))
		case "end":
			return
		}
	}
}

func printVec(vet []int) {
	fmt.Print("[")
	for i, val := range vet {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(val)
	}
	fmt.Println("]")
}

func str2vet(s string) []int {
	if s == "[]" {
		return nil
	}
	s = s[1 : len(s)-1]
	parts := strings.Split(s, ",")
	var vet []int
	for _, part := range parts {
		n, _ := strconv.Atoi(part)
		vet = append(vet, n)
	}
	return vet
}
