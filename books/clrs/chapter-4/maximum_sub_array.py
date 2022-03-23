def find_max_sub_array(array, start, end):
    if start == end:
        return start, end, array[start]

    middle = int((end + start)/2)

    left_start, left_end, left_sum = find_max_sub_array(array, start, middle)
    right_start, right_end, right_sum = find_max_sub_array(array, middle+1, end)
    cross_start, cross_end, cross_sum = find_cross_max_sub_array(array, start, middle, end)

    if left_sum > right_sum and left_sum > cross_sum:
        return left_start, left_end, left_sum
    elif right_sum > left_sum and right_sum > cross_sum:
        return right_start, right_end, right_sum
    else:
        return cross_start, cross_end, cross_sum



def find_cross_max_sub_array(array, start, middle, end):
    max_sub_array = []

    max_left = float("-inf")
    max_left_index = 0
    sum_left = 0

    i = middle
    while i >= start:
        sum_left += array[i]
        if sum_left > max_left:
            max_left = sum_left
            max_left_index = i

        i -= 1

    max_right = float("-inf")
    max_right_index = 0
    sum_right = 0
    j = middle + 1
    while j <= end:
        sum_right += array[j]
        if sum_right > max_right:
            max_right = sum_right
            max_right_index = j

        j += 1

    return (max_left_index, max_right_index, max_left + max_right)



def find_change_arrary(array):
    change_array = [0]
    i = 1
    while i < len(array):
        change_array.append(array[i]-array[i-1])
        i += 1

    return change_array


array = [100, 113, 110, 85, 105, 102, 86, 63, 81, 101, 94, 106, 101, 79, 94, 90, 97]
# array = [100, 113, 110, 85, 105, 102, 86, 63, 81]
# array = [100, 113, 110, 85, 105]
change_array = find_change_arrary(array)
print(change_array)

len_array = len(change_array) - 1
# print(find_cross_max_sub_array(change_array, 0, int(len_array/2), len_array))
print(find_max_sub_array(change_array,0, len_array))
