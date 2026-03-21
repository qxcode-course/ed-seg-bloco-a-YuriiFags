package main
import "fmt"
func main() {
   var Q int
   var D string 
   fmt.Scan(&Q,&D)

   type gomo struct{
        x,y int
   }
   
   cobra := make([]gomo,Q)

   for i := range cobra {
    fmt.Scan(&cobra[i].x, &cobra[i].y)
   }

   for i := Q - 1; i > 0; i-- {
    cobra[i] = cobra[i-1]
   }

  switch D {
  case "L":
    cobra[0].x--
  case "R":
    cobra[0].x++
  case "U":
    cobra[0].y--
  case "D":
    cobra[0].y++
  }

  for _, gomo := range cobra {
     fmt.Printf("%v %v\n",gomo.x,gomo.y)
  }
 
}
