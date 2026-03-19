package main
import (
    "fmt"
    "sort"
)
func main() {
    var total_figurinhas int
    var fig_Baruel int

    fmt.Scan(&total_figurinhas)
    fmt.Scan(&fig_Baruel)
    contagem := make(map[int]int)

    for i := 0; i < fig_Baruel; i++ {
        var num int
        fmt.Scan(&num)
        contagem[num]++
    }

    repetidas := []int{}

    for num,qtd := range contagem {
        if qtd > 1 {
            for i := 0; i < qtd - 1; i++ {
                repetidas = append(repetidas, num)
            }
        }
    }

sort.Ints(repetidas)
    if len(repetidas)== 0{
        fmt.Print("N")
    } else {
        for i := 0; i < len(repetidas); i++ {
            if i > 0 {
                fmt.Print(" ")
            }
            fmt.Print(repetidas[i])
        }
    }

    fmt.Println()
   
    faltantes := []int{}

    for i := 1 ; i <= total_figurinhas; i++ {
        if contagem[i] == 0 {
           faltantes = append(faltantes, i)
        }
    }

  if len(faltantes) == 0 {
    fmt.Print("N")
  } else {
    for i := 0; i < len(faltantes); i++ {
        if i > 0 {
            fmt.Print(" ")
        }
        fmt.Print(faltantes[i])
    }
  }

    fmt.Println()
}
