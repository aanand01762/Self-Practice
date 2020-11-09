# Link: https://leetcode.com/explore/challenge/card/30-day-leetcoding-challenge/528/week-1/3286/

class Solution:
    def moveZeroes(self, nums: List[int]) -> None:
        for i in range(len(nums)):
            # If num is zero, swap it with next non-zero num
            if nums[i] == 0:
                for nexti in range(i+1, len(nums)):
                    if nums[nexti] != 0:
                        temp = nums[i]
                        nums[i] = nums[nexti]
                        nums[nexti] = temp
                        break
        return nums
