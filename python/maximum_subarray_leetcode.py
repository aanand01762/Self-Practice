# Link: https://leetcode.com/explore/challenge/card/30-day-leetcoding-challenge/528/week-1/3285/
# reference: https://www.youtube.com/watch?v=86CQq3pKSUw

class Solution:
    def maxSubArray(self, nums: List[int]) -> int:
        curr_max , global_max = 0, min(nums)

        for i in nums:

            # current maximum sum at perticular index can be current element
            # or the sum of elements before it and current element
            curr_max = max(i , curr_max + i)

            # If current max is more than the global max which we found
            # erlier than updat the global max
            global_max = max(global_max, curr_max)
        return global_max
