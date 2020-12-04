package main

import (
    "fmt"
    "os"
    "bufio"
    "strconv"
)

func main() {

    const desiredSum = 2020

    file, err := os.Open("input.txt")
    defer file.Close()
    if err != nil {
        fmt.Println("Encountered an error: ", err)
        return
    }

    input := make([]int, 0)

    // populate input
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        cur, err := strconv.Atoi(scanner.Text())
        if err != nil {
            fmt.Println("Encountered an error: ", err)
            return
        }
        input = append(input, cur)
    }


    // cross join
    encountered := make(map[int]int)

    for i := 0; i < len(input); i++ {
        for j := i+1; j < len(input); j++ {
            encountered[input[i] + input[j]] = input[i] * input[j]
        }
    }

    // evaluate
    for _, val := range input {
        if enc, ok := encountered[desiredSum - val]; ok {
            fmt.Printf("Answer is %v\n", val * enc)
            return
        }
    }


    fmt.Printf("Didn't find combination of numbers to %d\n", desiredSum)
    return
}
