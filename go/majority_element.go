//link: https://leetcode.com/problems/majority-element/

package main

import "fmt"

func majorityElement(nums []int) int {
	elements_occurance := make(map[int]int)

	// Initialize the maximum occured element to the first element at start
	max_occuring_num := nums[0]

	// Iterate through each element and increment the occurence of element
	// and compare it with privously max occured element
	for _, i := range nums {
		elements_map[i]++
		if elements_occurance[i] > elements_occurance[max_occuring_num] {
			max_occuring_num = i
		}
	}
	return max_occuring_num
}
