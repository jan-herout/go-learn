package main

import "fmt"

func main() {    
    x, y := 3.0, 2.0               // float64 nebo float32, podle platformy
    div := x/y                     
    fmt.Println(div, x/y == 1.5)   // ==> CO TO VRÁTÍ?
    
    var a, b float32               // explicitní float32
    a, b = 1, 0.99                
    dif := a - b                    
    fmt.Println(dif, dif == 0.01)  // ==> CO TO VRÁTÍ?
}
