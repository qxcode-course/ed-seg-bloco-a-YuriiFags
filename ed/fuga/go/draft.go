package main
import "fmt"

func main(){
    var H,P,F,D int
    fmt.Scan(&H,&P,&F,&D)


   for{
     F = ((F + D) + 16) % 16 
     switch F {
    case P :
        fmt.Println("N")
        return
    case H:
        fmt.Println("S")
        return
    
    }
    //  if F == P {
    //     fmt.Println("N")
    //     return
    // }
    // if F == H {
    //     fmt.Println("S")
    //     return
    // }
   }

   

    
}