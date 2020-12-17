package main

import (
	"bufio"
	"log"
	"os"
	"strings"
	"regexp"
	"strconv"
)

const MASK_LENGTH = 36

func main() {
	memory := readInput1("./input")
	var sum1 int64 = 0
	for _, m := range memory {
		sum1 = sum1 + m
	}
	log.Println("What is the sum of all values left in memory after it completes?")
	log.Println(sum1)

	memory2 := readInput2("./input")
	var sum2 int64 = 0
	for _, m := range memory2 {
		sum2 = sum2 + m
	}
	log.Println("What is the sum of all values left in memory after it completes?")
	log.Println(sum2)
}

func readInput1(inputFile string) map[string]int64 {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	memory := make(map[string]int64)
	mask := ""
	re := regexp.MustCompile(`\[([0-9]+)\] = ([0-9]+)`)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "mask") {
			mask = strings.TrimPrefix(line, "mask = ")
			// log.Println("mask changed to", mask)
		} else {
			// mem[8] = 11
			mem := re.FindStringSubmatch(line)
			memIndex := mem[1]
			memValue := mem[2]
			maskedValue := applyMask(memValue, mask)
			// log.Println(memIndex, memValue, "=>", maskedValue)
			memory[memIndex] = maskedValue
		}
	}
	return memory
}

func readInput2(inputFile string) map[string]int64 {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	memory := make(map[string]int64)
	mask := ""
	re := regexp.MustCompile(`\[([0-9]+)\] = ([0-9]+)`)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "mask") {
			mask = strings.TrimPrefix(line, "mask = ")
			// log.Println("mask changed to", mask)
		} else {
			// mem[8] = 11
			mem := re.FindStringSubmatch(line)
			memIndex := mem[1]
			memValue := mem[2]
			maskedValue, _ := strconv.ParseInt(memValue, 10, 64)
			maskedMemIndexes := maskMemIndex(memIndex, mask)
			// log.Println(memIndex, memValue, "=>", maskedValue)
			for _, memI := range maskedMemIndexes {
				// log.Println(memI, "=>", maskedValue)
				memory[memI] = maskedValue
			}
		}
	}
	return memory
}

func applyMask(memValue string, mask string) int64 {
	value, err := strconv.Atoi(memValue)
	if err != nil {
		log.Fatal(err)
	}
	valueBin := []byte(strconv.FormatInt(int64(value), 2))
	maskedValue := make([]byte, MASK_LENGTH)
	
	for i := 0; i < MASK_LENGTH; i = i + 1 {
		valueI := i - (MASK_LENGTH - len(valueBin))
		if mask[i] != 'X' {
			maskedValue[i] = mask[i]
		} else if valueI >= 0 {
			maskedValue[i] = valueBin[valueI]
		} else {
			maskedValue[i] = '0'
		}
	}
	
	// log.Println("masked bin => ", string(maskedValue))
	if result, err := strconv.ParseInt(string(maskedValue), 2, 64); err == nil {
		return result
	} else {
		log.Fatal(err)
		return -1
	}
}

func maskMemIndex(memI string, mask string) []string {
		value, err := strconv.Atoi(memI)
	if err != nil {
		log.Fatal(err)
	}
	valueBin := []byte(strconv.FormatInt(int64(value), 2))
	maskedValue := make([]byte, MASK_LENGTH)
	floating := make([]int, 0)
	
	for i := 0; i < MASK_LENGTH; i = i + 1 {
		valueI := i - (MASK_LENGTH - len(valueBin))
		if mask[i] == '1' {
			maskedValue[i] = mask[i]
		} else if mask[i] == '0' && valueI >= 0 {
			maskedValue[i] = valueBin[valueI]
		} else if mask[i] == 'X' {
			floating = append(floating, i)
			maskedValue[i] = '0'
		} else {
			maskedValue[i] = '0'
		}
	}

	floatingVariants := make([]string, 0)
	floatingVariants = append(floatingVariants, string(maskedValue))

	for _, fI := range floating {
		for _, res := range floatingVariants {
			val := []byte(res)
			val[fI] = '1'
			floatingVariants = append(floatingVariants, string(val))
		}
	}
	
	// log.Println(floatingVariants)
	result := make([]string, 0)
	for _, res := range floatingVariants {
		if i, err := strconv.ParseInt(res, 2, 64); err == nil {
			result = append(result, strconv.FormatInt(i, 10))
		} else {
			log.Fatal(err)
			return nil
		}
	}
	return result
}