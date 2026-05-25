package main
import (
    "fmt"
    "sort"
      )
    

func Permutacao(s string) []string{
    chars := []rune(s)
    n := len(chars)

    usados := make([]bool,n)
    atual := make([]rune,0,n)

    resultado := []string{}

    var backtracking func() 
    backtracking = func() {
        if len(atual) == n {
            resultado = append(resultado, string(atual))
            return 
        }

        for i := 0; i < n; i++ {
            if !usados[i] {
                usados[i] = true
                atual = append(atual, chars[i])

                backtracking()

                atual = atual[:len(atual)-1]
                usados[i] = false
            }
        }
    }

    backtracking()
    
    sort.Strings(resultado)
    return resultado

}


func main() {
	var s string
	fmt.Scanln(&s)
	
	permutations := Permutacao(s)
	
	
	for _, perm := range permutations {
		fmt.Println(perm)
	}
}