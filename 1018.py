import sys
N, M = map(int, input().split())
boards = [i.strip("\n") for i in sys.stdin.readlines()]

mn = 64

for n in range(N-7):
    for m in range(M-7):
        temp = 0
        right = boards[n][m]
        checker = (n%2 == 1)^(m%2 == 1)
        for i in range(n, n+8):
            for j in range(m, m+8):
                is_right = (checker == (i%2 == 1)^(j%2 == 1))
                if is_right:
                    if right != boards[i][j]:
                        temp += 1
                else:
                    if right == boards[i][j]:
                        temp += 1
        if temp > 32:
            temp = 64-temp
        if mn > temp:
            mn = temp
print(mn)
        