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
		if v, ok := properties[key]; !ok {
			return false
		} else {
			switch v {
			case "byr":
				byr, err := strconv.Atoi(v)
				if err != nil {
					return false
				}
				if byr < 1920 || byr > 2002 {
					return false
				}
			case "iyr":
				iyr, err := strconv.Atoi(v)
				if err != nil {
					return false
				}
				if iyr < 2010 || iyr > 2020 {
					return false
				}
			case "eyr":
				eyr, err := strconv.Atoi(v)
				if err != nil {
					return false
				}
				if eyr < 2020 || eyr > 2030 {
					return false
				}
			case "hgt":
				// todo regex
			case "hcl":
				// todo regex
			case "ecl":
				if v != "amb" || v != "blu" || v != "brn" || v != "gry" || v != "grn" || v != "hzl" || v != "oth" {
					return false
				}
			case "pid":
				if len(v) != 9 {
					return false
				}
			default:
				return true
			}

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