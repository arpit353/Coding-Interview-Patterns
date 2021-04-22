/*
Problem Statement
Given a sorted array, create a new array containing squares of all the number of the input array in the sorted order.

Example 1:

Input: [-2, -1, 0, 2, 3]
Output: [0, 1, 4, 4, 9]

Example 2:

Input: [-3, -1, 0, 1, 2]
Output: [0 1 1 4 9]
*/

package main

import "fmt"

func getSquaredArray(arr []int) []int {

	start := 0
	l := len(arr)
	end := l - 1

	sol := make([]int, l)

	for start <= end {
		startSquare := arr[start]*arr[start]
		endSquare := arr[end]*arr[end]
		if startSquare >= endSquare {
			sol[l-1] = startSquare 
			start++
		} else {
			sol[l-1] = endSquare
			end--
		}
		l--;
	}

	return sol
}

func main() {

	arr1 := []int {-2, -1, 0, 2, 3}
	arr2 := []int {-3, -1, 0, 1, 2}

	fmt.Println(getSquaredArray(arr1))
	fmt.Println(getSquaredArray(arr2))

}