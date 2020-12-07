package main

import (
    "os"
    "bufio"
    "fmt"
)

func main() {
    file, err := getFile("input.txt")
    if err != nil {
        panic("input file is required")
    }

    sum := parse(file)

    fmt.Printf("Sum is: %v\n", sum)

}


func parse(file *os.File) int {
    scanner := bufio.NewScanner(file)

    answers := make(map[string]bool)
    curSum := 0
    for scanner.Scan() {
        line := scanner.Text()

        if line == "" {
            curSum += len(answers)
            answers = make(map[string]bool)
            continue
        }

        for _, v := range line {
            answers[string(v)] = true
        }
    }

    return curSum
}





func getFile(name string) (*os.File, error) {
    file, err := os.Open(name)
    if err != nil {
        fmt.Println("something went wrong opening the file")
        return nil, err
    }

    return file, nil
}

