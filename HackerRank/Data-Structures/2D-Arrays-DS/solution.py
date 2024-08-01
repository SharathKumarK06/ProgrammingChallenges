#!/bin/python3

import math
import os
import random
import re
import sys
from functools import reduce

#
# Complete the 'hourglassSum' function below.
#
# The function is expected to return an INTEGER.
# The function accepts 2D_INTEGER_ARRAY arr as parameter.
#

def hourglassSum(arr):
    sums = []
    for idy in range(0, len(arr)-2):
        for idx in range(0, len(arr[0])-2):
            # 00 + 01 + 02
            #    + 11 +
            # 20 + 21 + 22
            sumi = 0
            for i in range(3):
                sumi += arr[idy][idx+i]
                sumi += arr[idy+2][idx+i]
            sumi += arr[idy+1][idx+1]
            sums.append(sumi)
    return reduce(lambda x, y: x if x>y else y, sums)


if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    arr = []

    for _ in range(6):
        arr.append(list(map(int, input().rstrip().split())))

    result = hourglassSum(arr)

    fptr.write(str(result) + '\n')
    print(str(result) + '\n')

    fptr.close()

