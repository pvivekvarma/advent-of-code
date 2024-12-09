
package main

import (
  "fmt"
  "os"
  "bufio"
  "strings"
)

func main() {
  fmt.Println("Day 4: Ceres Search")

  inputs, err := os.Open("./input.txt")
  if err != nil {
    panic(err)
  }

  defer inputs.Close()

  fileScanner := bufio.NewScanner(inputs)
  fileScanner.Split(bufio.ScanLines)

  count := 0

  mat := make([][]string, 140)

  for fileScanner.Scan() {
    line := strings.Split(fileScanner.Text(), "")
    mat[count] = line
    count++
  }

  part1(mat)

  part2(mat)

}

func part1(mat [][]string) {
  total := 0
  for i := 0; i < 140; i++ {
    for j := 0; j < 140; j ++ {
      if mat[i][j] == "X" {
        if j >= 3 {
          // SAMX
          if join(mat[i][j-3], mat[i][j-2], mat[i][j-1]) == "SAM" {
            total++
          }
          if i >= 3 {
            // S
            //  A
            //   M
            //    X
            if join(mat[i-3][j-3], mat[i-2][j-2], mat[i-1][j-1]) == "SAM" {
              total++
            }
          }
          if 139 - i >= 3 {

            //    X
            //   M
            //  A
            // S
            if join(mat[i+1][j-1], mat[i+2][j-2], mat[i+3][j-3]) == "MAS" {
              total++
            }

          }
        }
        if 139 - j >= 3 {
          // XMAS
          if join(mat[i][j+1], mat[i][j+2], mat[i][j+3]) == "MAS" {
            total++
          }
        }
        if i >= 3 {
          // S
          // A
          // M
          // X
          if join(mat[i-3][j], mat[i-2][j], mat[i-1][j]) == "SAM" {
            total++
          }
          if 139 - j >= 3 {
            //    S
            //   A
            //  M
            // X
            if join(mat[i-3][j+3], mat[i-2][j+2], mat[i-1][j+1]) == "SAM" {
              total++
            }
            
          }
        }
        if 139 - i >= 3 {
          // X
          // M
          // A
          // S
          if join(mat[i+1][j], mat[i+2][j], mat[i+3][j]) == "MAS" {
            total++
          }
          if 139 - j >= 3 {
            // X
            //  M
            //   A
            //    S
            if join(mat[i+1][j+1], mat[i+2][j+2], mat[i+3][j+3]) == "MAS" {
              total++
            }
          }
        }
      }
    }
  }
  fmt.Printf("Result: %d\n", total)
}

func part2(mat [][]string) {
  total := 0
  for i := 1; i < 139; i++ {
    for j := 1; j < 139; j++ {
      if mat[i][j] == "A" {
        // M  M
        //  A
        // S  S
        if mat[i-1][j-1] == "M" && mat[i-1][j+1] == "M" && mat[i+1][j-1] == "S" && mat[i+1][j+1] == "S" {
          total ++
        }

        // M  S
        //  A
        // M  S
        if mat[i-1][j-1] == "M" && mat[i-1][j+1] == "S" && mat[i+1][j-1] == "M" && mat[i+1][j+1] == "S" {
          total ++
        }

        // S  S
        //  A
        // M  M
        if mat[i-1][j-1] == "S" && mat[i-1][j+1] == "S" && mat[i+1][j-1] == "M" && mat[i+1][j+1] == "M" {
          total ++
        }

        // S  M
        //  A
        // S  M
        if mat[i-1][j-1] == "S" && mat[i-1][j+1] == "M" && mat[i+1][j-1] == "S" && mat[i+1][j+1] == "M" {
          total ++
        }
      }
    }
  }
  fmt.Printf("Result: %d\n", total)
}

func join(a, b, c string) string{
  return a + "" + b + "" + c
}
