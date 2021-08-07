package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	s := strings.Split(scanner.Text(), ",")

	codes := make([]int, 0)
	for _, v := range s {
		vi, _ := strconv.Atoi(v)
		codes = append(codes, vi)
	}

	//partOne(codes)
	partTwo(codes)
}

func partOne(codes []int) {
	op, idx := codes[0], 0
	for op != 99 {

		switch op {
		case 1: sum(idx, &codes)
		case 2: multiply(idx, &codes)
		case 99: break
		default:
			panic("Oops, something went terribly wrong!")
		}

		idx += 4
		op = codes[idx]
	}

	fmt.Println(codes)
}

func partTwo(codes []int) {
	noun, found := 0, false
	for noun <= 99 {
		for verb := 0; verb < 100; verb++ {
			codesCopy := make([]int, len(codes))
			copy(codesCopy, codes)
			ans := try(noun, verb, codesCopy)
			if ans == 19690720 {
				fmt.Println(fmt.Sprintf("%d, %d", noun, verb))
				found = true
				break
			}
		}

		if found {
			break
		}

		noun++
	}
}

func try(noun int, verb int, codes []int) int {
	codes[1] = noun
	codes[2] = verb

	op, idx := codes[0], 0
	for op != 99 {

		switch op {
		case 1: sum(idx, &codes)
		case 2: multiply(idx, &codes)
		case 99: break
		default:
			panic("Oops, something went terribly wrong!")
		}

		idx += 4
		op = codes[idx]
	}

	return codes[0]
}

func sum(currPos int, codes *[]int) {
	pos1 := (*codes)[currPos+1]
	pos2 := (*codes)[currPos+2]
	pos3 := (*codes)[currPos+3]

	(*codes)[pos3] = (*codes)[pos1] + (*codes)[pos2]
}

func multiply(currPos int, codes *[]int) {
	pos1 := (*codes)[currPos+1]
	pos2 := (*codes)[currPos+2]
	pos3 := (*codes)[currPos+3]

	(*codes)[pos3] = (*codes)[pos1] * (*codes)[pos2]
}
