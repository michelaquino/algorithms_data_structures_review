
n = 11
count = 0

while n > 0:
    count += 1
    n = n & (n-1)

print(count)
