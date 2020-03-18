# https://www.hackerrank.com/challenges/new-year-chaos/problem

def minimumBribes(q):
    bribes = 0
    
    # Travese the list from right to left
    for i in reversed(range(len(q))):
        
        # check if element is not at right position 
        if q[i] != i+1:
            
            # if element if one position left then one bribe is given
            # swap to its right psition
            if (q[i-1] >= 0 and q[i-1] == i+1):
                q[i-1] = q[i]
                q[i] = i+1
                bribes += 1
            
            # if element if two positions left, then one bribe is given
            # swap to its right psition
            elif q[i-2] >= 0 and q[i-2] == i+1:
                q[i-2] = q[i-1]
                q[i-1] = q[i]
                q[i] = i+1
                bribes += 2
            
            # if not possible 
            else:
                print('Too chaotic')
                return
    print(bribes)
