#!/bin/python3

import math
import os
import random
import re
import sys
import operator

#
# Complete the 'nonDivisibleSubset' function below.
#
# The function is expected to return an INTEGER.
# The function accepts following parameters:
#  1. INTEGER k
#  2. INTEGER_ARRAY s
#

def nonDivisibleSubset(k, s):
    # Create an array to store the remainder count for each number
    remainderCount = [0] * k
    for num in s:
        remainderCount[num % k] += 1

    # Initialize the size of the non-divisible subset
    subsetSize = 0

    # If there is a remainder of 0, we can only include one number
    if remainderCount[0] > 0:
        subsetSize += 1

    # If k is even, we can only include one number with a remainder of k/2
    if k % 2 == 0 and remainderCount[k // 2] > 0:
        subsetSize += 1

    # Iterate through the remainder counts and add the larger count for each pair of remainders
    for i in range(1, (k + 1) // 2):
        subsetSize += max(remainderCount[i], remainderCount[k - i])
    return subsetSize

    
if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')
    first_multiple_input = input().rstrip().split()
    n = int(first_multiple_input[0])
    k = int(first_multiple_input[1])
    s = list(map(int, input().rstrip().split()))
    result = nonDivisibleSubset(k, s)
    fptr.write(str(result) + '\n')
    fptr.close()