package main

import (
  "fmt"
  "os"
  "strings"
  "strconv"
  "bufio"
)

func main() {
	fmt.Println("Day 2: Red-Nosed Reports")


  /** Todo
  1. Read inputs file.
  2. For each like in input file, check if the report is valid or not and increment the totoal value
  **/

  part1()
  part2()
}

func part1() {
  inputs, err := os.Open("./input.txt")
  if err != nil {
    panic(err)
  }

  defer inputs.Close()

  fileScanner := bufio.NewScanner(inputs)
  fileScanner.Split(bufio.ScanLines)

  var totalSafe int
  for fileScanner.Scan() {
    report := strings.Split(fileScanner.Text(), " ")
    
    safe := true
    inc := true
    dec := true

    for i := 0; i < len(report) - 1; i++ {
      j, err := strconv.Atoi(report[i])
      if err != nil {
        panic(fmt.Sprintf("Failed to parse %d to int", report[i]))
      }

      k, err := strconv.Atoi(report[i+1])
      if err != nil {
        panic(fmt.Sprintf("Failed to parse %d to int", report[i+1]))
      }

      if j > k && dec {
        inc = false
        if j - k > 3 {
          safe = false
        }
      } else if k > j && inc {
        dec = false
        if k - j > 3 {
          safe = false
        }
      } else {
        safe = false
      }
    }

    if safe {
      totalSafe++;
    }
  }
  

  fmt.Printf("Total number of safe reports: %d\n", totalSafe)
}





func part2() {
  inputs, err := os.Open("./input.txt")
  if err != nil {
    panic(err)
  }

  defer inputs.Close()

  fileScanner := bufio.NewScanner(inputs)
  fileScanner.Split(bufio.ScanLines)

  var totalSafe int

  for fileScanner.Scan() {
    report := strings.Split(fileScanner.Text(), " ")
    report = report[:]
    
    for index := -1; index < len(report); index ++ {

      safe := true
      inc := true
      dec := true

      modifiedReport := make([]string, 0)
      modifiedReport = append(modifiedReport, report[:]...)

      if index >= 0 {
        if index == len(report) - 1 {
          modifiedReport = append(modifiedReport[:index])
        } else {
          modifiedReport = append(modifiedReport[:index], report[index+1:]...)
        }
      }

      for i := 0; i < len(modifiedReport) - 1; i++ {
        j, err := strconv.Atoi(modifiedReport[i])
        if err != nil {
          panic(fmt.Sprintf("Failed to parse %d to int", modifiedReport[i]))
        }

        k, err := strconv.Atoi(modifiedReport[i+1])
        if err != nil {
          panic(fmt.Sprintf("Failed to parse %d to int", modifiedReport[i+1]))
        }

        if j > k && dec {
          inc = false
          if j - k > 3 {
            safe = false
          }
        } else if k > j && inc {
          dec = false
          if k - j > 3 {
            safe = false
          }
        } else {
          safe = false
        }
      }

      if safe {
        totalSafe++;
        break;
      }
    }
  }

  fmt.Printf("Total number of safe reports: %d\n", totalSafe)
}
