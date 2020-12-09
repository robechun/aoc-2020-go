package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := getFile("input.txt")
	if err != nil {
		fmt.Println(err)
		panic("file needs to be present!")
	}

	valid := parse(file)

	fmt.Printf("answer: %d\n", valid)
}

func parse(file *os.File) int {
	scanner := bufio.NewScanner(file)

	twoSum := make(map[int]bool)

	pre := 0

	for scanner.Scan() {
		line := scanner.Text()
		val, _ := strconv.Atoi(line)

		if pre < 25 {
			twoSum[val] = true
			pre++
		} else {
			if !check(twoSum, val) {
				return val
			}
			twoSum[val] = true
		}
	}

	return -1
}

func check(sums map[int]bool, curVal int) bool {
	for k := range sums {
		if _, ok := sums[curVal-k]; ok {
			return true
		}
	}

	return false
}

func getFile(name string) (*os.File, error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, err
	}

	return file, nil
}
