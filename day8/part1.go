package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
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

    arr  := make([]string, 10000)
    inst := make(map[int]bool)

    i := 0
    accum := 0

	for scanner.Scan() {
		line := scanner.Text()
        arr[i] = line
        i++
	}

    i = 0
    for {
        a := strings.Split(arr[i], " ")
        if inst[i] == true {
            return accum
        }

        v, _ := strconv.ParseFloat(a[1], 10)
        inst[i] = true

        if a[0] == "acc" {
            accum += int(v)
            i++
        } else if a[0] == "jmp" {
            i = i + int(v)
        } else {
            i++
        }
    }


	return accum
}


func getFile(name string) (*os.File, error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, err
	}

	return file, nil
}
