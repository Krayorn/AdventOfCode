package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

type camera struct {
	ID        int
	sides     map[string]string
	image     []string
	neighbors map[int]bool
}

func main() {
	textFile, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println(err)
	}

	byteValue, _ := ioutil.ReadAll(textFile)
	values := strings.Split(string(byteValue), "\n")

	allIngredients := make([]map[string]bool, len(values))
	allAllergens := make([]map[string]bool, len(values))

	for n, v := range values {
		split := strings.Split(v, " (contains ")
		ingredients := make(map[string]bool)
		allergens := make(map[string]bool)
		for _, i := range strings.Split(split[0], " ") {
			ingredients[i] = true
		}
		for _, a := range strings.Split(strings.TrimSuffix(split[1], ")"), ", ") {
			allergens[a] = true
		}

		allIngredients[n] = ingredients
		allAllergens[n] = allergens
	}

	allergensFound := make(map[string][]string)

	interact := true
	for interact {
		interact = false
		for i, allergens := range allAllergens {
			for allergen := range allergens {
				var intersect []string
				for ingredient := range allIngredients[i] {
					intersect = append(intersect, ingredient)
				}
				for j, allergensToCompare := range allAllergens {
					if i == j {
						continue
					}
					for allergenToCompare := range allergensToCompare {
						if allergenToCompare == allergen {
							for k := 0; k < len(intersect); k++ {
								if _, ok := allIngredients[j][intersect[k]]; !ok {
									intersect = remove(intersect, k)
									k--
								}
							}
						}
					}
				}
				for allergenToCompare, matchingIngredients := range allergensFound {
					if allergenToCompare == allergen {
						continue
					}
					if len(matchingIngredients) == 1 {
						if ok := stringInSlice(matchingIngredients[0], intersect); ok != -1 {
							intersect = remove(intersect, ok)
						}
					}
				}
				if ingredients, ok := allergensFound[allergen]; !ok || (len(intersect) != len(ingredients)) {
					interact = true
				}
				allergensFound[allergen] = intersect
			}
		}
	}

	total := 0
	for _, ingredients := range allIngredients {
		for ingredient := range ingredients {
			safe := true
			for _, ingredientsToCompare := range allergensFound {
				if ok := stringInSlice(ingredient, ingredientsToCompare); ok != -1 {
					safe = false
				}
			}
			if safe {
				total++
			}
		}
	}

	fmt.Println(total)

	allergenAlpha := make([]string, 0)
	for allergen := range allergensFound {
		allergenAlpha = append(allergenAlpha, allergen)
	}

	sort.Strings(allergenAlpha)

	var res []string
	for _, allergen := range allergenAlpha {
		res = append(res, allergensFound[allergen][0])
	}

	fmt.Println(strings.Join(res, ","))
}

func stringInSlice(a string, list []string) int {
	for i, b := range list {
		if b == a {
			return i
		}
	}
	return -1
}

func remove(s []string, i int) []string {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
