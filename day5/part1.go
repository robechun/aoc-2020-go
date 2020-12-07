package main

import (
    "os"
    "bufio"
    "fmt"
    "math"
)

func main() {
    file, err := getFile("input.txt")
    if err != nil {
        panic("input file is required")
    }

    maxID := parse(file)
    fmt.Printf("MaxID: %v\n", maxID)
}


type seat struct {
    row int
    column int
}

func parse(file *os.File) int {
    scanner := bufio.NewScanner(file)


    maxID := 0
    for scanner.Scan() {
        line := scanner.Text()
        seat := getSeat(line)

        maxID = max(maxID, seat.row * 8 + seat.column)
    }

    return maxID

}


func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}

func getSeat(line string) *seat {
    lowerRow := 0.0
    upperRow := 127.0

    // get row
    for i := 0; i < 7; i++ {
        mid := lowerRow + ((upperRow - lowerRow) / 2)
        if string(line[i]) == "F" {
            upperRow = math.Floor(mid)
        } else {
            lowerRow = math.Ceil(mid)
        }

    }

    lowerCol := 0.0
    upperCol := 7.0

    // get col
    for i := 7; i < 10; i++ {
        mid := lowerCol + ((upperCol - lowerCol) / 2)
        if string(line[i]) == "L" {
            upperCol = math.Floor(mid)
        } else {
            lowerCol = math.Ceil(mid)
        }
    }

    return &seat{
        row: int(lowerRow),
        column: int(lowerCol),
    }

}


func getFile(name string) (*os.File, error) {
    file, err := os.Open(name)
    if err != nil {
        fmt.Println("something went wrong opening the file")
        return nil, err
    }

    return file, nil
}

