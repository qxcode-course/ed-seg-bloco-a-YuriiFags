package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Join(slice []int, sep string) string {
	if len(slice) == 0 {
		return ""
	}
	result := fmt.Sprintf("%d", slice[0])
	for _, value := range slice[1:] {
		result += sep + fmt.Sprintf("%d", value)
	}
	return result
}

type MultiSet struct {
	data []int
}

func NewMultiSet(capacidade int) *MultiSet {
	return &MultiSet{
		data: make([]int, 0, capacidade),
	}
}

func (m *MultiSet) encontraIndice(val int) int {
	posInicial := 0
	posFinal := len(m.data)

	for posInicial < posFinal {
		meio := (posInicial + posFinal) / 2

		if m.data[meio] < val {
			posInicial = meio + 1
		} else {
			posFinal = meio
		}
	}
	return posFinal
}

func (m *MultiSet) insert(val int) {
	pos := m.encontraIndice(val)

	m.data = append(m.data, 0)

	copy(m.data[pos+1:], m.data[pos:])
	m.data[pos] = val
}

func (m*MultiSet) Contains(val int) bool {
	pos := m.encontraIndice(val)

	if pos < len(m.data) && m.data[pos] == val{
		return true
	}
	return false
}

func (m*MultiSet) Show() string {
	if len(m.data) == 0 {
		return "[]"
	}

	saida := "["

	for i,v := range m.data {
		if i > 0 {
			saida += ", "
		}
		saida += fmt.Sprint(v) 
	}
	saida += "]"
	return saida
}

func (m*MultiSet) Erase(val int){
	pos := m.encontraIndice(val)

	if pos > len(m.data) ||  m.data[pos] != val {
		fmt.Println("value not found")
	}

	for i := pos; i < len(m.data) - 1; i++ {
		m.data[i] = m.data[i+1]
	}
	m.data = m.data[:len(m.data)-1]
}

func (m*MultiSet) Count(val int) int {
	pos := m.encontraIndice(val)

	if pos > len(m.data) {
		return 0
	}
	contar := 0
	for i := pos; i < len(m.data);i++{
		if m.data[i] == val {
			contar++
		}
	}
	return contar
}

func (m*MultiSet) Unique() int {

	if len(m.data) == 0 {
		return 0
	}
	contar := 1
	for i := 0; i < len(m.data) - 1; i++{
		if m.data[i] != m.data[i+1]{
			contar++
		}
	}
	return contar
}

func (m*MultiSet) Clear() {
	m.data = m.data[:0]
}
func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)
	ms := NewMultiSet(0)
	_ = ms

	for scanner.Scan() {
		fmt.Print("$")
		line = scanner.Text()
		args := strings.Fields(line)
		fmt.Println(line)
		if len(args) == 0 {
			continue
		}
		cmd = args[0]

		switch cmd {
		case "end":
			return
		case "init":
			value, _ := strconv.Atoi(args[1])
			ms = NewMultiSet(value)
		case "insert":
			for _, part := range args[1:] {
				value, _ := strconv.Atoi(part)
				ms.insert(value)
			}
		case "show":
			fmt.Println(ms.Show())
		case "erase":
			value, _ := strconv.Atoi(args[1])
			ms.Erase(value)
		case "contains":
			value, _ := strconv.Atoi(args[1])
			fmt.Println(ms.Contains(value))
		case "count":
			value, _ := strconv.Atoi(args[1])
			fmt.Println(ms.Count(value))
		case "unique":
			fmt.Println(ms.Unique())
		case "clear":
			ms.Clear()
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
