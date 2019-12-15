
package main

import (
  "flag"
  "fmt"
)

func getTheFlag() *string {
  return flag.String("name", "everyone","greeting object")
}

func main() {

  var i16 int16
  i16 = -255
  var i8 int8
  i8 = int8(i16)
  fmt.Printf("i8 value:  %v!\n", i8)
  fmt.Printf("i16 value:  %v!\n", i16)
  fmt.Printf("error: %v!\n", string(-221))
  var aaa rune
  aaa = 23321
  fmt.Println("aaa value", aaa)

}