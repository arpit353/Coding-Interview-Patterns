/*
Problem Statement #
Given an array of positive numbers and a positive number ‘k’, find the maximum sum of any contiguous subarray of size ‘k’.

Example 1:

Input: [2, 1, 5, 1, 3, 2], k=3
Output: 9
Explanation: Subarray with maximum sum is [5, 1, 3].

Example 2:

Input: [2, 3, 4, 1, 5], k=2
Output: 7
Explanation: Subarray with maximum sum is [3, 4].
*/

package main

import (
	"fmt"
)

func maxSubArrayOfSizeK(arr []int, windowSize int) int {

	var maximumSum, windowSum, start, end int

	for end = 0; end < len(arr); end++ {
		windowSum += arr[end]
		if end >= windowSize-1 {
			if maximumSum <= windowSum {
				maximumSum = windowSum
			}
			windowSum -= arr[start]
			start++
		}
	}

	return maximumSum
}

func main() {
	arr1 := []int{2, 1, 5, 1, 3, 2}
	arr2 := []int{2, 3, 4, 1, 5}

	fmt.Println(maxSubArrayOfSizeK(arr1, 3))
	fmt.Println(maxSubArrayOfSizeK(arr2, 2))
}

/*
Time Complexity
The time complexity of the above algorithm will be O(N).

Space Complexity
The algorithm runs in constant space O(1).
*/
