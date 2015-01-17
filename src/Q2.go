package main
import  "fmt"

func main(){

  fmt.Printf("For loop:\n")
  for i := 0; i < 10; i++ {
    fmt.Printf("%d\n", i)
  }


  fmt.Printf("Goto loop:\n")
  i := 0
  Here: fmt.Printf("%d\n", i)

  i++

  if (i < 10) {
    goto Here
  }

  fmt.Printf("For loop with array:\n")

  var array [10]int

  for i := 0; i < 10; i++ {
    array[i] = i
  }

  fmt.Printf("%d\n", array)

}