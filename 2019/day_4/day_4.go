package main

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	start = 109165
	end = 576723
)

func main() {
	//partOne()
	partTwo()
}

func partOne() {
	total := 0
	for i := start; i <= end; i++ {
		if hasAdjacentDigits(i) && increasingDigits(i) {
			total++
		}
	}
	fmt.Println(total)
}

func partTwo() {
	total := 0
	for i := start; i <= end; i++ {
		if hasSpecialAdjacentDigits(i) && increasingDigits(i) {
			total++
		}
	}
	fmt.Println(total)
}

func hasAdjacentDigits(num int) bool {
	tmp := strconv.Itoa(num)
	s := strings.Split(tmp, "")
	for i := 1; i < len(s); i++ {
		if s[i-1] == s[i] {
			return true
		}
	}

	return false
}

func increasingDigits(num int) bool {
	tmp := strconv.Itoa(num)
	s := strings.Split(tmp, "")
	for i := 1; i < len(s); i++ {
		if s[i-1] > s[i] {
			return false
		}
	}

	return true
}

func hasSpecialAdjacentDigits(num int) bool {
	tmp := strconv.Itoa(num)
	s := strings.Split(tmp, "")
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			count := 0
			j := i
			for ; j < len(s); j++ {
				if s[j] == s[i] {
					count++
				} else {
					break
				}
			}
			if count == 2 {
				return true
			}
			i = j-1
		}
	}

	return false
}