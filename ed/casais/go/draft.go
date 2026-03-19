package main
import "fmt"
func main() {
    var N int
    pares := 0

    fmt.Scan(&N)

    descasados := make(map[int]int)

    for i := 0; i < N; i++ {
        var animal int
        fmt.Scan(&animal)

        if descasados[-animal] > 0{
            descasados[-animal]--
            pares++
        } else {
            descasados[animal]++
        }
    }
    fmt.Println(pares)
    
}
