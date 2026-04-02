package main
import "fmt"

func figurinhasRepetidas(figs []int) string {
	repetidas := make(map[int]int)
	saida := ""

	
	for _, valor := range figs {
		repetidas[valor]++
	}

	
	for numero, quantidade := range repetidas {
		if quantidade > 1 {
			for i := 0; i < quantidade-1; i++ {
				saida += fmt.Sprintf("%d ", numero)
			}
		}
	}

	if saida == "" {
		return "N"
	}

	return saida[:len(saida)-1]
}

func numeroFaltando(figs []int, total int) string {
	controle := make(map[int]bool)
	saida := ""

	for _, valor := range figs {
		controle[valor] = true
	}

	for i := 1; i <= total; i++ {
		if !controle[i] {
			saida += fmt.Sprintf( "%d ", i)
		}
	}
	if saida == ""{
		return "N"
	}
	return saida[:len(saida)-1]
}

func main () {
	var total_figurinhas int
	var qtd int

	fmt.Scan(&total_figurinhas,&qtd)
	figs := make([]int, qtd)


	for i:=0; i < qtd; i++ {
		fmt.Scan(&figs[i])
	}

	fmt.Println(figurinhasRepetidas(figs))
	fmt.Println(numeroFaltando(figs,total_figurinhas))
}
