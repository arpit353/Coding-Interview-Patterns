/*
Problem Statement
Given an array containing 0s and 1s, if you are allowed to replace no more than ‘k’ 0s with 1s, find the length of the longest contiguous subarray having all 1s.

Example 1:

Input: Array=[0, 1, 1, 0, 0, 0, 1, 1, 0, 1, 1], k=2
Output: 6
Explanation: Replace the '0' at index 5 and 8 to have the longest contiguous subarray of 1s having length 6.

Example 2:

Input: Array=[0, 1, 0, 0, 1, 1, 0, 1, 1, 0, 0, 1, 1], k=3
Output: 9
Explanation: Replace the '0' at index 6, 9, and 10 to have the longest contiguous subarray of 1s having length 9.
*/

package main

import "fmt"

func lengthOfLongestSubstring(arr []int, k int) int {
	var start, maxWindow, numZeros int
	for i, c := range arr {
		if c == 0 {
			numZeros++
		}
		for numZeros > k {
			if arr[start] == 0 {
				numZeros--
			}
			start++
		}
		if i-start+1 > maxWindow {
			maxWindow = i - start + 1
		}
	}
	return maxWindow
}

func main() {

	arr1 := []int{0, 1, 1, 0, 0, 0, 1, 1, 0, 1, 1}
	arr2 := []int{0, 1, 0, 0, 1, 1, 0, 1, 1, 0, 0, 1, 1}

	fmt.Println(lengthOfLongestSubstring(arr1, 2))
	fmt.Println(lengthOfLongestSubstring(arr2, 3))
}

/*
Time Complexity
The time complexity of the above algorithm will be O(N) where ‘N’ is the count of numbers in the input array.

Space Complexity
The algorithm runs in constant space O(1).
*/
