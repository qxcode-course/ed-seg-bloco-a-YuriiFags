package main
import "fmt"

func main(){
    var nome string
    var idade int
    var classificacao string

    fmt.Scan(&nome,&idade)

    if idade < 12 {
        classificacao = "crianca"
    } else if idade < 18 {
        classificacao = "jovem"
    } else if idade < 65 {
        classificacao = "adulto"
    } else if idade < 1000 {
        classificacao = "idoso"
    } else {
        classificacao = "mumia"
    }

    fmt.Println(nome,"eh",classificacao)
}