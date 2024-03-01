import random

numberQuantity = 1000
max_number = 10000000

afile = open("input.txt", "w")

for i in range(numberQuantity):
    line = str(random.randint(1, max_number)) + "\n"
    afile.write(line)

afile.close()
