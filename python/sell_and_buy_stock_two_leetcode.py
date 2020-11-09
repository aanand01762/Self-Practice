# link: https://leetcode.com/explore/challenge/card/30-day-leetcoding-challenge/528/week-1/3287/

class Solution:
    def maxProfit(self, prices: List[int]) -> int:

        # solution 1
        maxprofit = 0;
        # if the there is increase in the price of stock than
        # the prevoius day then add the difference to the profit
        for i  in range(1, len(prices)):
            if prices[i] > prices[i - 1]:
                maxprofit += prices[i] - prices[i - 1]
        return maxprofit

        '''
        Solution 2: peek and valley approach

        if len(prices) < 1:
            return 0

        i = maxprofit = 0
        peak = valley = prices[0]
        while i < len(prices) - 1:

            # Get the valley in the price of stock in continous fashion
            while i < len(prices) - 1 and prices[i] >= prices[i+1]:
                i += 1
            valley = prices[i]

             # Get the peak in the price of stock in continous fashion
            while i < len(prices) - 1 and prices[i] <= prices[i+1]:
                i += 1
            peak = prices[i]

            # Calcutate the profit and add to max profit
            maxprofit += peak -valley
        return maxprofit
        '''
