package main
import(
    "fmt"
    "math"
) 
func main() {
    var N_competidores int
    var melhor_Pontuacao = math.MaxFloat64
    var melhor_Indice = -1

    fmt.Scan(&N_competidores)

    for i:= 0; i < N_competidores;i++{
       var A,B int
        fmt.Scan(&A,&B)

       

        if A < 10 || B < 10 {
            continue
        }
        diferenca := math.Abs(float64(A-B))
        if diferenca < melhor_Pontuacao{
            melhor_Pontuacao = diferenca
            melhor_Indice = i
          
        }
    }
    if melhor_Indice == -1 {
        fmt.Println("sem ganhador")
    } else {
        fmt.Println(melhor_Indice)
    }
}
