package main

import "fmt"

func imprimir(num[]int, pos int) {
    fmt.Print("[ ")
    for i, valor := range num {
        if i == pos {
            fmt.Printf("%v> ",valor)
        } else {
            fmt.Printf("%v ",valor)
        }
    }
    fmt.Println("]")
}

func main() {
   var qtd_Numeros, num_Escolhido int
   fmt.Scan(&qtd_Numeros,&num_Escolhido)
   num := make([]int,qtd_Numeros)

   for i := 0 ; i < qtd_Numeros; i++{
    num[i] = i + 1
   }

   pos_Atual := num_Escolhido - 1
   for len(num) > 0 {
    imprimir(num,pos_Atual)

    if len(num) == 1 {
        break
    }
   
    morto := (pos_Atual + 1) % len(num)
    num = append(num[:morto],num[morto + 1:]...)
    pos_Atual = morto % len(num)
   }
}