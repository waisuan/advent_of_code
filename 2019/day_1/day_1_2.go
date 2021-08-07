package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		mass, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		sum += fuel(mass)
	}

	fmt.Println(sum)
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

func fuel(mass int) int {
	fuels := 0
	for mass > 0 {
		mass = mass / 3 - 2
		if mass > 0 {
			fuels += mass
		}
	}
	return fuels
}


