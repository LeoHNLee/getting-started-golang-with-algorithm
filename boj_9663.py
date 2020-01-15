import sys
from copy import deepcopy
sys.setrecursionlimit(10**8)
N = int(input())
ret = 0

def color(n, pos, chess):
    left = pos
    m = -1
    for i in range(n):
        chess[i] = chess[i][:pos] + "0" + chess[i][pos+1:]
    while left > 0 and m < n-1:
        left -= 1
        m += 1
        chess[m] = chess[m][:left] + "0" + chess[m][left+1:]
    m = -1
    while pos < N-1 and m < n-1:
        pos += 1
        m += 1
        chess[m] = chess[m][:pos] + "0" + chess[m][pos+1:]
    return chess

def check(n, pos, chess):
    if n == 1:
        global ret
        ret += 1
        return
    chess = color(n-1, pos, chess)
    for pos, v in enumerate(chess[0]):
        if int(v):
            check(n-1, pos, chess[1:])

for pos in range(N):
    chess = ["1"*N for _ in range(N)]
    check(N, pos, chess)

print(ret)