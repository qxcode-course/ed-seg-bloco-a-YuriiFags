package main

import "fmt"

func main() {

	fila := NewQueue[string]()

	for i := 'A'; i <= 'P'; i++{
		fila.Enqueue(string(i))
		//Aqui os times estão enfileirados
	}

	for i := 0; i < 15; i++ {
		var gols1, gols2 int
		fmt.Scan(&gols1,&gols2)

		time1 := fila.Dequeue()
		time2 := fila.Dequeue()

		if gols1 > gols2 {
			fila.Enqueue(time1)
		} else {
			fila.Enqueue(time2)
		}
	}

	fmt.Println(fila.Dequeue())
	//o time vencedor
}
