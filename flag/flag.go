package main 

import ( 
    "flag"
    "fmt"
) 
  
func main() { 
    const VERSION = "1.0"
    // Define a bool flag 
    boolVrsPtr := flag.Bool("version", true, VERSION) 
  
    // Parse command line  
    // into the defined flags 
    flag.Parse() 
  
    fmt.Println("Version:", *boolVrsPtr) 
}
