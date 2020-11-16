//link: https://leetcode.com/problems/count-primes/

package main

import "fmt"

func countPrimes(n int) int {
	isprime := make([]int, n)
	var count int

	// i*i will be less than n
	for i := 2; i*i < n; i++ {
		if isprime[i] == 0 {

			// Set value 0 at index which is multiple of any number and less than n
			for j := 2; j*i < n; j++ {
				isprime[i*j] = 1
			}
		}
	}

	// Count total prime numbers
	for i := 2; i < n; i++ {
		if isprime[i] == 0 {
			count++
		}
	}
	return count
}
