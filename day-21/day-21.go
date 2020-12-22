package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func main() {
	allerginMapToSet, ingredientCount := parseInput()
	foundAllergins := eliminateAllergins(copy(allerginMapToSet))
	solve(foundAllergins, ingredientCount)
	solveP2(foundAllergins)
}

func parseInput() (map[string][]string, map[string]int) {
	file, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	allerginMapToSet := make(map[string][]string)
	ingredientCount := make(map[string]int)

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.ReplaceAll(line, ")", "")
		splitInput := strings.Split(line, " (contains ")
		ingredients := strings.Split(splitInput[0], " ")
		allergens := strings.Split(splitInput[1], ", ")

		for _, v := range allergens {
			if _, ok := allerginMapToSet[v]; ok {
				allerginMapToSet[v] = filter(ingredients, func(test string) bool {
					val, ok := allerginMapToSet[v]
					if !ok {
						return false
					}
					return contains(val, test)
				})
			} else {
				allerginMapToSet[v] = ingredients
			}
		}

		for _, v := range ingredients {
			if val, ok := ingredientCount[v]; ok {
				ingredientCount[v] = val + 1
			} else {
				ingredientCount[v] = 1
			}
		}

	}
	return allerginMapToSet, ingredientCount
}

func eliminateAllergins(allerginMapToSet map[string][]string) map[string]string {
	foundAllergens := make(map[string]string)
	keys := mapKeysToArray(allerginMapToSet)
	for len(keys) > 0 {
		// Find findIndex key with size === 1 and we can delete that
		item := findIngredientWithOneItem(keys, allerginMapToSet)
		ingredient := allerginMapToSet[item][0]
		foundAllergens[item] = ingredient
		delete(allerginMapToSet, item)

		for k, v := range allerginMapToSet {
			ind := findIndex(v, ingredient)
			if ind != -1 {
				allerginMapToSet[k] = removeIndex(allerginMapToSet[k], ind)
			}
		}

		keys = mapKeysToArray(allerginMapToSet)
	}
	return foundAllergens
}

func solve(foundAllergins map[string]string, ingredientCount map[string]int) {
	definedIngredients := make(map[string]string)
	for _, v := range foundAllergins {
		definedIngredients[foundAllergins[v]] = v
	}
	arr := intKeys(ingredientCount)
	ret := filter(arr, func(t string) bool {
		for _, v := range foundAllergins {
			if v == t {
				return false
			}
		}
		return true
	})

	count := 0
	for _, ingred := range ret {
		count += ingredientCount[ingred]
	}
	fmt.Printf("Part 1: %d\n", count)
}

func solveP2(foundAllergins map[string]string) {
	arr := stringKeys(foundAllergins)
	sort.Strings(arr)
	sortedAllergins := make([]string, 0)
	for _, v := range arr {
		sortedAllergins = append(sortedAllergins, foundAllergins[v])
	}
	solution := strings.Join(sortedAllergins, ",")
	fmt.Printf("Part 2: %s\n", solution)
}

func intKeys(ingredientCount map[string]int) []string {
	arr := make([]string, 0)
	for k, _ := range ingredientCount {
		arr = append(arr, k)
	}
	return arr
}

func stringKeys(ingredientCount map[string]string) []string {
	arr := make([]string, 0)
	for k, _ := range ingredientCount {
		arr = append(arr, k)
	}
	return arr
}

func findIndex(arr []string, remove string) int {
	for i, v := range arr {
		if v == remove {
			return i
		}
	}
	return -1
}

func removeIndex(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}

func findIngredientWithOneItem(s []string, allerginMapToSet map[string][]string) string {
	for _, v := range s {
		if val, ok := allerginMapToSet[v]; ok {
			if len(val) == 1 {
				return v
			}
		}
	}
	return ""
}

func mapKeysToArray(allerginMapToSet map[string][]string) []string {
	keys := make([]string, len(allerginMapToSet))
	i := 0
	for k := range allerginMapToSet {
		keys[i] = k
		i++
	}
	return keys
}

func contains(slice []string, item string) bool {
	set := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		set[s] = struct{}{}
	}

	_, ok := set[item]
	return ok
}

func filter(ss []string, test func(string) bool) (ret []string) {
	for _, s := range ss {
		if test(s) {
			ret = append(ret, s)
		}
	}
	return
}

func copy(myMap map[string][]string) map[string][]string {
	newMap := make(map[string][]string)
	for k, v := range myMap {
		newMap[k] = v
	}
	return newMap
}
