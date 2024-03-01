def insertion_sort(array):
    i = 1
    while i < len(array):
        key = array[i]
        j = i-1

        while j >= 0 and array[j] > key:
            array[j+1] = array[j]
            j -= 1

        array[j+1] = key
        i += 1


array = [5,1,2,8,7]
insertion_sort(array)
