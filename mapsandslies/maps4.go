package main 
  
import "fmt"
  
// Main function 
func main() { 
  
    // Creating and initializing a map 
    m := map[int]string{ 
        90: "hello", 
        91: "how", 
        92: "are", 
        93: "you", 
        94: "friend", 
    } 
  
    fmt.Println("Original map: ", m) 
  
    // Adding new key-value pairs in the map 
    m[95] = "shiva"
    m[96] = "bharath"
    fmt.Println("Map after adding new key-value pair:\n", m)
 
    // Updating values of the map 
    m[91] = "akash"
    m[93] = "parthiban"
    fmt.Println("\nMap after updating values of the map:\n", m) 
} 