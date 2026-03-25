package main
import "fmt"

func marcar_quem_saiu(sairam[]int) map[int]bool{
	registro_saidas := make(map[int]bool)

	for _, id := range sairam {
		registro_saidas[id] =  true
	}
	return registro_saidas
}

func imprimir(fila[]int, saiu map[int]bool){
	for _, pessoa := range fila {
		if !saiu[pessoa]{
			fmt.Printf("%v ",pessoa)
		}
	}
	fmt.Println()
}

func main(){
	var qtd_pessoas, identificador int

	fmt.Scan(&qtd_pessoas)

	fila := make([]int, qtd_pessoas)
	for i := 0; i < qtd_pessoas; i++ {
		fmt.Scan(&fila[i])
	}

	fmt.Scan(&identificador)

	sairam := make([]int,identificador)

	for i := 0; i < identificador; i++ {
		fmt.Scan(&sairam[i])
	}

	registro_saidas := marcar_quem_saiu(sairam)
	imprimir(fila,registro_saidas)

}



