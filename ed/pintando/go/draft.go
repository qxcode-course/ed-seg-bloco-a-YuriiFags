package main
import(
    "fmt"
    "math"
) 


func main(){
    var A,B,C float64
    fmt.Scan(&A,&B,&C)

    P := (A+B+C)/2

    area := math.Sqrt(P*(P-A)*(P-B)*(P-C))
    fmt.Printf("%.2f\n",area)
    
}