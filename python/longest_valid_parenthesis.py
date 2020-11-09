# https://leetcode.com/problems/longest-valid-parentheses

class Solution:
    def longestValidParentheses(self, s: str) -> int:
        stack, t_max = [-1], 0
        for i in range(len(s)):
            if s[i] == ')' and s[stack[-1]] == '(' and stack[-1] != -1:
                stack.pop(-1)
                t_max = max((i -stack[-1]), t_max)
            else:
                stack.append(i)

        return t_max
