package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	value int
	next  *Node
	prev  *Node
}

type LList struct {
	root *Node
	size int
}

func NewLList() *LList {
	return &LList{
		root: nil,
		size: 0,
	}
}

func (ll *LList) String() string {
	saida := "["

	if ll.root == nil {
		return "[]"
	}

	atual := ll.root
	for atual != nil {
		saida += fmt.Sprintf("%v", atual.value)

		if atual.next != nil {
			saida += ", "
		}
		atual = atual.next
	}
	saida += "]"
	return saida

}

func (ll *LList) Size() int {
	return ll.size
}

func (ll *LList) PushBack(val int) {
	novoNo := &Node{
		value: val,
	}
	if ll.root == nil {
		ll.root = novoNo
		ll.size++
		return
	}
	atual := ll.root

	for atual.next != nil {
		atual = atual.next
	}
	atual.next = novoNo
	novoNo.prev = atual
	ll.size++

}

func (ll *LList) PushFront(val int) {
	novoNo := &Node{
		value: val,
	}

	if ll.root == nil {
		ll.root = novoNo
		ll.size++
		return
	} else {
		novoNo.next = ll.root
		ll.root.prev = novoNo
		ll.root = novoNo
		ll.size++
	}

}

func (ll *LList) Clear() {
	ll.root = nil
	ll.size = 0 
}

func (ll *LList) Front() *Node{
	return ll.root
}

func (ll *LList) Back() *Node {
	if ll.root == nil {
		return nil
	}

	atual := ll.root

	for atual.next != nil {
		atual = atual.next
	}
	return atual
}

func (ll *LList) replace(valorAntigo, valorNovo int) {
	atual := ll.root

	for atual != nil {
		if atual.value == valorAntigo {
			atual.value = valorNovo
			return
		}
		atual = atual.next
	}

}

func (ll *LList) Search(valorAntigo int) *Node {
	atual := ll.root

	for atual != nil {
		if atual.value == valorAntigo {
			return atual
		}
		atual = atual.next
	}
	return nil
}

func (ll *LList) Insert(node *Node, valorNovo int) {
		novoNo:= &Node {
		value: valorNovo,
	}

	novoNo.next = node
	novoNo.prev = node.prev

	if node.prev != nil {
		node.prev.next = novoNo
	} else {
		ll.root = novoNo
	}

	node.prev = novoNo

	ll.size++
}

func (ll *LList) Remove(node *Node){
	
	if node.prev != nil {
		node.prev.next = node.next
	} else {
		ll.root = node.next
	}

	if node.next != nil {
		node.next.prev = node.prev
	}
	ll.size--
}
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	ll := NewLList()

	for {
		fmt.Print("$")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		fmt.Println(line)
		args := strings.Fields(line)

		if len(args) == 0 {
			continue
		}

		cmd := args[0]

		switch cmd {
		case "show":
			fmt.Println(ll.String())
		case "size":
			fmt.Println(ll.Size())
		case "push_back":
			for _, v := range args[1:] {
				num, _ := strconv.Atoi(v)
				ll.PushBack(num)
			}
		case "push_front":
			for _, v := range args[1:] {
				num, _ := strconv.Atoi(v)
				ll.PushFront(num)
			}
		case "pop_back":
			// ll.PopBack()
		case "pop_front":
			// ll.PopFront()
		case "clear":
			ll.Clear()
		case "walk":
			fmt.Print("[ ")
			for node := ll.Front(); node != nil; node = node.next {
				fmt.Printf("%v ", node.value)
			}
			fmt.Print("]\n[ ")
			for node := ll.Back(); node != nil; node = node.prev {
				fmt.Printf("%v ", node.value)
			}
			fmt.Println("]")
		case "replace":
			oldvalue, _ := strconv.Atoi(args[1])
			newvalue, _ := strconv.Atoi(args[2])
			node := ll.Search(oldvalue)
			if node != nil {
				node.value = newvalue
			} else {
				fmt.Println("fail: not found")
			}
		case "insert":
			oldvalue, _ := strconv.Atoi(args[1])
			newvalue, _ := strconv.Atoi(args[2])
			node := ll.Search(oldvalue)
			if node != nil {
				ll.Insert(node, newvalue)
			} else {
				fmt.Println("fail: not found")
			}
		case "remove":
			oldvalue, _ := strconv.Atoi(args[1])
			node := ll.Search(oldvalue)
			if node != nil {
				ll.Remove(node)
			} else {
				fmt.Println("fail: not found")
			}
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
