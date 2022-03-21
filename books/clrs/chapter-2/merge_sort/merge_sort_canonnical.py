def merge_sort_canonnical(array, start, end):
    # One element
    if start >= end:
        return array[start:start+1]

    middle = int((start + end) / 2)
    right = merge_sort_canonnical(array, start, middle)
    left = merge_sort_canonnical(array, middle+1, end)
    array_merged = merge_cannonical(array, start, middle, end)

    return array_merged



def merge_cannonical(array, start, middle, end):
    size_left = middle - start + 1
    size_right = end - middle

    left = array[start:middle+1]
    right = array[middle+1:end+1]

    i = 0
    j = 0
    k = start

    while k <= end:
        if i < len(right) and j < len(left):
            if right[i] <= left[j]:
                array[k] = right[i]
                i += 1
            else:
                array[k] = left[j]
                j += 1
        elif i < len(right):
            array[k] = right[i]
            i += 1
        elif j < len(left):
            array[k] = left[j]
            j += 1

        k += 1


    return array

a = [5, 2, 4, 7, 1, 3, 2, 6]
print(merge_sort_canonnical(a, 2, 5))
print(merge_sort_canonnical(a, 1, 6))

# print(merge_cannonical(a, 2, 3, 5))
