def set_bit(number, position):
    i = 1 << position
    result = number | i
    print("number: {0}; position: {1}; result: {2}".format(number, position, result))

set_bit(4, 0)
set_bit(4, 1)
set_bit(4, 2)
set_bit(4, 3)
