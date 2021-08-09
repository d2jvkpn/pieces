#! python3
# -*- coding: utf-8 -*-

def less (a, b): 
    x, y = str(a), str(b)
    n = max(len(x), len(y))
    x2, y2 = x.ljust(n, '0'), y.ljust(n, '0')

    return x2 < y2 or (x2 == y2 and len(x) > len(y))

class myInt(int):
    def __lt__(a, b):
        return less(a, b)
    def __gt__(a, b):
        return not less(a, b)

def intArry2number(arr):
    x = sorted([myInt(i) for i in arr], reverse=True)
    return int(''.join([str(i) for i in x]))

print(intArry2number([30, 3, 34, 5, 9]))
print(intArry2number([3, 30, 34, 5, 9]))
