package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

)

func localizarSoma(num []int, k int) bool{
    var dfs func(int, int) bool
   dfs = func(i, soma int) bool {
        if soma == k {
            return true
        }
        if i == len(num) {
            return false
        }
        if dfs(i+1, soma + num[i]) {
            return true
        }
        if dfs(i+1,soma) {
            return true
        }
        return false
   }
   return dfs(0,0)
}

func main() {
   scanner := bufio.NewScanner(os.Stdin)

   scanner.Scan()
   linha1 := strings.Fields(scanner.Text())
   n, _ := strconv.Atoi(linha1[0])
   k, _ := strconv.Atoi(linha1[1])

   scanner.Scan()
   elementos := strings.Fields(scanner.Text())
   num := make([]int, n)

   for i := 0; i < n; i++ {
    num[i], _ = strconv.Atoi(elementos[i])
   }
   fmt.Println(localizarSoma(num, k))
}