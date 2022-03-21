def merge_sort(array):
    # One element
    if len(array) < 2:
        return array

    middle = int(len(array) / 2)
    right = array[:middle]
    left = array[middle:]

    right = merge_sort(right)
    left = merge_sort(left)
    array_merged = merge(left, right)

    return array_merged



def merge(left, right):
    i = 0
    j = 0

    array_merged = []

    while i < len(right) or j < len(left):
        if i < len(right) and j < len(left):
            if right[i] <= left[j]:
                array_merged.append(right[i])
                i += 1
            else:
                array_merged.append(left[j])
                j += 1
        elif i < len(right):
            array_merged.append(right[i])
            i += 1
        elif j < len(left):
            array_merged.append(left[j])
            j += 1


    return array_merged

a = [3, 1, 2, 6, 5, 2, 4, 7]
array_sorted = merge_sort(a)
print(array_sorted)
