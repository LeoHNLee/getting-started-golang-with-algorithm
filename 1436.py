n = int(input())
i = 665
p = 0
while p < n:
    i += 1
    j = str(i)
    for k in range(len(j), 2, -1):
        is_d = int(j[:k])%1000 == 666
        if is_d:
            p += 1
            break
print(i)