def update_bit(number, position, value):
    number = number & ~(1 << position)
    value = value << position


    result = number | value
    print("number: {0}; position: {1}; result: {2}".format(number, position, result))

update_bit(4, 0, 1)
update_bit(4, 1, 1)
update_bit(4, 2, 1)
update_bit(4, 3, 1)
