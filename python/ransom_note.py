#https://www.hackerrank.com/challenges/ctci-ransom-note/problem

def checkMagazine(magazine, note):
 items, c = dict(), 1
    for i in magazine:
        if i not in items:
            items[i] = 1
        else:
            items[i] += 1
    for j in note:
        if j not in magazine or items[j] == 0:
            c = 0
            break
        else:
            items[j] -= 1
    print('Yes') if c  else print('No')
