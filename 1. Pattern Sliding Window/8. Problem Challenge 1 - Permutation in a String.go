/*
Problem Challenge 1

Permutation in a String (hard)
Given a string and a pattern, find out if the string contains any permutation of the pattern.

Permutation is defined as the re-arranging of the characters of the string. For example, “abc” has the following six permutations:

abc
acb
bac
bca
cab
cba
If a string has ‘n’ distinct characters it will have n!n! permutations.

Example 1:

Input: String="oidbcaf", Pattern="abc"
Output: true
Explanation: The string contains "bca" which is a permutation of the given pattern.

Example 2:

Input: String="odicf", Pattern="dc"
Output: false
Explanation: No permutation of the pattern is present in the given string as a substring.

Example 3:

Input: String="bcdxabcdy", Pattern="bcdyabcdx"
Output: true
Explanation: Both the string and the pattern are a permutation of each other.

Example 4:

Input: String="aaacb", Pattern="abc"
Output: true
Explanation: The string contains "acb" which is a permutation of the given pattern.
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

func findPermutation(str string, pattern string) bool {

	map1 := make(map[rune]int)
	map2 := make(map[rune]int)

	for _, c := range pattern {
		map2[c]++
	}
	l := len(pattern)
	for i, c := range str {
		map1[c]++
		if i >= l {
			t := rune(str[i-l])
			map1[t]--
			if map1[t] == 0 {
				delete(map1, t)
			}
		}
		if mapsEqual(map1, map2) {
			return true
		}
	}

	return false
}

func main() {

	fmt.Println(findPermutation("oidbcaf", "abc"))
	fmt.Println(findPermutation("odicf", "dc"))
	fmt.Println(findPermutation("bcdxabcdy", "bcdyabcdx"))
	fmt.Println(findPermutation("aaacb", "abc"))

}
