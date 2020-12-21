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
	rules, inputs := parseInput()
	regex := computeRegex(rules[0], rules, make(map[string]string, 0))

	totalSum := 0
	for _, v := range inputs {
		if ok, _ := regexp.MatchString("^" + regex + "$", v); ok {
			totalSum+= 1
		}
	}
	fmt.Printf("Part 1: %d\n", totalSum)

	rules[0] = "8 11"
	rules[8] = "42 | 42 8"
	rules[11] = "42 31 | 42 11 31"

	regex42 := computeRegex(rules[42], rules, make(map[string]string, 0))
	regex31 := computeRegex(rules[31], rules, make(map[string]string, 0))

	newRegex := fmt.Sprintf("^(?P<g42>(%s)+)(?P<g31>(%s)+)$", regex42, regex31)

	totalSumP2 := 0
	for _, v := range inputs {
		reg := regexp.MustCompile(newRegex)
		match := reg.FindStringSubmatch(v)
		result := make(map[string]string)
		for i, name := range reg.SubexpNames() {
			if i != 0 && name != "" && match != nil {
				result[name] = match[i]
			}
		}
		if val1, ok := result["g42"]; ok {
			if val2, ok2 := result["g31"]; ok2 {
				group42Regex := regexp.MustCompile(regex42)
				group42Matches := group42Regex.FindAllStringIndex(val1, -1)

				group31Regex := regexp.MustCompile(regex31)
				group31Matches := group31Regex.FindAllStringIndex(val2, -1)
				if len(group42Matches) > len(group31Matches) {
					totalSumP2 += 1
				}
			}
		}
	}
	fmt.Printf("Part 2: %d\n", totalSumP2)
}

func parseInput()  (map[int]string, []string){
	file, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	spacesFound := 0
	rules := make(map[int]string)
	inputs := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			spacesFound += 1
		}
		if spacesFound == 0 {
			a := strings.Split(line, ":")
			c, _ := strconv.Atoi(a[0])
			rules[c] = strings.Trim(a[1], " ")
		} else if line != "" {
			inputs = append(inputs, line)
		}
	}
	return rules, inputs
}

func computeRegex(value string, rules map[int]string, cache map[string]string) string {
	// Use the cache if its already been computed
	if val, ok := cache[value]; ok {
		return val
	}

	if strings.HasPrefix(value, "\"") {
		item := strings.ReplaceAll(value, "\"", "")
		cache[value] = item
		return item
	} else if strings.Contains(value, " | "){
		options := strings.Split(value, " | ")
		firstRegex := computeRegex(options[0], rules, cache)
		secondRegex :=  computeRegex(options[1], rules, cache)
		result := fmt.Sprintf("(%s|%s)", firstRegex, secondRegex);
		cache[value] = result
		return result
	} else {
		items := strings.Split(value, " ")
		result := ""
		for _, v := range items {
			ind, _ := strconv.Atoi(v)
			result += computeRegex(rules[ind], rules, cache)
		}
		cache[value] = result
		return result
	}
}
