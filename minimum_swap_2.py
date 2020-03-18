# https://www.hackerrank.com/challenges/minimum-swaps-2/problem

def minimumSwaps(arr):
    swap = 0
    index = [0]*len(arr)
    for i, value in enumerate(arr):
        index[value-1] = i
    
    for i in range(len(arr)):
        if arr[i] != i+1:
            Index = index[i]
            arr[Index] = arr[i]
            arr[i] = i+1

            index[arr[Index]-1] = Index
            index[i] = i
            swap += 1
    return swap
