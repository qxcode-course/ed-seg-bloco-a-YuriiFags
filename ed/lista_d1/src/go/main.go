package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct{
	Value int
	next *Node
	prev *Node
}
type LList struct{
	root *Node
}

func NewLList() *LList {
	return &LList{
		root: nil,
	}
}

func (ll *LList) String() string {
	saida := "["

	if ll.root == nil {
		return "[]"
	}
	atual := ll.root

	for atual != nil {
		saida += strconv.Itoa(atual.Value)
		if atual.next != nil {
			saida += ", "
		}
		atual = atual.next
	}
	saida += "]"
	return saida
}

func (ll *LList) Size() int {
	contar := 0
	atual := ll.root

	for atual != nil {
		contar++
		atual = atual.next
	}
	return contar 
}

func (ll *LList) PushFront(val int){
	novoNo := &Node {
		Value: val, next: nil, prev: nil,
	}
	if ll.root == nil {
		ll.root = novoNo
	} else {
		novoNo.next = ll.root
		ll.root.prev = novoNo
		ll.root = novoNo
	}
}

func(ll *LList) Clear(){
	ll.root = nil
}

func(ll *LList) PushBack(val int) {
	novoNo := &Node{Value: val, prev: nil, next: nil}

	if ll.root == nil {
		ll.root = novoNo
		return
	}
	atual := ll.root
	for atual.next != nil {
		atual = atual.next
	}
	atual.next = novoNo
}

func(ll *LList) PopBack() {
	if ll.root == nil || ll.root.next == nil {
		ll.root = nil
		return
	}
	v := ll.root

	for v.next.next != nil {
		v = v.next
	}
	v.next = nil
}

func(ll *LList) PopFront(){
	if ll.root == nil {
		return
	}
	ll.root = ll.root.next
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
			ll.PopBack()
		case "pop_front":
			ll.PopFront()
		case "clear":
			ll.Clear()
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
