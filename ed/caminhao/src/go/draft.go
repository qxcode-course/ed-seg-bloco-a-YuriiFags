package main
import "fmt"
func main() {
    var n int

    fmt.Scan(&n)

    total := 0
    tanque := 0 
    inicio := 0

    for i := 0; i < n; i++ {
        var gas, dist int
        fmt.Scan(&gas, &dist)

        saldo := gas - dist
        total += saldo
        tanque += saldo

        if tanque < 0 {
            inicio = i + 1
            tanque = 0
        }
    }
    if total < 0 {
        fmt.Println(-1)
    } else {
        fmt.Println(inicio)
    }
}