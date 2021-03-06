# https://www.hackerrank.com/challenges/counting-valleys/problem

#!/bin/python3

import math
import os
import random
import re
import sys

# Complete the countingValleys function below.
def countingValleys(n, s):
    c ,sum = 0, 0
    for i in range(n):
        if sum == 0 and s[i] == 'D':
            c += 1
            sum -= 1
        elif s[i] == 'D':
            sum -= 1
        else:
            sum += 1
    return  c
if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    n = int(input())

    s = input()

    result = countingValleys(n, s)

    fptr.write(str(result) + '\n')

    fptr.close()
