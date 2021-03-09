package github

import "sort"

func countLanguages(langs []string) map[string]int {
	duplicate := map[string]int{}

	for _, item := range langs {
		_, exist := duplicate[item]

		if exist {
			duplicate[item]++
		} else {
			duplicate[item] = 1
		}
	}

	return duplicate
}

func calcStarsOrForks(langs []string, ints []int) map[string]int {
	newMap := map[string]int{}

	for i, v := range langs {
		newMap[v] += ints[i]
	}

	return newMap
}

func sortMap(langs map[string]int) (languages []string, count []int) {
	var strVal []string
	var intVal []int

	keys := make([]string, 0, len(langs))
	for key := range langs {
		keys = append(keys, key)
	}

	sort.Slice(keys, func(i, j int) bool {
		return langs[keys[i]] > langs[keys[j]]
	})

	for _, key := range keys {
		intVal = append(intVal, langs[key])
		strVal = append(strVal, key)
	}

	return strVal, intVal
}
