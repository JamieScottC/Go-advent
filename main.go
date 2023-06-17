package main

import (
  "fmt" 
  "os"
)

func main(){
  // input, err := os.ReadFile("./dayone.txt")
  // if (err != nil){
  //   panic(err)

  // }
  f, err := os.Open("./dayone.txt")
  if (err != nil){
    panic(err)
  }
  b1 := make([]byte, 6)
  n1, err := f.Read(b1)
  fmt.Print(string(b1[:n1]))
  
}

