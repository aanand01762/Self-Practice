#https://leetcode.com/problems/trapping-rain-water

class Solution:
    def trap(self, height: List[int]) -> int:
        n, water = len(height), 0

        # Traverse each tower to calculate water at the perticular point
        for i in range(n):
            left = right = i

            # Intialize tower with maximum hieght to left and right of the
            # tower with tower's own length
            lmax, rmax = height[i], height[i]

            # Calulate the hieght of the maximum longest tower to left of tower
            while left > 0:
                left -= 1
                if lmax < height[left]:
                    lmax = height[left]

            # Calulate the hieght of the maximum longest tower to right of tower
            while right < n-1:
                right += 1
                if rmax < height[right]:
                    rmax = height[right]

            # Add water at that point to total water
            water += (min(lmax,rmax) - height[i])
        return water
