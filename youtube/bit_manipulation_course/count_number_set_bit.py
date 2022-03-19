
n = 12
count = 0

while n > 0:
    if n & 1 > 0:
        count += 1

    n = n >> 1

print(count)
