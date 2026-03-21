package main

import "fmt"

func main() {
    var qtd, E int
    fmt.Scan(&qtd, &E)

    vivos := make([]bool, qtd)

    for i := 0; i < qtd; i++ {
        vivos[i] = true
    }

    contagemVivos := qtd
    pos := E - 1

   
    fmt.Print("[ ")
    for i := 0; i < qtd; i++ {
        if i == pos {
            fmt.Printf("%d> ", i+1)
        } else {
            fmt.Printf("%d ", i+1)
        }
    }
    fmt.Println("]")

    for contagemVivos > 1 {
        prox := (pos + 1) % qtd

        for vivos[prox] == false {
            prox = (prox + 1) % qtd
        }

        vivos[prox] = false
        contagemVivos--

        pos = (prox + 1) % qtd
        for vivos[pos] == false {
            pos = (pos + 1) % qtd
        }

        fmt.Print("[ ")
        for i := 0; i < qtd; i++ {
            if vivos[i] {
                if i == pos {
                    fmt.Printf("%d> ", i+1)
                } else {
                    fmt.Printf("%d ", i+1)
                }
            }
        }
        fmt.Println("]")
    }
}