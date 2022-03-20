# Given a real number between 0 and 1 (e.g 0.72) that is passed in as a double, print the binary representation.
# If the number cannot be represented accurately in binary with at most 32 characters, print "ERROR"

def print_real_number_as_binary(number):
    if number >= 1 or number <= 0:
        print("ERROR")

    original_number = number

    number_binary = ""

    while number > 0:
        if len(number_binary) > 32:
            number_binary = "ERROR"
            break

        double_number = number * 2
        if double_number >= 1:
            number_binary += "1"
            number = double_number-1
        else:
            number_binary += "0"
            number = double_number

    print("RESULT for {1} -> {0}".format(number_binary, original_number))


print_real_number_as_binary(0.72)
print_real_number_as_binary(0.5)
