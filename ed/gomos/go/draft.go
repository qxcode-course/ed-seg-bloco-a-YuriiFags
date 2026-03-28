package main
import "fmt"

func main(){
  var Q int
  var D string
  fmt.Scan(&Q,&D)

  type Gomos struct{
     x, y int
  }

  cobra := make([]Gomos, Q)

  for i := range cobra {
    fmt.Scan(&cobra[i].x, &cobra[i].y)
  }
  
  for i := Q - 1; i >= 0; i-- {
    cobra[i] = cobra[i - 1]
  }

  s

}