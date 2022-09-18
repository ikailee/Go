package main

import "fmt"

func main() {
	ages := make(map[string]int)
	ages["Kevin"] = 40
	fmt.Printf("Kevin is %d", ages["Kevin"])

	grades := map[string]float32{
		"Kevin": 9.99,
	}
	fmt.Printf("Kevin is %.2f", grades["Kevin"])

	var visit map[string]bool
	visit = make(map[string]bool)
	visit["A"] = true

	fruits := []string{
		"apple",
		"banana",
		"kiwi",
		"apple",
	}
	fmt.Printf("duplicated furit %s", findDuplicated(fruits))
}

func findDuplicated(words []string) string {
	dupMap := make(map[string]bool)

	for i := 0; i < len(words); i++ {
		if !dupMap[words[i]] {
			dupMap[words[i]] = true
		} else {
			return words[i]
		}
	}

	return ("")
}
