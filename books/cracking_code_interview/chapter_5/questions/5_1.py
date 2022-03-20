def printBit(name, integer):
    if integer >= 0:
        print("{0}: ".format(name) + bin(integer)[2:].zfill(32))
    else:
        print("{0}: ".format(name) + bin(2**32 + integer)[2:])

def insert_m_into_n(n, m, i, j):
    # step 1
    #   create mask to clear N from i to j
    #
    # step 2
    #   clear N with the mask
    #
    # step 3
    #   shift M left i times
    #
    # step 4
    #  n_cleared OR m_shifted
    printBit("m", m)
    printBit("n", n)

    allOne = ~0
    printBit("allOne", allOne)

    leftMask = allOne << (j + 1)
    printBit("leftMask", leftMask)

    rightMask = (1 << i) - 1
    printBit("rightMask", rightMask)

    mask = leftMask | rightMask
    printBit("mask", mask)

    n_cleared = n & mask
    printBit("n_cleared", n_cleared)

    m_shifted = m << i
    printBit("m_shifted", m_shifted)

    printBit("result", n_cleared | m_shifted)

insert_m_into_n(1024, 19, 2, 6)
