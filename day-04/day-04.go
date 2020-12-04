package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	readPassports("./input")
}

func readPassports(inputFile string) {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	validPassportsP1 := 0
	validPassportsP2 := 0
	newPassport := make(map[string]string)
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			// Check the old map to see if its valid
			if isValidPassportP1(newPassport) {
				validPassportsP1 += 1
			}
			if isValidPassportP2(newPassport) {
				validPassportsP2 += 1
			}
			// Re-set the map and start verifying a new passport
			newPassport = make(map[string]string)
		} else {
			entries := strings.Split(line, " ")
			for _, e := range entries {
				entry := strings.Split(e, ":")
				key := entry[0]
				value := entry[1]
				newPassport[key] = value
			}
		}
	}

	// Just in case the last passport was checked before ending the for loop
	if isValidPassportP1(newPassport) {
		validPassportsP1 += 1
	}

	if isValidPassportP2(newPassport) {
		validPassportsP2 += 1
	}

	fmt.Printf("There are %s valid passports\n", strconv.Itoa(validPassportsP1))
	fmt.Printf("There are %s valid passports\n", strconv.Itoa(validPassportsP2))
}

func isValidPassportP1(passport map[string]string) bool {
	_, byr := passport["byr"]
	_, iyr := passport["iyr"]
	_, eyr := passport["eyr"]
	_, hgt := passport["hgt"]
	_, hcl := passport["hcl"]
	_, ecl := passport["ecl"]
	_, pid := passport["pid"]
	return byr &&
		iyr &&
		eyr &&
		hgt &&
		hcl &&
		ecl &&
		pid
}

func isValidPassportP2(passport map[string]string) bool {
	byrValue, byr := passport["byr"]
	iyrValue, iyr := passport["iyr"]
	eyrValue, eyr := passport["eyr"]
	hgtValue, hgt := passport["hgt"]
	hclValue, hcl := passport["hcl"]
	eclValue, ecl := passport["ecl"]
	pidValue, pid := passport["pid"]
	return byr && checkByr(byrValue) &&
		iyr && checkIyr(iyrValue) &&
		eyr && checkEyr(eyrValue) &&
		hgt && checkHeight(hgtValue) &&
		hcl && checkHcl(hclValue) &&
		ecl && checkEcl(eclValue) &&
		pid && checkPid(pidValue)
}

func checkByr(byrValue string) bool {
	return byrValue >= "1920" && byrValue <= "2002" && len(byrValue) == 4
}

func checkIyr(iyrValue string) bool {
	return iyrValue >= "2010" && iyrValue <= "2020" && len(iyrValue) == 4
}

func checkEyr(eyrValue string) bool {
	return eyrValue >= "2020" && eyrValue <= "2030" && len(eyrValue) == 4
}

func checkHeight(height string) bool {
	if strings.HasSuffix(height, "cm") {
		num := strings.Replace(height, "cm", "", 1)
		return num >= "150" && num <= "193"
	} else if strings.HasSuffix(height, "in") {
		num := strings.Replace(height, "in", "", 1)
		return num >= "59" && num <= "76"
	}
	return false
}

func checkHcl(hclValue string) bool {
	match, _ := regexp.MatchString("^#[a-f0-9]{6}$", hclValue)
	return match
}

func checkEcl(eclValue string) bool {
	return eclValue == "amb" ||
		eclValue == "blu" ||
		eclValue == "brn" ||
		eclValue == "gry" ||
		eclValue == "grn" ||
		eclValue == "hzl" ||
		eclValue == "oth"
}

func checkPid(pidValue string) bool {
	match, _ := regexp.MatchString("^[0-9]{9}$", pidValue)
	return match
}
