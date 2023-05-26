package main

import "fmt"

// * pointer
// & adress

func changeLocal(num *[5]int) { // *
  num[0] = 55

  fmt.Println("\ninside function:\n ", &num)

}
func main() {
  num := [...]int{5, 6, 7, 8, 8}
  fmt.Println("\nbefore passing to function:\n ", &num)

  changeLocal(&num) //num is passed by value
  fmt.Println("\nadress of num:\n  ", &num)

  fmt.Println("\nafter passing to function: ", &num)
}
