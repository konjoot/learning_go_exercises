package main

import "fmt"

// Average
func main() {
  var sum float64
  array := [...]float64{ 1.0, 2.1, 3.2, 4.3, 5.4 }
  slice := array[:]

  switch len(slice){
    case 0:
      fmt.Printf("%f\n", 0.0)
      return
    default:
      for _, numb := range slice { sum += numb }
  }

  fmt.Printf("%f\n", sum/float64(len(slice)))
}
