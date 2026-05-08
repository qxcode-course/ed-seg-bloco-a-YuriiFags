package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func MagicSearch(slice []int, value int) int {
	inicio := 0
	fim := len(slice)

	for inicio < fim {
		meio := (inicio+fim)/2
		if slice[meio] == value {
			index := meio

			for index + 1 < len(slice) && slice[index+1] == value {
				index++
			}
			return index
		} else if slice[meio] < value {
			inicio = meio + 1
		} else {
			fim = meio
		}
	}
	return inicio
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	parts := strings.Fields(scanner.Text())
	slice := make([]int, 0, 1)
	for _, elem := range parts[1 : len(parts)-1] {
		value, _ := strconv.Atoi(elem)
		slice = append(slice, value)
	}

	scanner.Scan()
	value, _ := strconv.Atoi(scanner.Text())
	result := MagicSearch(slice, value)
	fmt.Println(result)
}
