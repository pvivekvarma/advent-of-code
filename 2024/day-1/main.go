package main

import (
  "fmt"
  "os"
  "strings"
  "strconv"
  "bufio"
  "sort"
)

func main() {
	fmt.Println("Day 1: Historian Hysteria")

  inputs, err := os.Open("./input.txt")
  if err != nil {
    panic(err)
  }

  defer inputs.Close()

  /** Todo
  1. Fetch list 1 and 2 from the file
  2. Sort list 1 and 2
  3. Calculate the distances between each pair / or just calculate the total distance 
  **/

  fileScanner := bufio.NewScanner(inputs)
  fileScanner.Split(bufio.ScanLines)

  var list1, list2 []int

  for fileScanner.Scan() {
    line := fileScanner.Text()
    items := strings.Split(line, "   ")
    i, err := strconv.Atoi(items[0])
    if err != nil {
      panic(fmt.Sprintf("Failed to parse string %s to integer", items[0]))
    }
    list1 = append(list1, i) 

    i, err = strconv.Atoi(items[1])
    if err != nil {
      panic(fmt.Sprintf("Failed to parse string %s to integer", items[1]))
    }
    list2 = append(list2, i)
  }

  if len(list1) != len(list2) {
    panic("List length is not the same")
  }

  sort.Ints(list1)
  sort.Ints(list2)

  part1(list1, list2)
  part2(list1, list2)

}

func part1(list1, list2 []int) {
  var diff int
  for i := range(list1) {
    if list1[i] < list2[i] {
      diff += list2[i] - list1[i]
    } else {
      diff += list1[i] - list2[i]
    }
  }

  fmt.Printf("Difference: %d\n", diff)
}


func part2(list1, list2 []int) {
  var sim int
  var prev int
  var times int
  for i := range(list1) {
    x := list1[i]

    if x == prev {
      sim += x*times
    } else {
      times = findTimes(list2, x)
      sim += x*times
    }
    prev = x
  }

  fmt.Printf("Similarity: %d\n", sim)

}


func findTimes(list []int, x int) int {
  var times int
  for i := range(list) {
    if list[i] > x {
      break;
    }
    if list[i] == x {
      times += 1
    }
  }

  return times
}

