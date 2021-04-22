/*
Problem Statement
Given an array of sorted numbers, remove all duplicates from it. You should not use any extra space; after removing the duplicates in-place return the new length of the array.

Example 1:

Input: [2, 3, 3, 3, 6, 9, 9]
Output: 4
Explanation: The first four elements after removing the duplicates will be [2, 3, 6, 9].

Example 2:

Input: [2, 2, 2, 11]
Output: 2
Explanation: The first two elements after removing the duplicates will be [2, 11].
*/

package main

import "fmt"

func removeDuplicates(arr []int) int {
	var length int
	start := -1 

	for i,c := range arr {
		if start == -1 || c != arr[start] {
			start = i
			arr[length] = c
			length++
		}
	}

	return length
}

func main() {

	arr1 := []int{2, 3, 3, 3, 6, 9, 9}
	arr2 := []int{2, 2, 2, 11}

	fmt.Println(removeDuplicates(arr1))
	fmt.Println(removeDuplicates(arr2))

}