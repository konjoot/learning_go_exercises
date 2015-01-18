package main

import "fmt"

func main(){
  a := "A"
  out := ""
  for i := 0; i < 100; i++ {
    out += a
    fmt.Printf("%s\n", out)
  }
}