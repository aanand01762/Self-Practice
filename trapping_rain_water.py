class Solution:
    def trap(self, height: List[int]) -> int:
        n, water = len(height), 0
        for i in range(n):
            left = right = i
            lmax, rmax = height[i], height[i]

            while left > 0:
                left -= 1
                if lmax < height[left]:
                    lmax = height[left]


            while right < n-1:
                right += 1
                if rmax < height[right]:
                    rmax = height[right]


            water += (min(lmax,rmax) - height[i])
        return water
