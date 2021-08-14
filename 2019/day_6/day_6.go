package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	//partOne()
	partTwo()
}

func partOne() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	tracker := make(map[string][]string)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), ")")
		tracker[s[0]] = append(tracker[s[0]], s[1])
	}

	fmt.Println(count("COM", tracker, 0))
}

func partTwo() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	tracker := make(map[string][]string)
	trackerReversed := make(map[string][]string)
	var start string
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), ")")
		tracker[s[0]] = append(tracker[s[0]], s[1])

		trackerReversed[s[1]] = append(trackerReversed[s[1]], s[0])
		trackerReversed[s[0]] = append(trackerReversed[s[0]], s[1])
		if s[1] == "YOU" {
			start = s[0]
		}
	}

	find(start, map[string]bool{"YOU": true}, 0, trackerReversed, tracker)
}

func find(curr string, prev map[string]bool, count int, trackerReversed map[string][]string, tracker map[string][]string) bool {
	for _, v := range tracker[curr] {
		if v == "SAN" {
			fmt.Println(count)
			return true
		}
	}

	for _, v := range trackerReversed[curr] {
		if _, ok := prev[v]; ok {
			continue
		}
		prev[v] = true
		if find(v, prev, count + 1, trackerReversed, tracker) {
			return true
		}
	}

	return false
}

func count(curr string, tracker map[string][]string, depth int) int {
	sum := 0
	for _, v := range tracker[curr] {
		sum += count(v, tracker, depth + 1)
	}

	return depth + sum
}
