package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	fmt.Println("Day 6: Guard Gallivant")

	inputs, err := os.Open("./input.txt")
	//inputs, err := os.Open("./sample-input.txt")
	if err != nil {
		panic(err)
	}

	defer inputs.Close()

	fileScanner := bufio.NewScanner(inputs)
	fileScanner.Split(bufio.ScanLines)

	total := 1
	mappedArea := make([][]string, 0)
	positionMap := make([][]int, 0)
	initialPosition := []int{-1, -1}

	area := 0

	for fileScanner.Scan() {
		line := fileScanner.Text()
		mappedPart := strings.Split(line, "")
		mappedArea = append(mappedArea, mappedPart)
		positionMap = append(positionMap, make([]int, len(mappedPart)))
		ij := slices.Index(mappedPart, "^")
		if ij >= 0 {
			initialPosition[0] = area
			initialPosition[1] = ij
		}
		area++
	}

	//right := []int{0,1}
	//left := []int{0, -1}
	//up := []int{-1, 0}
	//down := []int{1, 0}
	out := false
	currentDirection := []int{-1, 0}
	currentPosition := initialPosition

  y := len(mappedArea)
  x := len(mappedArea[0])

	positionMap[initialPosition[0]][initialPosition[1]] = 1
	for !out {
		newPosition := stepForward(currentPosition, currentDirection)
    np0 := newPosition[0]
    np1 := newPosition[1]
		if np0 >= y || np1 >= x || np0 < 0 || np1 < 0 {
			out = true
			currentPosition = newPosition
			//fmt.Printf("Out %v - %v%v\n", newPosition, y, x)
		} else if mappedArea[np0][np1] == "#" {
			currentDirection = turnRight(currentDirection)
			//fmt.Printf("Obstacle ahead, turning right. New direction: %v\n", currentDirection)
		} else {
			currentPosition = newPosition
			if positionMap[currentPosition[0]][currentPosition[1]] != 1 {
				positionMap[currentPosition[0]][currentPosition[1]] = 1
				total++
			}
		}
	}
	fmt.Printf("Total: %d\n", total)

	possibleObstructions := 0

	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			out = false
	    currentPosition = initialPosition
	    currentDirection = []int{-1, 0}
      obstructionsMap := make([][]string, y)
      for k := 0; k < y; k++ {
        obstructionsMap[k] = make([]string, x)
      }
			if positionMap[i][j] != 1 {
				continue
			}

			for !out {
				newPosition := stepForward(currentPosition, currentDirection)
        np0 := newPosition[0]
        np1 := newPosition[1]
        currentDirectionString := fmt.Sprintf("%d%d", currentDirection[0], currentDirection[1])
				if np0 >= y || np1 >= x || np0 < 0 || np1 < 0 {
					out = true
          continue
				} else if (mappedArea[np0][np1] == "#") || (np0 == i && np1 == j) {
          //fmt.Printf("Obstacle ahead!! obstacle position: %v, currentPosition: %v, direction: %v\n", []int{np0, np1}, []int{currentPosition[0], currentPosition[1]}, currentDirection)

          if obstructionsMap[np0][np1] != "" {
            if currentDirectionString == obstructionsMap[np0][np1] {
              possibleObstructions++
						  break
            } else {
              obstructionsMap[np0][np1] = currentDirectionString
            }
          } else {
            obstructionsMap[np0][np1] = currentDirectionString
          }
					currentDirection = turnRight(currentDirection)
        } else { 
          currentPosition = newPosition
        }
			}
		}
	}

	fmt.Printf("Possible Obstructions: %d\n", possibleObstructions)

}

func stepForward(position, direction []int) []int {
	return []int{position[0] + direction[0], position[1] + direction[1]}
}

func turnRight(current []int) []int {
	var newDirection []int
	if current[0] == 0 && current[1] == 1 {
		newDirection = []int{1, 0}
	} else if current[0] == 1 && current[1] == 0 {
		newDirection = []int{0, -1}
	} else if current[0] == 0 && current[1] == -1 {
		newDirection = []int{-1, 0}
	} else if current[0] == -1 && current[1] == 0 {
		newDirection = []int{0, 1}
	}

	return newDirection
}
