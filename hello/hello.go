package main
import (
   "fmt"
   "github.com/pmthompson/strpkg"
)
func main() {
   fmt.Println("Hello, new gopher world -- " + strpkg.Reverse("Hello, new gopher world"))
   // fmt.Println("Hello, new gopher world -- ")
   // fmt.Println( strpkg.Reverse("Hello, new gopher world") )
}
