package helpers

import (
	"sort"
)

//CountDuplicates func counts duplicates in string slice
func CountDuplicates(strSlice []string) map[string]float64 {
	duplicate := map[string]float64{}

	for _, v := range strSlice {
		_, exist := duplicate[v]

		if exist {
			duplicate[v]++
		} else {
			duplicate[v] = 1
		}
	}

	return duplicate
}

//CountLanguagesCommit counts duplicates and adds up commits values
func CountLanguagesCommit(strSlice []string, floatSlice []float64) map[string]float64 {
	duplicate := map[string]float64{}

	for i, v := range strSlice {
		_, exist := duplicate[v]

		if exist {
			duplicate[v] += floatSlice[i]
		} else {
			duplicate[v] = floatSlice[i]
		}
	}

	return duplicate
}

//CalcStarsOrForks iterates over slice and adds up it's values
func CalcStarsOrForks(strings []string, integers []float64) map[string]float64 {
	newMap := map[string]float64{}

	for i, v := range strings {
		newMap[v] += integers[i]
	}

	return newMap
}

//SortMap splits map into two slices and sorts them
func SortMap(someMap map[string]float64) (strings []string, integers []float64) {
	var strSlice []string
	var intSlice []float64

	keys := make([]string, 0, len(someMap))
	for key := range someMap {
		keys = append(keys, key)
	}

	sort.Slice(keys, func(i, j int) bool {
		return someMap[keys[i]] > someMap[keys[j]]
	})

	for _, key := range keys {
		intSlice = append(intSlice, someMap[key])
		strSlice = append(strSlice, key)
	}

	return strSlice, intSlice
}
