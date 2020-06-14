# Link: https://leetcode.com/explore/challenge/card/30-day-leetcoding-challenge/528/week-1/3283/

class Solution:
    def singleNumber(self, nums: List[int]) -> int:
        numbers = []

        # Check if number has ocured before
        for num in nums:

            # At first occurance, add to list
            if num not in numbers:
                numbers.append(num)

            # Remove at second occurance
            else:
                numbers.remove(num)
