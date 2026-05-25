package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func podColocarDígito(s []rune, posição int, digito rune, L int) bool {
   
    for i := 0; i < posição; i++ {
        if s[i] == digito {
            distância := posição - i
            if distância <= L {
                return false
            }
        }
    }

   
    for i := posição + 1; i < len(s); i++ {
        if s[i] != '.' && s[i] == digito {
            distância := i - posição
            if distância <= L {
                return false
            }
        }
    }

    return true
}

func preencherString(s []rune, posição int, L int) bool {
    
    if posição == len(s) {
        return true
    }

   
    if s[posição] != '.' {
        return preencherString(s, posição+1, L)
    }

    
    for dígito := 0; dígito < L; dígito++ {
        caractere := rune('0' + dígito)

        if podColocarDígito(s, posição, caractere, L) {
            s[posição] = caractere

           
            if preencherString(s, posição+1, L) {
                return true
            }

            
            s[posição] = '.'
        }
    }

    return false
}

func mainDist() {
    scanner := bufio.NewScanner(os.Stdin)

    
    scanner.Scan()
    sequência := strings.TrimSpace(scanner.Text())

    
    scanner.Scan()
    L, _ := strconv.Atoi(strings.TrimSpace(scanner.Text()))

    
    caracteres := []rune(sequência)

   
    preencherString(caracteres, 0, L)

  
    resultado := string(caracteres)
    fmt.Println(resultado)
}