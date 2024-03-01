
def clear_all_bits_from_most_significant_throught_i(number, position):
    mask = -1 << (position + 1)
    print("mask {0}. bin: {1}".format(mask, bin(mask)))
    result = number & mask
    print("number: {0}; position: {1}; result: {2}; binary: {3}".format(number, position, result, bin(result)))

clear_all_bits_from_most_significant_throught_i(5, 0)
clear_all_bits_from_most_significant_throught_i(5, 1)
clear_all_bits_from_most_significant_throught_i(5, 2)
clear_all_bits_from_most_significant_throught_i(5, 3)
