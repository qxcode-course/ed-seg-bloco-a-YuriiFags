package main

import (
	"bufio"
	"fmt"
	"os"
	
)


func podeColocar(seq []rune,pos int, valor rune, L int) bool {
    inicio := pos - L

    if inicio < 0{
        inicio = 0
    }

    fim := pos + L
    if fim >= len(seq) {
        fim = len(seq)-1
    }

    for i := inicio; i <= fim; i++ {
        if i == pos {
            continue
        }
        if seq[i] == valor {
            return false
        }
    }
    return true
}

func backtracking(seq[]rune, pos int, L int)bool {
    if pos == len(seq){
        return true
    }
    if seq[pos] != '.' {
        return backtracking(seq, pos+ 1, L)
    }
    for d := 0; d <= L; d++{
        valor := rune('0' + d)
        if podeColocar(seq,pos, valor,L){
            seq[pos] = valor
            if backtracking(seq,pos+1,L){
                return true
            }
            seq[pos] = '.'
        }
    }
    return false
}

func main(){
    in := bufio.NewReader(os.Stdin)
    var s string
    var L int

    fmt.Fscan(in, &s)
    fmt.Fscan(in,&L)

    seq := []rune(s)
    backtracking(seq,0,L)
    fmt.Println(string(seq))
}