max_number = 10000000
bitmap = []

for i in range (0, max_number):
    bitmap.append(0)

f = open("input.txt","r")
for line in f:
    number = int(line.rstrip())
    bitmap[number] = 1

output_file = open("output.txt", "w")
for i, bit in enumerate(bitmap):
    if bit == 1:
        line = str(i) + "\n"
        output_file.write(line)

output_file.close()
