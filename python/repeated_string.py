# https://www.hackerrank.com/challenges/repeated-string/problem

def repeatedString(s, n):
    str_len, count = len(s), 0
    multiple = (n // str_len)
    remaining = n % str_len
    for i in s:
        if i == 'a':
            count +=1
    count *= multiple
    for i in range(remaining):
        if s[i] == 'a':
            count +=1
    return count
