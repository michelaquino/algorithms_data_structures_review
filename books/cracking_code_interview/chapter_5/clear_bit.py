
def clear_bit(number, position):
    i = ~(1 << position)
    result = number & i
    print("number: {0}; position: {1}; result: {2}".format(number, position, result))

clear_bit(4, 0)
clear_bit(4, 1)
clear_bit(4, 2)
clear_bit(4, 3)
