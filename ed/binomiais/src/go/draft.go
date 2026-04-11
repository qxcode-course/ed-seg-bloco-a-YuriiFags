package main
import "fmt"


func calcularCoeficiente(n,k int) int {
    if k == 0 {
        return 1
    }
    if k == n {
        return 1
    }

    if k >= 1 && k <= n - 1 {
        return calcularCoeficiente(n-1,k-1) + calcularCoeficiente(n-1,k)
    }
    return calcularCoeficiente(n,k)
}

func main() {
    var n,k int
    fmt.Scan(&n,&k)

    fmt.Println(calcularCoeficiente(n,k))
}
