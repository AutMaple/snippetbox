package main

import "fmt"

func main() {
  arr := []int{1,2,3,4,5}
  for _, value := range arr {
    fmt.Println(value)
  }
  test(arr...)
}

func test(arr ...int) {
  for _, value := range arr {
    fmt.Println(value)
  }  
}
