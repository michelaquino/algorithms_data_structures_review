# Flip bit to win:
# You have an integer and you cann flip exactly one bit from 0 to a 1.
# Write a code to find the legth of the longest sequence of 1s you could create
#
# Ex:
# Input: 1775 (or 11011101111)
# Output: 8
#
# Approach
# 1) You get an array of the sequence of 0's and 1's, ex for 11011101111, array starting with 0, from righ to left:
# [0, 4, 1, 3, 1, 2]
#
# 2) Merge the sequences of 1's when 0 representation is 1.
# On the array above, the index 2 is an 0 representation, so sum 4 and 3 plus the new 1, you have a sequence of 8 number 1
#

def get_array_representation(number):
    looking_for = 0
    array = []

    count = 0
    # 32 bits
    for i in range (32):
        if (number & 1) != looking_for:
            looking_for = looking_for ^ 1
            array.append(count)
            count = 0

        count += 1
        number = number >> 1

    return array


def flip_bit_to_win(number):
    array = get_array_representation(number)

    longest_sequence = 0

    lookinng_at = 0
    index = 0
    while index < len(array):

        # only 1 number 0
        if array[index] == 1:
            value_to_compare = 0

            if index == 0 and index+1 < len(array):
                value_to_compare = array[index+1]

            elif index == len(array) - 1:
                value_to_compare = array[index-1]

            else:
                value_to_compare = array[index-1] + array[index+1]


            if value_to_compare > longest_sequence:
                longest_sequence = value_to_compare + 1

        index += 2

    return longest_sequence



result = flip_bit_to_win(1775)
print(result)
