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

    encountered := make(map[int]bool)

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        cur, err := strconv.Atoi(scanner.Text())
        if err != nil {
            fmt.Println("Encountered an error: ", err)
            return
        }
        if encountered[desiredSum - cur] {
            fmt.Printf("Answer is: %v\n", cur * (desiredSum - cur))
            return
        }

        encountered[cur] = true
    }


    fmt.Printf("Didn't find combination of numbers to %d\n", desiredSum)
    return
}
