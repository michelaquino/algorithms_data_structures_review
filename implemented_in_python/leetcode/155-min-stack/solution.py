# https://leetcode.com/problems/min-stack/

import math


class Node:
    def __init__(self, value, min):
        self.min = min
        self.value = value
        self.next = None


class MinStack:
    def __init__(self):
        self.head = None

    def push(self, val: int) -> None:
        if self.head is not None:
            minVal = min(self.head.min, val)
        else:
            minVal = val

        node = Node(val, minVal)
        head = self.head
        node.next = head
        self.head = node

    def pop(self) -> None:
        val = self.head.value
        self.head = self.head.next

    def top(self) -> int:
        return self.head.value

    def getMin(self) -> int:
        return self.head.min


# Your MinStack object will be instantiated and called as such:
# obj = MinStack()
# obj.push(val)
# obj.pop()
# param_3 = obj.top()
# param_4 = obj.getMin()
minStack = MinStack()
# print(minStack.push(-2))
# print(minStack.push(0))
# print(minStack.push(-3))
# print(minStack.getMin())  # return -3
# print(minStack.pop())
# print(minStack.top())  # return 0
# print(minStack.getMin())  # return -2

print("push: ", minStack.push(2147483646))
print("push: ", minStack.push(2147483646))
print("push: ", minStack.push(2147483647))

print("top: ", minStack.top())
print("pop: ", minStack.pop())
print("getMin: ", minStack.getMin())

print("pop: ", minStack.pop())
print("getMin: ", minStack.getMin())
print("pop: ", minStack.pop())

print("push: ", minStack.push(2147483647))
print("top: ", minStack.top())
print("getMin: ", minStack.getMin())

print("push: ", minStack.push(-2147483648))
print("top: ", minStack.top())
print("getMin: ", minStack.getMin())

print("pop: ", minStack.pop())
print("getMin: ", minStack.getMin())
