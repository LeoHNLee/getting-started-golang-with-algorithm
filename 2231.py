import sys

N = sys.stdin.readline()
digit = len(N)
N = int(N)
ret = 0
for n in range(N-digit*9, N):
    if n < 0:
        continue
    temp = n
    n = str(n)
    for i in n:
        temp += int(i)
    if temp == N:
        ret = n
        break
print(ret)