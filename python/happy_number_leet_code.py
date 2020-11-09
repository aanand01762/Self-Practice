# Link: https://leetcode.com/explore/challenge/card/30-day-leetcoding-challenge/528/week-1/3284/

class Solution:
    def isHappy(self, n: int) -> bool:

        # Create a list which will store
        # the produced numbers
        numbers = []

        while(1):
            squaresum = 0

            # Get the sum of squared digits
            while(n):
                squaresum += (n % 10) * (n % 10)
                n = int(n / 10)
            n = squaresum

            # Check if sum is 1
            if n == 1:
                return True

            # Check if the  number is not produced
            # earlier and sequence is not repeating
            if n in numbers:
                return False

            # Store the produced number
            numbers.append(n)
