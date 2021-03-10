package helpers

import "sort"

func CountDuplicates(strSlice []string) map[string]int {
	duplicate := map[string]int{}

	for _, item := range strSlice {
		_, exist := duplicate[item]

		if exist {
			duplicate[item]++
		} else {
			duplicate[item] = 1
		}
	}

	return duplicate
}

func CalcStarsOrForks(strings []string, integers []int) map[string]int {
	newMap := map[string]int{}

	for i, v := range strings {
		newMap[v] += integers[i]
	}

	return newMap
}

func SortMap(someMap map[string]int) (strings []string, integers []int) {
	var strVal []string
	var intVal []int

	keys := make([]string, 0, len(someMap))
	for key := range someMap {
		keys = append(keys, key)
	}

	sort.Slice(keys, func(i, j int) bool {
		return someMap[keys[i]] > someMap[keys[j]]
	})

	for _, key := range keys {
		intVal = append(intVal, someMap[key])
		strVal = append(strVal, key)
	}

	return strVal, intVal
}
