package main 
  
import "fmt"
  
func main() { 
  
    // Creating a map 
    // Using make() function 
    var My_map = make(map[float64]string) 
    fmt.Println(My_map) 
  
    // As we already know that make() function 
    // always returns a map which is initialized 
    // So, we can add values in it 
    My_map[1.3] = "shiva"
    My_map[1.5] = "bharath"
    fmt.Println(My_map) 
} 