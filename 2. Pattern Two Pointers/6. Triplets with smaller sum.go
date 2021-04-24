/*
Problem Statement
Given an array arr of unsorted numbers and a target sum, count all triplets in it such that arr[i] + arr[j] + arr[k] < target where i, j, and k are three different indices. Write a function to return the count of such triplets.

Example 1:

Input: [-1, 0, 2, 3], target=3
Output: 2
Explanation: There are two triplets whose sum is less than the target: [-1, 0, 3], [-1, 0, 2]

Example 2:

Input: [-1, 4, 2, 1, 3], target=5
Output: 4
Explanation: There are four triplets whose sum is less than the target:
   [-1, 1, 4], [-1, 1, 3], [-1, 1, 2], [-1, 2, 3]
*/

package main

import (
	"fmt"
	"sort"
)

func tripletsWithSmallerSum(arr []int, target int) int{

	sort.Ints(arr)

	var count int

	for i,c := range arr {

		start := i+1
		end := len(arr)-1

		for start < end {
			if(arr[start]+ arr[end]+c < target) {
				count += (end - start)
				start++
			} else {
				end--
			}
		}

	} 

	return count
}

func main() {

	arr1 := []int {-1, 0, 2, 3}
	arr2 := []int {-1, 4, 2, 1, 3}

	fmt.Println(tripletsWithSmallerSum(arr1, 3))
	fmt.Println(tripletsWithSmallerSum(arr2, 5))

}