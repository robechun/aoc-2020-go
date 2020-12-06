package main

import (
    "fmt"
    "os"
    "bufio"
    "strings"
    "strconv"
    "time"
)

/*
## Problem:

To try to debug the problem, they have created a list (your puzzle input) of passwords (according to the corrupted database) and the corporate policy when that password was set.

For example, suppose you have the following list:

1-3 a: abcde
1-3 b: cdefg
2-9 c: ccccccccc

Each line gives the password policy and then the password. The password policy indicates the lowest and highest number of times a given letter must appear for the password to be valid. For example, 1-3 a means that the password must contain a at least 1 time and at most 3 times.

In the above example, 2 passwords are valid. The middle password, cdefg, is not; it contains no instances of b, but needs at least 1. The first and third passwords are valid: they contain one a or nine c, both within the limits of their respective policies.

How many passwords are valid according to their policies?
*/
func main() {
    file, err := getFile("input.txt")
    defer file.Close()
    if err != nil {
        panic("we need the file")
    }

    start := time.Now()
    valid := walk(file, start)
    duration := time.Since(start)

    fmt.Printf("Valid passwords are: %d\n", valid)
    fmt.Printf("time taken: %v\n", duration.Nanoseconds())

    return
}

func walk(file *os.File, t time.Time) int {
    scanner := bufio.NewScanner(file)

    sum := 0
    for scanner.Scan() {
        line := scanner.Text()
        sum += parse(line)
    }

    fmt.Printf("walk after scan duration %v\n", time.Since(t).Nanoseconds())

    return sum

}

func parse(line string) int {
    parsed := strings.Split(line, " ")
    bounds := strings.Split(parsed[0], "-")
    lower, err := strconv.Atoi(bounds[0])
    if err != nil {
        fmt.Println(err)
    }
    upper, err := strconv.Atoi(bounds[1])
    if err != nil {
        fmt.Println(err)
    }

    lowerMatch := parsed[2][lower - 1] == parsed[1][0]
    upperMatch := parsed[2][upper - 1] == parsed[1][0]

    if lowerMatch != upperMatch {
        return 1
    }

    return 0
}



func getFile(filename string) (*os.File, error) {
    file, err := os.Open(filename)
    if err != nil {
        fmt.Println("Something went wrong", err)
        return nil, err
    }

    return file, nil
}


