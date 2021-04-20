/*
String Anagrams (hard)
Given a string and a pattern, find all anagrams of the pattern in the given string.

Anagram is actually a Permutation of a string. For example, “abc” has the following six anagrams:

abc
acb
bac
bca
cab
cba
Write a function to return a list of starting indices of the anagrams of the pattern in the given string.

Example 1:

Input: String="ppqp", Pattern="pq"
Output: [1, 2]
Explanation: The two anagrams of the pattern in the given string are "pq" and "qp".

Example 2:

Input: String="abbcabc", Pattern="abc"
Output: [2, 3, 4]
Explanation: The three anagrams of the pattern in the given string are "bca", "cab", and "abc".
*/

package main

import "fmt"

func mapsEqual(map1, map2 map[rune]int) bool {
	if len(map1) != len(map2) {
		return false
	}

	for k, v := range map1 {
		if map2[k] != v {
			return false
		}
	}
	return true
}

func stringAnagram(a, b string) []int {

	sol := make([]int, 0, len(a))

	map1 := make(map[rune]int)
	map2 := make(map[rune]int)

	for _, c := range b {
		map2[c]++
	}

	l := len(b)

	for i, c := range a {
		if i >= l {
			t := rune(a[i-l])
			map1[t]--
			if map1[t] == 0 {
				delete(map1, t)
			}
		}
		map1[c]++

		if mapsEqual(map1, map2) {
			sol = append(sol, i-l+1)
		}
	}

	return sol
}

func main() {

	fmt.Println(stringAnagram("ppqp", "pq"))
	fmt.Println(stringAnagram("abbcabc", "abc"))

}
