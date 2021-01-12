package main 

import ( 
    "flag"
    "fmt"
) 

const VERSION = "1.0"


func main() { 
    
    // Define a bool flag 
    flagPresence:= flag.Bool("version", false, "affiche la version du programme")

  
    // Parse command line  
    // into the defined flags 
    flag.Parse() 
  
    if *flagPresence {
        fmt.Println(VERSION) 
    }else{
        fmt.Println("pas de flag version")
    }
}
