package main

import "fmt"
import "unicode/utf8"

func main(){
  a := "A"
  out := ""
  for i := 0; i < 100; i++ {
    out += a
    fmt.Printf("%s\n", out)
  }

  str := "asSASA ddd dsjkdsjs dk"
  fmt.Printf("%d\n", utf8.RuneCount([]byte(str)))

  strr := []rune(str)
  strr[4] = 'a'
  strr[5] = 'b'
  strr[6] = 'c'
  str = string(strr)
  fmt.Printf("%s\n", str)
}