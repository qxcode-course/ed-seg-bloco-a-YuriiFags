package main
import "fmt"
func main() {
    var s string
    fmt.Scan(&s)

    pilha := []rune{}
    balanceado := true

    for _, c := range s {
        switch c {
        case '(', '[':
            pilha = append(pilha, c)
        case ')':
            if len(pilha) == 0 || pilha[len(pilha)-1] != '('{
                balanceado = false
                break
            }
            pilha = pilha[:len(pilha)-1]
        case ']':
            if len(pilha) == 0 || pilha[len(pilha)-1] != '['{
                balanceado = false
                break
            }
            pilha = pilha[:len(pilha)-1]
        }
        if !balanceado{
            break
        }
    }
    if balanceado && len(pilha) == 0 {
        fmt.Println("balanceado")
    } else {
        fmt.Println("nao balanceado")
    }
}