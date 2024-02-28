package main

import (
	"fmt"
	"sort"
	"strings"
)

// Given an array of strings strs, group the anagrams together. You can return the answer in any order.
// An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase,
// typically using all the original letters exactly once.
// Example 1: Input: strs = ["eat","tea","tan","ate","nat","bat"]
// Output: [["bat"],["nat","tan"],["ate","eat","tea"]]
// Example 2: Input: strs = [""]
// Output: [[""]]
// Example 3: Input: strs = ["a"]
// Output: [["a"]]

func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)
	for _, str := range strs {
		s := strings.Split(str, "")
		sort.Strings(s)
		// fmt.Printf("sorted: %v\n", s)
		sortedStr := strings.Join(s, "")
		m[sortedStr] = append(m[sortedStr], str)
	}

	// fmt.Printf("m: %v\n", m)

	var res [][]string
	for _, v := range m {
		res = append(res, v)
	}
	return res
}

func main() {
	str1 := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Printf("groupAnagrams(str1): %v\n", groupAnagrams(str1))

	str2 := []string{""}
	fmt.Printf("groupAnagrams(str2): %v\n", groupAnagrams(str2))

	str3 := []string{"a"}
	fmt.Printf("groupAnagrams(str3): %v\n", groupAnagrams(str3))
}
