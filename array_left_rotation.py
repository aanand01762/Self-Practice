# https://www.hackerrank.com/challenges/ctci-array-left-rotation/problem

def rotLeft(a, d):
    l, new_list = len(a), []
    iterations = l-d
    iterator = d
    while(iterations):
        new_list.append(a[iterator])
        iterations -= 1
        iterator += 1
    for i in range(d):
        new_list.append(a[i])
    return new_list
