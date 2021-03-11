package helpers

import (
	"sort"

	"github.com/shurcooL/githubv4"
)

//CountDuplicates func counts duplicates in string slice
func CountDuplicates(strSlice []string) map[string]int {
	duplicate := map[string]int{}

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

func CountLanguagesCommit(strSlice []string, intSlice []githubv4.Int) map[string]githubv4.Int {
	duplicate := map[string]githubv4.Int{}

	for i, v := range strSlice {
		_, exist := duplicate[v]

		if exist {
			duplicate[v] += intSlice[i]
		} else {
			duplicate[v] = intSlice[i]
		}
	}

	return duplicate
}

//CalcStarsOrForks iterates over slice and calculates
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
