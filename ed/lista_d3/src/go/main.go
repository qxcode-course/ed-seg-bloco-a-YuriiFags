package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	Value int
	next  *Node
	prev  *Node
	root  *Node
}

type LList struct {
	root *Node
	size int
}

func NewLList() *LList {
	list := &LList{}
	list.root = &Node{root: nil}
	list.root.next = list.root
	list.root.prev = list.root
	list.root.root = list.root // nó sentinela aponta pra si mesmo
	return list
}

func (l *LList) String() string {
	values := []string{}

	for node := l.root.next; node != l.root; node = node.next {
		values = append(values, strconv.Itoa(node.Value))
	}

	return "[" + strings.Join(values, ", ") + "]"
}

func (l *LList) PushBack(value int) {
	l.insertBefore(l.root, value)
}

func (l *LList) insertBefore(mark *Node, value int) {
	n := &Node{
		Value: value,
		root:  l.root,
	}
	n.prev = mark.prev
	n.next = mark
	mark.prev.next = n
	mark.prev = n
	l.size++
}

func str2list(serial string) *LList {
	serial = serial[1 : len(serial)-1]
	ll := NewLList()
	if serial == "" {
		return ll
	}
	for _, p := range strings.Split(serial, ",") {
		value, _ := strconv.Atoi(p)
		ll.PushBack(value)
	}
	return ll
}

func equals(lla, llb *LList) bool {
	if lla.size != llb.size {
		return false
	}

	a := lla.root.next
	b := llb.root.next

	for a != lla.root && b != llb.root {
		if a.Value != b.Value {
			return false
		}

		a = a.next
		b = b.next
	}

	return true
}

func addsorted(ll *LList, value int) {
	node := ll.root.next

	for node != ll.root && node.Value < value {
		node = node.next
	}

	ll.insertBefore(node, value)
}

func reverse(ll *LList) *LList {
	newList := NewLList()

	for node := ll.root.prev; node != ll.root; node = node.prev {
		newList.PushBack(node.Value)
	}

	return newList
}

func merge(lla, llb *LList) *LList {
	newList := NewLList()

	a := lla.root.next
	b := llb.root.next

	for a != lla.root && b != llb.root {
		if a.Value <= b.Value {
			newList.PushBack(a.Value)
			a = a.next
		} else {
			newList.PushBack(b.Value)
			b = b.next
		}
	}

	for a != lla.root {
		newList.PushBack(a.Value)
		a = a.next
	}

	for b != llb.root {
		newList.PushBack(b.Value)
		b = b.next
	}

	return newList
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

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
		case "compare":
			lla := str2list(args[1])
			llb := str2list(args[2])
			if equals(lla, llb) {
				fmt.Println("iguais")
			} else {
				fmt.Println("diferentes")
			}
		case "addsorted":
			lla := NewLList()
			for i := 1; i < len(args); i++ {
				value, _ := strconv.Atoi(args[i])
				addsorted(lla, value)
			}
			fmt.Println(lla.String())
		case "reverse":
			lla := str2list(args[1])
			lla = reverse(lla)
			fmt.Println(lla.String())
		case "merge":
			lla := str2list(args[1])
			llb := str2list(args[2])
			merged := merge(lla, llb)
			fmt.Println(merged.String())
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
