
array = [4, 3, 9, 1]

result = 0

for n in array:
    result = result ^ (n << 1)

print(result)
