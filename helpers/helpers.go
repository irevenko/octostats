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

//CountLanguagesCommit counts duplicates and adds up commits values
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

//CalcStarsOrForks iterates over slice and adds up it's values
func CalcStarsOrForks(strings []string, integers []int) map[string]int {
	newMap := map[string]int{}

	for i, v := range strings {
		newMap[v] += integers[i]
	}

	return newMap
}

//SortMap splits map into two slices and sorts them
func SortMap(someMap map[string]int) (strings []string, integers []int) {
	var strSlice []string
	var intSlice []int

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

func SortIntv4Map(someMap map[string]githubv4.Int) (strings []string, integers []githubv4.Int) {
	var strSlice []string
	var intSlice []githubv4.Int

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
