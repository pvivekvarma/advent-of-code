
package main

import (
  "fmt"
  "os"
  "strconv"
  "bufio"
  "regexp"
)

func main() {
  fmt.Println("Day 3: Mull It Over")

  inputs, err := os.Open("./input.txt")
  if err != nil {
    panic(err)
  }

  defer inputs.Close()

  mulr, _ := regexp.Compile("mul\\((\\d{1,3}),(\\d{1,3})\\)")


  fileScanner := bufio.NewScanner(inputs)
  fileScanner.Split(bufio.ScanLines)

  total := 0
  count := 0

  for fileScanner.Scan() {
    allMatches := mulr.FindAllString(fileScanner.Text(), -1)
    
    for _, match := range(allMatches) {
      subMatch := mulr.FindStringSubmatch(match)
      count++
      i, err := strconv.Atoi(subMatch[1])
      if err != nil {
        panic(fmt.Sprintf("Failed to convert %d to int", subMatch[1]))
      }

      j, err := strconv.Atoi(subMatch[2])
      if err != nil {
        panic(fmt.Sprintf("Failed to convert %d to int", subMatch[2]))
      }

      total += i*j
    }
  }

  fmt.Printf("Result: %d\n", total)
  fmt.Printf("Count: %d\n", count)

  inputs2, err := os.Open("./input.txt")
  if err != nil {
    panic(err)
  }

  defer inputs2.Close()

  fileScanner = bufio.NewScanner(inputs2)
  fileScanner.Split(bufio.ScanLines)


  //Using lookarounds (Not supported by Go) - mulrcon, _ := regexp.Compile("(?<=don't\\(\\))([\\S\\s]*?)(?=(?:do\\(\\))|(?:\\z))")
  
  //Match between don't() and do() - Not using lookarounds
  //mulrcon, _ := regexp.Compile("(?:don't\\(\\))([\\S\\s]*?)(?:(?:do\\(\\))|(?:\\z))")

  //Match between don't() and do() - Not using lookarounds
  mulrcon, _ := regexp.Compile("(?:(?:\\A)|(?:do\\(\\)))([\\S\\s]*?)(?:(?:don't\\(\\))|(?:\\z))")
  
  //mulrcon, _ := regexp.Compile("(?:\\A)([\\S\\s]*)(?:\\z)")

  remove := 0
  count = 0

  for fileScanner.Scan() {
    text := fileScanner.Text()
    conditionalMatches := mulrcon.FindAllString(text, -1)

    for _, cMatch := range(conditionalMatches) {
      allMatches := mulr.FindAllString(cMatch, -1)

      for _, match := range(allMatches) {

        //fmt.Println(match)

        subMatch := mulr.FindStringSubmatch(match)
        count++
        i, err := strconv.Atoi(subMatch[1])
        if err != nil {
          panic(fmt.Sprintf("Failed to convert %d to int", subMatch[1]))
        }

        j, err := strconv.Atoi(subMatch[2])
        if err != nil {
          panic(fmt.Sprintf("Failed to convert %d to int", subMatch[2]))
        }

        remove += i*j
      }
    }
  }

  fmt.Printf("Result: %d\n", remove)
  fmt.Printf("Count: %d\n", count)
}
