import sys
from itertools import combinations
N, M = map(int, sys.stdin.readline().split())
cards = combinations(map(int, sys.stdin.readline().split()), 3)
mx = 0
for card in cards:
    temp = sum(card)
    if temp <= M:
        if temp > mx:
            mx = temp
            ret = temp
print(ret)