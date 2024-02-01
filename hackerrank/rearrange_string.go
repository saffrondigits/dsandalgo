// Given a string, construct a new string by rearranging the original string and deleting characters as needed. Return the alphabetically largest string that can be constructed respecting a limit as to how many consecutive characters can be the same.
// Example s= 'baccc"
// k=2
// The largest string, alphabetically, is 'coca' but it is not allowed because it uses the character*& more than 2 times consecutively. Therefore, the answer is 'cebca!
// Function Description
// Complete the function getLargestString in the editor below.
// getLargestString has the following parameters:
// string [nJ: the original string
// int k: the maximum number of identical consecutive characters the new string can have
// Returns:
// string: the alphabetically largest string that can be constructed that has no more than k identical consecutive characters
// Constraints
// • 13ns 109.
// • 15k≤109
// • The string s contains only lowercase English letters.ay

// package main

// func getLargestString (s string, k int32) string {
// 	// Write your code here
// }

// Simple input 0
// STDIN
// _______
// zzzazz
// 2

// Sample output 0
// zzazz

// Explanation: One 'z' must be removed so no more than 2 consecutive chare the same

// Simple input 1
// STDIN
// _______
// axxzzx
// 2

// Sample output 1
// zzxxax

// Explanation: The character 'a' must seperate the 2 'x' characters so that no more than 2 consecutive characters are the same

// input
// _______
// azzzzz
// 2

// Code Output(stdout)
// zzazzz

// Expected Output
// zzazz

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func getLargestString(s string, k int32) string {
	count := make([]int, 26)
	for _, ch := range s {
		count[ch-'a']++
	}

	pairs := make([]pair, 0, 26)
	for i, c := range count {
		if c > 0 {
			pairs = append(pairs, pair{byte(i + 'a'), c})
		}
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].count > pairs[j].count || (pairs[i].count == pairs[j].count && pairs[i].ch > pairs[j].ch)
	})

	res := make([]byte, 0, len(s))
	last := make([]byte, 0, len(s))
	for len(pairs) > 0 {
		added := false
		for i, p := range pairs {
			if p.count > 0 && (len(last) < int(k) || last[len(last)-int(k)] != p.ch) {
				res = append(res, p.ch)
				last = append(last, p.ch)
				pairs[i].count--
				added = true
				break
			}
		}
		if !added {
			if len(last) > 0 {
				last = last[:len(last)-1]
			} else {
				break
			}
		} else {
			sort.Slice(pairs, func(i, j int) bool {
				return pairs[i].count > pairs[j].count || (pairs[i].count == pairs[j].count && pairs[i].ch > pairs[j].ch)
			})
		}
	}

	return string(res)
}

type pair struct {
	ch    byte
	count int
}

func main1() { // TODO: change to main from main1
	reader := bufio.NewReader(os.Stdin)
	s, _ := reader.ReadString('\n')
	s = strings.TrimSpace(s)
	kStr, _ := reader.ReadString('\n')
	k, _ := strconv.Atoi(strings.TrimSpace(kStr))
	fmt.Println(getLargestString(s, int32(k)))
}
