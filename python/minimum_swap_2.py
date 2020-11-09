# https://www.hackerrank.com/challenges/minimum-swaps-2/problem

def minimumSwaps(arr):
    swap = 0
    indexs = [0]*len(arr)

    # Iterate index and value together
    # store index of the value at the index which which is value
    for i, value in enumerate(arr):
        indexs[value-1] = i

    for i in range(len(arr)):

        # Check if the value is not at right index
        if arr[i] != i+1:

            # Get the right index of value
            index_value = indexs[i]

            # Swap the right value which is at Index and value at i
            arr[index_value] = arr[i]
            arr[i] = i+1

            # Now index of pervious arr[i] is index_value
            # so update it in indexs list
            indexs[arr[index_value]-1] = index_value

            # Now arr[i] or i+1 is at right index so update indexs list also
            indexs[i] = i

            swap += 1
    return swap
