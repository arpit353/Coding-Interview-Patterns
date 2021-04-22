/*
Problem Statement
Given an array of unsorted numbers, find all unique triplets in it that add up to zero.

Example 1:

Input: [-3, 0, 1, 2, -1, 1, -2]
Output: [-3, 1, 2], [-2, 0, 2], [-2, 1, 1], [-1, 0, 1]
Explanation: There are four unique triplets whose sum is equal to zero.

Example 2:

Input: [-5, 2, -1, -2, 3]
Output: [[-5, 2, 3], [-2, -1, 3]]
Explanation: There are two unique triplets whose sum is equal to zero.
*/

package main

import (
	"fmt"
	"sort"
)

func tripletsSumToZero(nums []int) [][]int{

	var results [][]int
	sort.Ints(nums)

	for i := 0; i < len(nums)-2; i++ {

		// Check for repeat
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		target, left, right := -nums[i], i+1, len(nums)-1
		
		// Find Pairs
		for left < right {
			
			sum := nums[left] + nums[right]
			
			if sum == target {
				
				results = append(results, []int{nums[i], nums[left], nums[right]})
				
				left++
				right--
				
				// Check for repeat
				for left < right && nums[left] == nums[left-1] {
					left++
				}
				
				// Check for repeat
				for left < right && nums[right] == nums[right+1] {
					right--
				}
			} else if sum > target {

				right--

			} else if sum < target {
				
				left++
				
			}
		}
	}
	
	return results
}


func main() {

	arr1 := []int {-3, 0, 1, 2, -1, 1, -2}
	arr2 := []int {-5, 2, -1, -2, 3}

	fmt.Println(tripletsSumToZero(arr1))
	fmt.Println(tripletsSumToZero(arr2))

}