package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main(){
  // input, err := os.ReadFile("./dayone.txt")
  // if (err != nil){
  //   panic(err)

  // }
  filename := "./dayone.txt"
  f, err := os.Open(filename)
  if (err != nil){
    panic(err)
  }

  fileinfo, err := os.Stat(filename)
  b1 := make([]byte, fileinfo.Size())
  n1, err := f.Read(b1)
  fileString := string(b1[:n1])

  var elves []int
  elves = make([]int, 500)
  var currentMeal []byte
  currentElf := 0
  i := 0
  var maxCals float64 = 0.0

  for i < len(fileString){
    if fileString[i] == 10{
      meal, _ := strconv.Atoi(string(currentMeal))
      elves[currentElf] += meal
      currentMeal = []byte{}
      if i + 1 < len(fileString) && fileString[i + 1] == 10{
        maxCals = math.Max(float64(maxCals), float64(elves[currentElf]))
        currentElf++
      }
      i += 1
      continue
    }
    currentMeal = append(currentMeal, fileString[i])
    i += 1
  }
  fmt.Print(maxCals)
}

