/*
Problem Statement #
Given a string, find the length of the longest substring in it with no more than K distinct characters.

Example 1:

Input: String="araaci", K=2
Output: 4
Explanation: The longest substring with no more than '2' distinct characters is "araa".

Example 2:

Input: String="araaci", K=1
Output: 2
Explanation: The longest substring with no more than '1' distinct characters is "aa".

Example 3:

Input: String="cbbebi", K=3
Output: 5
Explanation: The longest substrings with no more than '3' distinct characters are "cbbeb" & "bbebi".
*/

package main

import "fmt"

func longestSubstringWithKDistinct(s string, k int) int {
	var start int
	var maximumSize int

	m := make(map[string]int)

	for i, c := range s {
		m[string(c)]++
		for len(m) > k {
			temp := string(s[start])
			m[temp]--
			if m[temp] == 0 {
				delete(m, temp)
			}
			start++
		}

		if i-start+1 >= maximumSize {
			maximumSize = i - start + 1
		}
	}

	return maximumSize
}

func main() {

	s1 := "araaci"
	s2 := "araaci"
	s3 := "cbbebi"

	fmt.Println(longestSubstringWithKDistinct(s1, 2))
	fmt.Println(longestSubstringWithKDistinct(s2, 1))
	fmt.Println(longestSubstringWithKDistinct(s3, 3))

}

/*

Time Complexity
The time complexity of the above algorithm will be O(N) where ‘N’ is the number of characters in the input string.
The outer for loop runs for all characters and the inner while loop processes each character only once, therefore the time complexity of the algorithm will be O(N+N) which is asymptotically equivalent to O(N).

Space Complexity
The space complexity of the algorithm is O(K), as we will be storing a maximum of ‘K+1’ characters in the HashMap.
*/
