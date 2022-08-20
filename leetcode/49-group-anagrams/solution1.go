package main

import "fmt"

// I think will use the fact that alphabets have to be unique then count them
func getIndex(character rune) rune {
	return character - 96
}

func groupAnagrams(strs []string) [][]string {

	M := make(map[string][]string)
	for _, v := range strs {
		characterCounts := make([]int32, 26)
		for _, character := range v {
			characterCounts[getIndex(character)] += 1
		}

		key := ""
		for i, count := range characterCounts {
			if count > 0 {
				key += fmt.Sprintf("%d%c", count, i)
			}
		}

		if anagrams, ok := M[key]; ok {
			M[key] = append(anagrams, v)
		} else {
			M[key] = []string{v}
		}
	}
	ans := make([][]string, len(M))
	index := 0
	for _, v := range M {
		ans[index] = v
		index++
	}
	return ans
}
