
def get_bit(number, position):
    i = 1 << position
    result = (number & i) > 0
    print("number: {0}; position: {1}; result: {2}".format(number, position, result))

get_bit(4, 0)
get_bit(4, 1)
get_bit(4, 2)
get_bit(4, 3)
