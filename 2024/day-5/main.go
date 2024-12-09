
package main

import (
  "fmt"
  "os"
  "bufio"
  "strings"
  "strconv"
  "slices"
)

func main() {
  fmt.Println("Day 5: Print Queue")

  inputs, err := os.Open("./input.txt")
  //inputs, err := os.Open("./sample-input.txt")
  if err != nil {
    panic(err)
  }

  defer inputs.Close()

  fileScanner := bufio.NewScanner(inputs)
  fileScanner.Split(bufio.ScanLines)

  rules := make([][]string, 100)

  parseRules := true
  total := 0
  fixedTotal := 0

  for fileScanner.Scan() {
    line := fileScanner.Text()
    if parseRules {
      if line == "" {
        parseRules = false
      } else {
        r := strings.Split(line, "|")
        r0, err := strconv.Atoi(r[0])
        if err != nil {
          panic("Error")
        }
        rules[r0] = append(rules[r0], r[1])
      }
    } else {
      updates := strings.Split(line, ",")
      isValidUpdate := checkIfValidUpdate(rules, updates)
      if isValidUpdate {
        mid, err := strconv.Atoi(updates[len(updates)/2])
        if err != nil {
          panic("Failed to parse update")
        }

        total += mid
      } else {
        fixed := makeValid(rules, updates)
        mid, err := strconv.Atoi(fixed[len(fixed)/2])
        if err != nil {
          panic("Failed to parse update")
        }

        fixedTotal += mid
      }
    }
  }

  fmt.Println()
  fmt.Printf("Total: %d\n", total)
  fmt.Printf("Fixed Total: %d\n", fixedTotal)
}

func checkIfValidUpdate(rules [][]string, updates []string) bool {
  isValidUpdate := true
  for i := 0; i < len(updates); i++ {
    ui, err := strconv.Atoi(updates[i])
    if err != nil {
      panic("Failed to parse update")
    }
    if len(rules[ui]) > 0 {
      for j := i + 1; j < len(updates); j++{
        if !slices.Contains(rules[ui], updates[j]) {
          isValidUpdate = false
          break
        }
      }
      for j := i - 1; j >= 0 ; j--{
        if slices.Contains(rules[ui], updates[j]) {
          isValidUpdate = false
          break
        }
      }
    } else {
      continue
    }
  }

  return isValidUpdate
}

func makeValid(rules [][]string, updates []string) []string { 
  for i := len(updates) - 1; i > 0; i-- {
    ui, err := strconv.Atoi(updates[i])
    if err != nil {
      panic("Failed to parse update")
    }
    if len(rules[ui]) > 0 {
      for j := i - 1; j >= 0; j-- {
        if slices.Contains(rules[ui], updates[j]) {
          toMove := updates[j]
          fixed := append(updates[:j], updates[j+1:i+1]...)
          fixed = append(fixed, toMove)
          fixed = append(fixed, updates[i+1:]...)

          updates = fixed
          i = len(updates)
          break
        }
      }
    }
  }

  return updates
}
