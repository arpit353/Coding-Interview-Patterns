/*
Problem Statement
Given a string with lowercase letters only, if you are allowed to replace no more than ‘k’ letters with any letter, find the length of the longest substring having the same letters after replacement.

Example 1:

Input: String="aabccbb", k=2
Output: 5
Explanation: Replace the two 'c' with 'b' to have a longest repeating substring "bbbbb".

Example 2:

Input: String="abbcb", k=1
Output: 4
Explanation: Replace the 'c' with 'b' to have a longest repeating substring "bbbb".

Example 3:

Input: String="abccde", k=1
Output: 3
Explanation: Replace the 'b' or 'd' with 'c' to have the longest repeating substring "ccc".
*/

package main

import "fmt"

func characterReplacement(s string, k int) int {

	var maxSize, start, maxWindowChar int

	m := make(map[rune]int)

	for i, c := range s {
		m[c]++
		if m[c] > maxWindowChar {
			maxWindowChar = m[c]
		}

		if i-start+1-maxWindowChar > k {
			m[rune(s[start])]--
			start++
		}

		if i-start+1 > maxSize {
			maxSize = i - start + 1
		}
	}

	return maxSize
}

func characterReplacementWithoutMap(s string, k int) int {

	var maxSize, start, maxWindowChar int

	m := make([]int, 52)

	for i, c := range s {
		m[c-'A']++
		if m[c-'A'] > maxWindowChar {
			maxWindowChar = m[c-'A']
		}

		if i-start+1-maxWindowChar > k {
			m[s[start]-'A']--
			start++
		}

		if i-start+1 > maxSize {
			maxSize = i - start + 1
		}
	}

	return maxSize
}

func main() {

	str1 := "aabccbb"
	str2 := "abbcb"
	str3 := "abccde"

	fmt.Println(characterReplacement(str1, 2))
	fmt.Println(characterReplacement(str2, 1))
	fmt.Println(characterReplacement(str3, 1))

	fmt.Println(characterReplacementWithoutMap(str1, 2))
	fmt.Println(characterReplacementWithoutMap(str2, 1))
	fmt.Println(characterReplacementWithoutMap(str3, 1))
}

/*

Time Complexity
The time complexity of the above algorithm will be O(N) where ‘N’ is the number of letters in the input string.

Space Complexity
As we are expecting only the lower case letters in the input string, we can conclude that the space complexity will be O(26), to store each letter’s frequency in the HashMap, which is asymptotically equal to O(1).

*/
