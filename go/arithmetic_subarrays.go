//link: https://leetcode.com/problems/arithmetic-subarrays/

package main

import "sort"

func checkArithmeticSubarrays(nums []int, l []int, r []int) []bool {
	var bools []bool
	for i, _ := range l {

		// Create a sub slice with given indices and sort that
		new_list := make([]int, r[i]+1-l[i])
		copy(new_list, nums[l[i]:(r[i]+1)])
		sort.Ints(new_list)

		// Get the diference of first two elements and intialize the flag
		diff := new_list[1] - new_list[0]
		flag := 0

		// Check if diff between any two consecutive elements is not same
		// then change flag and add bool value of false for that slice
		for j := 2; j < (r[i] + 1 - l[i]); j++ {
			if (new_list[j] - new_list[j-1]) != diff {
				flag = 1
				break
			}
		}
		if flag == 1 {
			bools = append(bools, false)
		} else {
			bools = append(bools, true)
		}
	}
	return bools
}
