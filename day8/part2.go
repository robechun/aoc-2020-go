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

        v, _ := strconv.ParseFloat(a[1], 10)

        if a[0] == "acc" {
            accum += int(v)
            inst[i] = true
            i++
        } else if a[0] == "jmp" {
            tmp := arr[i]
            arr[i] = "nop +0"
            nm := make(map[int]bool)
            na := accum
            for k,v := range inst {
                nm[k] = v
            }
            if a, b := traverse(i, na, arr, nm); b {
                return a
            }
            arr[i] = tmp

            inst[i] = true
            i = i + int(v)
        } else {
            tmp := arr[i]
            arr[i] = "jmp " + string(int(v))
            nm := make(map[int]bool)
            na := accum
            for k,v := range inst {
                nm[k] = v
            }
            if a, b := traverse(i, na, arr, nm); b {
                return a
            }
            arr[i] = tmp

            inst[i] = true
            i++
        }
    }


	return accum
}

func traverse(i int, accum int, arr []string, inst map[int]bool) (int, bool) {

    for {
        if arr[i] == "" {
            return accum, true
        }
        a := strings.Split(arr[i], " ")
        if inst[i] == true {
            return 0, false
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

    return 0, false
}


func getFile(name string) (*os.File, error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, err
	}

	return file, nil
}
