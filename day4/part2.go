package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
	"regexp"
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
			switch key {
			case "byr":
				byr, err := strconv.Atoi(v)
				if err != nil {
					return false
				}
                if !inBetween(byr, 1920, 2002) {
					return false
				}
			case "iyr":
				iyr, err := strconv.Atoi(v)
				if err != nil {
					return false
				}
                if !inBetween(iyr, 2010, 2020) {
					return false
				}
			case "eyr":
				eyr, err := strconv.Atoi(v)
				if err != nil {
					return false
				}
                if !inBetween(eyr, 2020, 2030) {
					return false
				}
			case "hgt":
                matched, err := regexp.MatchString(`^[0-9]*(cm|in)$`, v)
                if !matched || err != nil {
                    return false
                }

                hgt, _ := strconv.Atoi(v[0:len(v)-2])
                if strings.Contains(v, "cm") && !inBetween(hgt, 150, 193) {
                    return false
                }

                if strings.Contains(v, "in") && !inBetween(hgt, 59, 76) {
                    return false
                }
			case "hcl":
                matched, err := regexp.MatchString(`^#[a-z0-9]{6}$`, v)
                if !matched || err != nil {
                    return false
                }
			case "ecl":
				if v != "amb" && v != "blu" && v != "brn" && v != "gry" && v != "grn" && v != "hzl" && v != "oth" {
					return false
				}
			case "pid":
                matched, err := regexp.MatchString(`^[0-9]{9}$`, v)
                if !matched || err != nil {
                    return false
                }
			default:
                continue
			}

		}
	}

	return true
}

func inBetween(v, lower, upper int) bool {
    return v >= lower && v <= upper
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
