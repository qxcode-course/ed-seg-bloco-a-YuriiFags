package main

import (
	"fmt"
	
	
)

// mostra a lista com o elemento sword destacado
func ToStr(l *DList[int], sword *DNode[int]) string {
	saida := ""
	contar := 0

	n := l.root.next

	for n != l.root {
		if contar > 0 {
			saida += " "
		}
		if n == sword {
			saida += fmt.Sprintf(">%v<",n.Value)
		} else {
			saida += fmt.Sprint(n.Value)
		}
		n = n.next
		contar++
	}
	return "[" + saida + "]"
}

// move para frente na lista circular
func Next(l *DList[int], it *DNode[int]) *DNode[int] {
	if it == nil {
		return nil
	}

	next := it.Next()
	if next == l.root {
		return l.root.next
	}
	return next
}
func main() {
	var qtd, chosen int
	fmt.Scan(&qtd, &chosen)
	fmt.Println(qtd, chosen)
	l := NewDList[int]()
	for i := 1; i <= qtd; i++ {
		l.PushBack(i)
	}
	sword := l.Front()
	for range chosen - 1 {
		sword = Next(l, sword)
	}
	for range qtd - 1 {
		fmt.Println(ToStr(l, sword))
		l.Erase(Next(l, sword))
		sword = Next(l, sword)
	}
	fmt.Println(ToStr(l, sword))
}
