package main
import "fmt"

func main(){
  for i := 1; i < 101; i++ {
    switch{
      case i % 15 == 0:
        fmt.Printf("FizzBuzz\n")
        continue
      case i % 3 == 0:
        fmt.Printf("Fizz\n")
        continue
      case i % 5 == 0:
        fmt.Printf("Buzz\n")
        continue
    }
    fmt.Printf("%d\n", i)
  }
}