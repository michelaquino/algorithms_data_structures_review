https://www.hackerrank.com/challenges/one-week-preparation-kit-new-year-chaos/problem


---
Resolution explanation

Kept thinking I needed to track the number of swaps needed if I were to bubble sort the array, but realized I was overcomplicating the problem.

Since the limitation is "no more than two bribes", we simply maintain a window of expected sticker values of size 3:

    the sticker value we expect if no bribes occurred
    the sticker value we expect if one bribe occurred
    the sticker value we expect if two bribes occurred

We create a list to represent the window, expected = [1,2,3]

To avoid the alloc/deallocs with append() and pop() you could probably improve efficiency with an array size four and something fancy to maintain the pointers.

We then iterate through through the line of people, q.

    if q[i] equals expected[0], no bribes occurred.
    else if q[i] equals expected[1], increment bribes by 1
    else if q[i] equals expected[2], increment bribes by 2
    else a larger bribed occurred; print "Too chaotic" and return

Be sure to update expected at each iteration:

    at the beginning, append the next expected value, expected[2]+1
    If no bribes occurred, remove expected[0]
    If one bribe occurred, remove expected[1]
    If two bribes occurred, remove expected[2]

Here's an example:

# q = [3,1,4,6,2,5]
# e = [1,2,3]
#  q[0]==3==e[2], bribes+=2, e.append(3+1), e.pop(2); e = [1,2,4]
#  q[1]==1==e[0], bribes+=0, e.append(4+1), e.pop(0); e = [2,4,5]
#  q[2]==4==e[1], bribes+=1, e.append(5+1), e.pop(1); e = [2,5,6]
#  q[3]==6==e[2], bribes+=2, e.append(6+1), e.pop(2); e = [2,5,7]
#  q[4]==2==e[0], bribes+=0, e.append(7+1), e.pop(0); e = [5,7,8]
#  q[5]==5==e[0], bribes+=0, e.append(8+1), e.pop(0); e = [7,8,9]
