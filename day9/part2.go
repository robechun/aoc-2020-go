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

	items := make([]int, 100000, 100000)

	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		val, _ := strconv.Atoi(line)
		items[i] = val
		i++
	}

	s := 0
	e := 0
	curSum := 0
	flag := false
	for s <= e {
		if !flag {
			curSum += items[e]
		}
		fmt.Println(curSum, s, e, items[s], items[e])

		if curSum == 373803594 {
			fmt.Println(s, e)
			fmt.Println(items[s] + items[e])
			return thing(s, e, items)
		}

		flag = false
		for curSum > 373803594 {
			curSum -= items[s]
			s++
			flag = true
		}

		if !flag {
			e++
		}
	}

	//     s        e
	// 10 10 5 15 20

	return -1
}

func thing(s, e int, items []int) int {
	low, high := 3738035940, 0
	for i := s; i < e; i++ {
		if items[i] < low {
			low = items[i]
		}
		if items[i] > high {
			high = items[i]
		}
	}

	return low + high
}

func getFile(name string) (*os.File, error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, err
	}

	return file, nil
}
