package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
)

func main() {
	file, err := getFile("input.txt")
	if err != nil {
		fmt.Println(err)
		panic("file needs to be present!")
	}

	valid := parse(file)

	fmt.Printf("Valid passports: %d\n", valid)
}

func parse(file *os.File) int {
	sum := 0

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		properties := make(map[string]string)

		if line != "" {
			fillProperties(properties, strings.Split(line, " "))
			blah(scanner, properties)
		}

		if validPassport(properties) {
			sum++
		}
	}

	return sum
}

func blah(scanner *bufio.Scanner, properties map[string]string) {
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			fillProperties(properties, strings.Split(line, " "))
		} else {
			break
		}
	}
}

func validPassport(properties map[string]string) bool {
	var requiredKeys = [7]string{ "byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid" }

	for _, key := range requiredKeys {
		if _, ok := properties[key]; !ok {
			return false
		}
	}

	return true
}

func fillProperties(mappedProperties map[string]string, properties []string) {
	for _, val := range properties {
		kv := strings.Split(val, ":")
		mappedProperties[kv[0]] = kv[1]
	}

}

func getFile(name string) (*os.File, error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, err
	}

	return file, nil
}
