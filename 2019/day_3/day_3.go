package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x, y int
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	wireOne := strings.Split(scanner.Text(), ",")

	scanner.Scan()
	wireTwo := strings.Split(scanner.Text(), ",")

	//partOne(wireOne, wireTwo)
	partTwo(wireOne, wireTwo)
}

func partOne(wireOne []string, wireTwo []string) {
	path := make(map[Point]int)

	oriX, oriY := 0, 0
	for move := 0; move < len(wireOne); move++ {
		dirOne := string(wireOne[move][0])
		dirAmnt, _ := strconv.Atoi(wireOne[move][1:])
		for ; dirAmnt > 0; dirAmnt-- {
			destination(dirOne, &oriX, &oriY)
			path[Point{x: oriX, y: oriY}] = 1
		}
	}

	minDistance := math.MaxInt64
	oriX, oriY = 0, 0
	for move := 0; move < len(wireTwo); move++ {
		dirTwo := string(wireTwo[move][0])
		dirAmnt, _ := strconv.Atoi(wireTwo[move][1:])
		for ; dirAmnt > 0; dirAmnt-- {
			destination(dirTwo, &oriX, &oriY)
			if _, ok := path[Point{x: oriX, y: oriY}]; ok {
				dist := manhattanDistance(0, 0, oriX, oriY)
				if dist < minDistance {
					minDistance = dist
				}
			}
		}
	}

	fmt.Println(minDistance)
}

func partTwo(wireOne []string, wireTwo []string) {
	path := make(map[Point]int)

	oriX, oriY, steps := 0, 0, 0
	for move := 0; move < len(wireOne); move++ {
		dirOne := string(wireOne[move][0])
		dirAmnt, _ := strconv.Atoi(wireOne[move][1:])
		for ; dirAmnt > 0; dirAmnt-- {
			destination(dirOne, &oriX, &oriY)
			steps++
			path[Point{x: oriX, y: oriY}] = steps
		}
	}

	minSum := math.MaxInt64
	oriX, oriY, steps = 0, 0, 0
	for move := 0; move < len(wireTwo); move++ {
		dirTwo := string(wireTwo[move][0])
		dirAmnt, _ := strconv.Atoi(wireTwo[move][1:])
		for ; dirAmnt > 0; dirAmnt-- {
			destination(dirTwo, &oriX, &oriY)
			steps++
			if val, ok := path[Point{x: oriX, y: oriY}]; ok {
				if val + steps < minSum {
					minSum = val + steps
				}
			}
		}
	}

	fmt.Println(minSum)
}

func destination(dir string, x *int, y *int) {
	if dir == "R" {
		*x += 1
	} else if dir == "L" {
		*x -= 1
	} else if dir == "U" {
		*y += 1
	} else if dir == "D" {
		*y -= 1
	}
}

func manhattanDistance(x1 int, y1 int, x2 int, y2 int) int {
	return abs(x1-x2) + abs(y1-y2)
}

func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}