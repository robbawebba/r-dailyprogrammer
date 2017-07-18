package main
 import (
   "fmt"
   "os"
 )

 func main() {
   fmt.Println("Hello from main")
   args := os.Args[1:]
   if len(args) == 0 {
     fmt.Println("too few arguments")
   }
 }
