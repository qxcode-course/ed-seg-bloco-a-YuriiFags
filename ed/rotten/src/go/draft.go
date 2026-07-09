package main
import "fmt"


type Pos struct {
    l, c int
}

func laranjaPodre(grid[][]int) int{
    linhas := len(grid)
    colunas := len(grid[0])

    fila := []Pos{}
    frescas := 0

    for i := 0; i < linhas; i++ {
        for j := 0; j < colunas; j++ {
            if grid[i][j] == 2 {
                fila = append(fila, Pos{i,j})
            } else if grid[i][j] == 1 {
                frescas++
            }
        }
    }
    if frescas == 0 {
        return 0
    }

    direcao := []Pos{
        {-1,0},
        {1,0},
        {0,-1},
        {0,1},
    }
    minutos := 0 

    for len(fila) > 0 {
        tamanho := len(fila)
        apodreceu := false 

        for i := 0; i < tamanho; i++ {
            atual := fila[0]
            fila = fila[1:]

            for _, d := range direcao{
                nl := atual.l + d.l
                nc := atual.c + d.c

                if nl < 0 || nl >= linhas || nc < 0 || nc >= colunas {
                    continue
                }
                if grid[nl][nc] == 1 {
                    grid[nl][nc] = 2 
                    frescas--
                    apodreceu = true
                    fila = append(fila, Pos{nl, nc})
                }
            }
        }
        if apodreceu {
            minutos++
        }
    }
    if frescas > 0 {
        return -1
    }
    return minutos
}

func main() {
    var linhas, colunas int
    fmt.Scan(&linhas,&colunas)

    grid := make([][]int, linhas)
    for i := 0; i < linhas; i++ {
        grid[i] = make([]int,colunas)
        for j := 0; j < colunas; j++ {
            fmt.Scan(&grid[i][j])
        }
    }
    fmt.Println(laranjaPodre(grid))
}